package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"google.golang.org/grpc"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

func serveStaticFiles() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("error while getting the runtime binary path; %v", err)
	}
	staticDir := strings.Replace(
		dir[0:strings.LastIndex(dir, "/main")]+"/static", "bazel-bin/", "", 1)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(staticDir)))
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("error while serving static files; %v", err)
	}
}

func main() {
	go serveStaticFiles()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := tpb.RegisterTickTackToeHandlerFromEndpoint(ctx, mux, "localhost:12345", opts)
	if err != nil {
		log.Fatalf("registering the endpoint failed; %v", err)
	}

	if err := http.ListenAndServe(":8080", wsproxy.WebsocketProxy(mux)); err != nil {
		log.Fatalf("listening to the port failed; %v", err)
	}
}
