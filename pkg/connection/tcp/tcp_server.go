package tcp

import (
	"net"
	"reflect"
	"strconv"
	"sync"
	"syscall"
	"time"

	"shadowss/pkg/access"
	"shadowss/pkg/config"
	"shadowss/pkg/multiuser/apiserverproxy"

	"shadowss/pkg/connection/tcp/ssclient"
	"shadowss/pkg/crypto"

	"github.com/golang/glog"
	"golang.org/x/net/context"
)

//TCPServer maintain a listener
type TCPServer struct {
	Config          *config.ConnectionInfo
	quit            chan struct{}
	uploadTraffic   int64 //request upload traffic
	downloadTraffic int64 //request download traffic
	lastActiveTime  time.Time

	//mutex Mutex used to serialize access to the dictionary
	mapMutex  *sync.Mutex
	dataMutex *sync.Mutex

	//store connection count for server port
	portConnectionCount map[int]uint
}

type connector struct {
	client     *ssclient.Client
	serverConn map[string]net.Conn //proxy to remote connnection
}

//NewTCPServer create a TCPServer
func NewTCPServer(cfg *config.ConnectionInfo) *TCPServer {
	if cfg.MaxConnection <= 0 {
		cfg.MaxConnection = 50
	}
	return &TCPServer{
		Config:              cfg,
		lastActiveTime:      time.Now(),
		quit:                make(chan struct{}),
		mapMutex:            new(sync.Mutex),
		dataMutex:           new(sync.Mutex),
		portConnectionCount: make(map[int]uint),
	}
}

//Stop implement quit go routine
func (tcpSrv *TCPServer) Stop() {
	glog.V(5).Infof("tcp server close %v\r\n", tcpSrv.Config)
	close(tcpSrv.quit)
}

//Traffic ollection traffic for client,return upload traffic and download traffic
func (tcpSrv *TCPServer) Traffic() (int64, int64) {
	lock(tcpSrv.dataMutex)
	upload := tcpSrv.uploadTraffic
	download := tcpSrv.downloadTraffic

	//clear traffic
	tcpSrv.uploadTraffic = 0
	tcpSrv.downloadTraffic = 0
	unlock(tcpSrv.dataMutex)

	return upload, download
}

func (tcpSrv *TCPServer) handleRequest(ctx context.Context, acceptConn net.Conn) {

	srvPort := acceptConn.LocalAddr().(*net.TCPAddr).Port
	lock(tcpSrv.mapMutex)
	limitConn, ok := tcpSrv.portConnectionCount[srvPort]
	unlock(tcpSrv.mapMutex)
	if ok {
		if limitConn > uint(tcpSrv.Config.MaxConnection) {
			glog.Errorf("limit connection, refuse this connection(%v)", acceptConn)
			acceptConn.Close()
			return
		}
	}
	limitConn++

	lock(tcpSrv.mapMutex)
	tcpSrv.portConnectionCount[srvPort] = limitConn
	unlock(tcpSrv.mapMutex)

	reqAddr := acceptConn.RemoteAddr().String()
	timeout := time.Duration(tcpSrv.Config.Timeout*2) * time.Second
	crypto, err := crypto.NewCrypto(tcpSrv.Config.EncryptMethod, tcpSrv.Config.Password)
	if err != nil {
		glog.Errorf("create crypto error :%v", err)
		limitConn--
		lock(tcpSrv.mapMutex)
		tcpSrv.portConnectionCount[srvPort] = limitConn
		unlock(tcpSrv.mapMutex)

		acceptConn.Close()

		return
	}
	client := ssclient.NewClient(acceptConn, crypto)
	connHelper := &connector{
		client:     client,
		serverConn: make(map[string]net.Conn),
	}
	defer func() {
		connHelper.client.Close()

		limitConn--
		lock(tcpSrv.mapMutex)
		tcpSrv.portConnectionCount[srvPort] = limitConn
		unlock(tcpSrv.mapMutex)
	}()

	ssProtocol, err := connHelper.client.ParseTcpReq()
	if err != nil {
		glog.Errorf("get invalid request  from %s error %v\r\n", reqAddr, err)
		return
	}
	if ssProtocol.OneTimeAuth != tcpSrv.Config.EnableOTA {
		glog.Errorf("not match one time auth with config(%v:%v)\r\n", ssProtocol.OneTimeAuth, tcpSrv.Config.EnableOTA)
		return
	}

	remoteAddr := &net.TCPAddr{
		IP:   ssProtocol.DstAddr.IP,
		Port: ssProtocol.DstAddr.Port,
	}
	host := apiserverproxy.FilterRequest(remoteAddr)

	remote, err := net.Dial("tcp", host)
	if err != nil {
		if ne, ok := err.(*net.OpError); ok && (ne.Err == syscall.EMFILE || ne.Err == syscall.ENFILE) {
			glog.Errorf("dial error:%v\r\n", err)
		} else {
			glog.Errorf(" connecting to:%v occur err:%v", host, err)
		}
		return
	}
	defer remote.Close()

	type result struct {
		uploadTraffic   int64
		downloadTraffic int64
	}

	var wg sync.WaitGroup
	pipeResult := make(chan result, 1)
	go func() {
		wg.Add(1)
		upload, download := PipeData(ctx, connHelper.client, remote, timeout)
		pipeResult <- result{upload, download}
		wg.Done()
	}()

	for {

		select {
		case <-ctx.Done():
			glog.V(5).Infof("handle %s read requet will be done\n", reqAddr)
			wg.Wait()
			return
		case result := <-pipeResult:
			lock(tcpSrv.dataMutex)
			tcpSrv.uploadTraffic += result.uploadTraffic
			tcpSrv.downloadTraffic += result.downloadTraffic
			unlock(tcpSrv.dataMutex)

			glog.V(5).Infof("handle %s read requet will be done with result %+v\n", reqAddr, result)
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}

}

//Run start a tcp listen for user
func (tcpSrv *TCPServer) Run() {

	port := tcpSrv.Config.Port
	//allways realloc port for user except 48888
	if port != 48888 {
		port = 0
	}

	portStr := strconv.Itoa(port)
	ln, err := net.Listen("tcp", ":"+portStr)
	if err != nil {
		glog.Errorf("tcp server(%v) error: %v\n", port, err)
	}
	defer ln.Close()

	//update user port if input config port equal 0
	tcpSrv.Config.Port = ln.Addr().(*net.TCPAddr).Port

	var ctx context.Context
	ctx, cancel := context.WithCancel(context.TODO())

	defer func() {
		cancel()
		ln.Close()
		access.TurnoffLocalPort(tcpSrv.Config.Port, string("tcp"))
	}()

	//open port
	err = access.OpenLocalPort(tcpSrv.Config.Port, string("tcp"))
	if err != nil {
		glog.Warningf("open port(%d) on local host firewall error %v", tcpSrv.Config.Port, err)
	}

	var wg sync.WaitGroup
	type accepted struct {
		conn net.Conn
		err  error
	}
	for {
		c := make(chan accepted, 1)
		go func() {
			glog.V(5).Infof("wait for accept on %v\r\n", tcpSrv.Config.Port)
			var conn net.Conn
			conn, err = ln.Accept()
			c <- accepted{conn: conn, err: err}
		}()

		select {
		case <-tcpSrv.quit:
			glog.Infof("Receive Quit singal for %s\r\n", port)
			return
		case accept := <-c:
			if accept.err != nil {
				glog.V(5).Infof("accept error: %v\n", accept.err)
				continue
			}
			reqAddr := accept.conn.RemoteAddr().String()

			glog.V(5).Infof("accept remote client: %v\n", reqAddr)
			wg.Add(1)
			go func() {
				tcpSrv.lastActiveTime = time.Now()
				tcpSrv.handleRequest(ctx, accept.conn)
				wg.Done()
			}()
		}
	}
}

func (tcpSrv *TCPServer) Compare(client *config.ConnectionInfo) bool {
	return reflect.DeepEqual(*tcpSrv.Config, *client)
}

func (tcpSrv *TCPServer) GetListenPort() int {
	return tcpSrv.Config.Port
}

func (tcpSrv *TCPServer) GetConfig() config.ConnectionInfo {
	return *tcpSrv.Config
}

func (tcpSrv *TCPServer) GetLastActiveTime() time.Time {
	return tcpSrv.lastActiveTime
}

func lock(mutex *sync.Mutex) {
	mutex.Lock()
}

func unlock(mutex *sync.Mutex) {
	mutex.Unlock()
}
