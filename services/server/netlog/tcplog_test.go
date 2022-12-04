package netlog

import (
	"context"
	"net"
	"testing"
	"time"
)

func TestTcpLogServer(t *testing.T) {
	s := NewTCPLogServer("127.0.0.1", 9988)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			t.Log(err)
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		tcpAddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:9988")
		conn, err := net.DialTCP("tcp4", nil, tcpAddr)
		if err != nil {
			t.Log(err)
			return
		}
		defer t.Log("Conn Close")
		defer conn.Close()
		_, err = conn.Write([]byte("hello world\n"))
		if err != nil {
			t.Log("Write", err)
			return
		}
		_, err = conn.Write([]byte("hello golang\n"))
		if err != nil {
			t.Log("Write", err)
			return
		}
	}()

	time.Sleep(5 * time.Second)
	s.Shutdown(context.Background())

}
