package monitors

import (
	"fmt"
	"log"
	"net"
	"time"
  util "github.com/gerty-monit/core/util"
)

type TcpMonitor struct {
	title       string
	description string
	host        string
	port        int
	buffer      util.CircularBuffer
	opts        *TcpMonitorOptions
}

type TcpMonitorOptions struct {
	Checks  int
	Timeout time.Duration
}

var DefaultTcpMonitorOptions = TcpMonitorOptions{
	Checks:  5,
	Timeout: 10 * time.Second,
}

func NewTcpMonitorWithOptions(title, description, host string, port int, opts *TcpMonitorOptions) *TcpMonitor {
	if opts == nil {
		opts = &DefaultTcpMonitorOptions
	}
	buffer := util.NewCircularBuffer(opts.Checks)
	return &TcpMonitor{title, description, host, port, buffer, opts}
}

func NewTcpMonitor(title, description, host string, port int) *TcpMonitor {
	return NewTcpMonitorWithOptions(title, description, host, port, nil)
}

func (monitor *TcpMonitor) Check() int {
	log.Printf("checking monitor %s", monitor.Name())
	address := fmt.Sprintf("%s:%d", monitor.host, monitor.port)
	conn, err := net.DialTimeout("tcp", address, monitor.opts.Timeout)

	if err == nil {
		defer conn.Close()
		monitor.buffer.Append(OK)
		return OK
	} else {
		log.Printf("tcp monitor check failed, error: %v", err)
		monitor.buffer.Append(NOK)
		return NOK
	}
}

func (monitor *TcpMonitor) Values() []int {
	return monitor.buffer.Values
}

func (monitor *TcpMonitor) Name() string {
	return monitor.title
}

func (monitor *TcpMonitor) Description() string {
	return monitor.description
}
