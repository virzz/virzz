package netlog

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/mozhu1024/virzz/logger"
)

type TCPLogServer struct {
	host     string
	port     int
	timeout  int
	doneChan chan int
}

func runTCPLogServer(host string, port int, timeout ...int) (string, error) {
	tcpServer := NewTCPLogServer(host, port)
	go func() {
		err := tcpServer.ListenAndServe()
		if err != nil {
			logger.Error(err)
		}
	}()
	interrupt := make(chan os.Signal, 1)
	sig := <-interrupt
	if err := tcpServer.Shutdown(context.Background()); err != nil {
		logger.Error("TCPLog Server Shutdown Error", err)
	}
	logger.Info("TCPLog Server Shutdown ...")
	if sig == os.Interrupt {
		return "tcplog server was interruped by system signal", nil
	}
	return "tcplog server was killed", nil
}

func NewTCPLogServer(host string, port int, timeout ...int) *TCPLogServer {
	_timeout := 5
	if len(timeout) > 0 { // && timeout[0] > 0
		_timeout = timeout[0]
	}
	server := TCPLogServer{
		host:    host,
		port:    port,
		timeout: _timeout,
	}
	return &server
}

func (s *TCPLogServer) ListenAndServe() (err error) {
	s.doneChan = make(chan int, 1)
	ip := net.ParseIP(s.host)
	if ip == nil {
		return fmt.Errorf("host is error : %s", s.host)
	}
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("listen fail err", err)
		return err
	}
	logger.SuccessF("TCPLog Server Listen on: %s", addr)
	return s.Serve(ln)
}

func (s *TCPLogServer) Shutdown(ctx context.Context) error {
	// TODO: context.Context
	s.doneChan <- 1
	return nil
}

func (s *TCPLogServer) Serve(l net.Listener) error {
	defer l.Close()
	tempDelay := time.Duration(1) * time.Second
	for {
		client, err := l.Accept()
		if err != nil {
			logger.Error("accept fail err", err)
			select {
			case <-s.doneChan:
				close(s.doneChan)
				return fmt.Errorf("TCPLog Server Closed")
			default:
			}
			logger.ErrorF("TCPLog Accept error: %v; retrying in %v", err, tempDelay)
			time.Sleep(tempDelay)
			continue
		}
		go process(client)
	}
}

func process(client net.Conn) {
	rAddr := client.RemoteAddr().String()
	defer logger.WarnF("Conn %s closed", rAddr)
	defer client.Close()
	reader := bufio.NewReader(client)
	// TODO: Control Close client conn
	var buf [4096]byte
	for {
		n, err := reader.Read(buf[:])
		if err != nil {
			if err != io.EOF {
				logger.ErrorF("[%s] Read %s", rAddr, err)
			}
			break
		}
		// TODO: Log to Database ?
		logger.InfoF("[%s]: \n%s", rAddr, string(buf[:n-1]))
	}
}
