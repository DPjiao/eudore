package server

import (
	"context"
	"github.com/eudore/eudore/protocol"
	"github.com/eudore/eudore/protocol/fastcgi"
	"github.com/eudore/eudore/protocol/http"
	"github.com/eudore/eudore/protocol/http2"
	"github.com/eudore/eudore/protocol/server"
	"testing"
)

func TestServerHttp(t *testing.T) {
	server := &server.Server{
		Handler: protocol.HandlerFunc(func(ctx context.Context, w protocol.ResponseWriter, r protocol.RequestReader) {
			w.Header().Add("Server", "simple server")
			w.Write([]byte("hello http server. your remote addr is " + r.RemoteAddr()))
		}),
	}
	server.SetHandler(http.NewHttpHandler(server.Handler))
	t.Log(server.ListenAndServe(":8085", nil))
}

func TestServerTls(t *testing.T) {
	server := &server.Server{
		Handler: protocol.HandlerFunc(func(ctx context.Context, w protocol.ResponseWriter, r protocol.RequestReader) {
			w.Push("/favicon.ico", nil)
			w.Header().Add("Server", "simple server")
			w.Write([]byte("hello http server. your remote addr is " + r.RemoteAddr()))
		}),
	}
	server.SetHandler(http.NewHttpHandler(server.Handler))
	server.SetNextHandler("h2", http2.NewServer(server.Handler))
	t.Log(server.ListenAndServeTls(":8085", "/etc/nginx/openssl/wejass.com/wejass.com.cer", "/etc/nginx/openssl/wejass.com/wejass.com.key", nil))
}

func TestServerFastcgi(t *testing.T) {
	server := &server.Server{
		Handler: protocol.HandlerFunc(func(ctx context.Context, w protocol.ResponseWriter, r protocol.RequestReader) {
			w.Header().Add("Server", "simple server")
			w.Write([]byte("hello http server. your remote addr is " + r.RemoteAddr()))
		}),
	}
	server.SetHandler(fastcgi.NewServer(server.Handler))
	t.Log(server.ListenAndServe(":8085", nil))
}
