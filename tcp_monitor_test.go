package gerty

import (
	"testing"
	"time"
)

func TestShouldPingValidHostAndPort(t *testing.T) {
	host := "google.com"
	port := 80
	monitor := NewTcpMonitor("Tcp Ok", "this monitor pings google.com", host, port)
	status := monitor.Check()
	if status != OK {
		t.Fatalf("error while checking host %s:%d", host, port)
	}
}

func TestShouldFailTcpOnTimeout(t *testing.T) {
	// non-routeable IP address.
	host := "10.255.255.1"
	port := 80
	opts := TcpMonitorOptions{Checks: 5, Timeout: 1 * time.Second}
	monitor := NewTcpMonitorWithOptions("Tcp Timeout Monitor", "This monitor should timeout", host, port, &opts)
	status := monitor.Check()
	if status != NOK {
		t.Fatalf("http monitor should timeout and fail")
	}
}
