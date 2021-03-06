package proxyserver

import (
	"fmt"
	"shadowss/pkg/config"
	"shadowss/pkg/connection/tcp"
	"shadowss/pkg/connection/udp"
	"time"

	"github.com/golang/glog"
)

//Servers hold on a list of proxyserver
type Servers struct {
	tcpSrvMap map[string]ProxyServer //config line name to srv interface map
	udpSrvMap map[string]ProxyServer //config line name to srv interface map
	enableUDP bool
}

//UserInfo contains user connection information
type UserInfo struct {
	ConnectInfo    *config.ConnectionInfo
	LastActiveTime time.Time
}

//NewServers create a servers
func NewServers(udp bool) *Servers {
	return &Servers{
		tcpSrvMap: make(map[string]ProxyServer),
		udpSrvMap: make(map[string]ProxyServer),
		enableUDP: udp,
	}
}

func (srv *Servers) storeSrv(tcp ProxyServer, udp ProxyServer, cfg *config.ConnectionInfo) {
	srv.tcpSrvMap[cfg.Name] = tcp

	glog.Infof("store id %d  tcp :%p\r\n", cfg.ID, tcp)

	if srv.enableUDP {
		srv.udpSrvMap[cfg.Name] = udp
	}
}

//CheckServer Compare configure if exist
func (srv *Servers) CheckServer(client *config.ConnectionInfo) (bool, bool) {

	var equal bool
	tcpSrv, exist := srv.tcpSrvMap[client.Name]
	if exist {
		equal = tcpSrv.Compare(client)
	}

	return exist, equal
}

//GetListenPort get listen port by config user
func (srv *Servers) GetListenPort(client *config.ConnectionInfo) (int, error) {
	tcpSrv, exist := srv.tcpSrvMap[client.Name]
	if exist {
		return tcpSrv.GetListenPort(), nil
	}

	return 0, fmt.Errorf("not found this service")
}

//GetTraffic collection traffic for user,return upload traffic and download traffic
func (srv *Servers) GetTraffic(client *config.ConnectionInfo) (int64, int64, error) {
	var tcpUpload, tcpDownload, udpUpload, udpDownload int64
	tcpSrv, exist := srv.tcpSrvMap[client.Name]
	if exist {
		tcpUpload, tcpDownload = tcpSrv.Traffic()
		glog.V(5).Infof("Got %d Tcp traffic upload %d download:%d\r\n", client.Port, tcpUpload, tcpDownload)
	}

	if srv.enableUDP {
		udpSrv, exist := srv.udpSrvMap[client.Name]
		if exist {
			udpUpload, udpDownload = udpSrv.Traffic()
			glog.V(5).Infof("Got %d udp traffic upload %d download:%d\r\n", client.Port, udpUpload, udpDownload)
		} else {
			glog.Errorf("use udp relay but not found\r\n")
		}
	}

	return tcpUpload + udpUpload, tcpDownload + udpDownload, nil
}

//StopServer stop server only
func (srv *Servers) StopServer(client *config.ConnectionInfo) {
	tcpSrv, ok := srv.tcpSrvMap[client.Name]
	if !ok {
		glog.Warningf("not found tcp server %s\r\n", client.Port)
	} else {
		tcpSrv.Stop()

	}

	if srv.enableUDP {
		udpSrv, ok := srv.udpSrvMap[client.Name]
		if !ok {
			glog.Warningf("not found tcp server %s\r\n", client.Port)
		} else {
			udpSrv.Stop()

		}
	}
}

//CleanUpServer delete server from proxy manage. not use
func (srv *Servers) CleanUpServer(client *config.ConnectionInfo) {

	_, ok := srv.tcpSrvMap[client.Name]
	if !ok {
		glog.Warningf("not found tcp server %s\r\n", client.Port)
	} else {
		delete(srv.tcpSrvMap, client.Name)
	}

	if srv.enableUDP {
		_, ok := srv.udpSrvMap[client.Name]
		if !ok {
			glog.Warningf("not found tcp server %s\r\n", client.Port)
		} else {
			delete(srv.udpSrvMap, client.Name)
		}
	}

}

//StartWithConfig create new server for users
func (srv *Servers) StartWithConfig(v *config.ConnectionInfo) {

	glog.V(5).Infof("Start with %v at %p\r\n", v, v)

	tcpSrv := tcp.NewTCPServer(v)
	go tcpSrv.Run()

	//we need to ensure tcp and udp on same port
	time.Sleep(1 * time.Second)
	port := tcpSrv.GetListenPort()
	//fix config for udp
	v.Port = port

	var udpSrv ProxyServer
	if srv.enableUDP {
		udpSrv = udp.NewUDPServer(v)
		go udpSrv.Run()
	}

	srv.storeSrv(tcpSrv, udpSrv, v)
}

//Start create new server for user
func (srv *Servers) Start() {
	for idx := range config.ServerCfg.Clients {
		config := &config.ServerCfg.Clients[idx]
		srv.StartWithConfig(config)
	}
}

//GetUsersConfig get current all proxy config user
func (srv *Servers) GetUsersConfig() []config.ConnectionInfo {
	var users []config.ConnectionInfo
	for _, handler := range srv.tcpSrvMap {
		config := handler.GetConfig()
		users = append(users, config)
	}

	return users
}

//GetUsersInfor get current all proxy config user info
func (srv *Servers) GetUsersInfor() []UserInfo {
	var users []UserInfo
	for _, handler := range srv.tcpSrvMap {
		config := handler.GetConfig()
		lastTime := handler.GetLastActiveTime()

		item := UserInfo{}
		item.ConnectInfo = &config
		item.LastActiveTime = lastTime

		users = append(users, item)
	}

	return users
}
