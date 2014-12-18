package httputil

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func Listen(nw, laddr string) (net.Listener, error) {
	return net.Listen(nw, laddr)
}

func ListenFD(fd uint) (net.Listener, error) {
	file := os.NewFile(uintptr(fd), "")
	defer file.Close()
	return net.FileListener(file)
}

func ListenAndServe(addr string, handler http.Handler) error {
	ln, err := Listen("tcp", addr)
	if err != nil {
		return err
	}
	return Serve(ln, handler)
}

func ListenFDAndServe(fd uint, handler http.Handler) error {
	ln, err := ListenFD(fd)
	if err != nil {
		return err
	}
	return Serve(ln, handler)
}

func Serve(ln net.Listener, handler http.Handler) error {
	srv := &Server{}
	return srv.Serve(ln, handler)
}

func (srv *Server) Serve(ln net.Listener, handler http.Handler) error {
	addr := ln.Addr().String()
	kaln := tcpKeepAliveListener{ln.(*net.TCPListener)}
	httpSrv := &http.Server{Addr: addr, Handler: handler}
	srv.Server = httpSrv
	srv.signalHandler(kaln)
	return httpSrv.Serve(kaln)
}

func (srv *Server) signalHandler(ln net.Listener) {
	sigCh := make(chan os.Signal, 10)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer close(sigCh)
		<-sigCh
		signal.Stop(sigCh)
		ln.Close()
	}()
}

// tcpKeepAliveListener is copied from net/http.
// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}
