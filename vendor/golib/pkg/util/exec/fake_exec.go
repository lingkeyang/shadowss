package exec

import (
	"fmt"
)

// A simple scripted Interface type.
type FakeExec struct {
	CommandScript []FakeCommandAction
	CommandCalls  int
	LookPathFunc  func(string) (string, error)
}

type FakeCommandAction func(cmd string, args ...string) Cmd

func (fake *FakeExec) Command(cmd string, args ...string) Cmd {
	if fake.CommandCalls > len(fake.CommandScript)-1 {
		panic(fmt.Sprintf("ran out of Command() actions. Could not handle command [%d]: %s args: %v", fake.CommandCalls, cmd, args))
	}
	i := fake.CommandCalls
	fake.CommandCalls++
	return fake.CommandScript[i](cmd, args...)
}

func (fake *FakeExec) LookPath(file string) (string, error) {
	return fake.LookPathFunc(file)
}

// A simple scripted Cmd type.
type FakeCmd struct {
	Argv                 []string
	CombinedOutputScript []FakeCombinedOutputAction
	CombinedOutputCalls  int
	CombinedOutputLog    [][]string
	Dirs                 []string
}

func InitFakeCmd(fake *FakeCmd, cmd string, args ...string) Cmd {
	fake.Argv = append([]string{cmd}, args...)
	return fake
}

type FakeCombinedOutputAction func() ([]byte, error)

func (fake *FakeCmd) SetDir(dir string) {
	fake.Dirs = append(fake.Dirs, dir)
}

func (fake *FakeCmd) CombinedOutput() ([]byte, error) {
	if fake.CombinedOutputCalls > len(fake.CombinedOutputScript)-1 {
		panic("ran out of CombinedOutput() actions")
	}
	if fake.CombinedOutputLog == nil {
		fake.CombinedOutputLog = [][]string{}
	}
	i := fake.CombinedOutputCalls
	fake.CombinedOutputLog = append(fake.CombinedOutputLog, append([]string{}, fake.Argv...))
	fake.CombinedOutputCalls++
	return fake.CombinedOutputScript[i]()
}

func (fake *FakeCmd) Output() ([]byte, error) {
	return nil, fmt.Errorf("unimplemented")
}

// A simple fake ExitError type.
type FakeExitError struct {
	Status int
}

func (fake *FakeExitError) String() string {
	return fmt.Sprintf("exit %d", fake.Status)
}

func (fake *FakeExitError) Error() string {
	return fake.String()
}

func (fake *FakeExitError) Exited() bool {
	return true
}

func (fake *FakeExitError) ExitStatus() int {
	return fake.Status
}
