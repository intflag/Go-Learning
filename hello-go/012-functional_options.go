package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

//1 使用函数
func NewDefaultServer(addr string, port int) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
}

func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
}

//2 使用构造模式
type ServerBuilder struct {
	Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
	sb.Server.Addr = addr
	sb.Server.Port = port
	return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
	sb.Server.Protocol = protocol
	return sb
}

func (sb *ServerBuilder) WithTimeout(timeout time.Duration) *ServerBuilder {
	sb.Server.Timeout = timeout
	return sb
}

func (sb *ServerBuilder) WithTLS(tls *tls.Config) *ServerBuilder {
	sb.Server.TLS = tls
	return sb
}

func (sb *ServerBuilder) Build() (Server, error) {
	return sb.Server, nil
}

//3 函数式编程
type Option func(*Server)

func Protocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func MaxConns(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	return &srv, nil
}

func main() {
	//Go 语言不支持重载，只能使用不同的函数名称来进行创建
	server1, _ := NewDefaultServer("localhost", 8080)
	fmt.Println("server1", server1)
	server2, _ := NewTLSServer("192.168.0.1", 8080, nil)
	fmt.Println("server2", server2)

	//构造模式
	sb := ServerBuilder{}
	server3, _ := sb.Create("locolhost", 3306).
		WithProtocol("tcp").
		WithTimeout(5 * time.Second).
		Build()

	fmt.Println("server3", server3)

	//函数式编程
	server4, _ := NewServer("192.168.1.1", 6789)
	fmt.Println("server4", server4)
	server5, _ := NewServer("192.168.1.1", 6789, Protocol("udp"))
	fmt.Println("server5", server5)
}
