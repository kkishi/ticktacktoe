package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"google.golang.org/grpc"

	tpb "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
)

var (
	enableWsproxyLogging = flag.Bool("enable_wsproxy_logging", false,
		"Whether to enable wsproxy logging.")
)

func serveStaticFiles() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("error while getting the runtime binary path; %v", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(strings.Replace(
		dir[0:strings.LastIndex(dir, "/main")]+"/static", "bazel-bin/", "", 1))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(
		http.Dir(dir[0:strings.LastIndex(dir, "/main")]+"/static"))))
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("error while serving static files; %v", err)
	}
}

type binaryMarshaler struct {
}

func (m *binaryMarshaler) ContentType() string {
	// Content type is irrelevant for websocket.
	return ""
}

func (m *binaryMarshaler) Marshal(v interface{}) ([]byte, error) {
	var p proto.Message
	// See https://github.com/grpc-ecosystem/grpc-gateway/blob/c6f7a5ac629444a556bb665e389e41b897ebad39/runtime/handler.go
	switch v := v.(type) {
	case proto.Message:
		p = v
	case map[string]proto.Message:
		p = v["result"]
	default:
		return nil, fmt.Errorf("non-proto value is provided: %#v", v)
	}
	b, err := proto.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("error while marshaling %#v; %v", p, err)
	}
	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(b64, b)
	log.Printf("binaryMarshaler.Marshal proto: %s, bytes: %v, base64: %s",
		p.String(), b, string(b64))
	return b64, nil
}

func (m *binaryMarshaler) Unmarshal(data []byte, v interface{}) error {
	p, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("non-proto value is provided: %#v", v)
	}
	b := make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	if n, err := base64.StdEncoding.Decode(b, data); err != nil {
		return fmt.Errorf("base64 decode failed; %v", err)
	} else {
		b = b[0:n]
	}
	if err := proto.Unmarshal(b, p); err != nil {
		return fmt.Errorf("error while unmarshaling proto; %v", err)
	}
	log.Printf("binaryMarshaler.Unmarshal proto: %s, bytes: %v, base64: %s",
		p.String(), b, string(data))
	return nil
}

type decoder struct {
	reader io.Reader
}

func (d *decoder) Decode(v interface{}) error {
	p, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("non-proto value is provided: %#v", v)
	}

	for {
		b64 := make([]byte, 1024)
		if n, err := d.reader.Read(b64); err != nil {
			return fmt.Errorf("no more data to read; %v", err)
		} else {
			b64 = b64[0:n]
		}
		if len(b64) == 1 && b64[0] == 10 {
			continue
		}
		b := make([]byte, base64.StdEncoding.DecodedLen(len(b64)))
		if n, err := base64.StdEncoding.Decode(b, b64); err != nil {
			return fmt.Errorf("base64 decode failed; %v", err)
		} else {
			b = b[0:n]
		}
		if err := proto.Unmarshal(b, p); err != nil {
			return fmt.Errorf("unmarshal failed; %v", err)
		}
		log.Printf("decoder.Decode proto: %s, bytes: %v, base64: %s",
			p.String(), b, string(b64))
		return nil
	}
}

func (m *binaryMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return &decoder{
		reader: r,
	}
}

func (m *binaryMarshaler) NewEncoder(w io.Writer) runtime.Encoder {
	log.Fatal("not implemented")
	return nil
}

type wsproxyLogger struct {
}

func (l *wsproxyLogger) Warnln(args ...interface{}) {
	if *enableWsproxyLogging {
		log.Println(args...)
	}
}

func (l *wsproxyLogger) Debugln(args ...interface{}) {
	if *enableWsproxyLogging {
		log.Println(args...)
	}
}

func main() {
	flag.Parse()

	go serveStaticFiles()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &binaryMarshaler{}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := tpb.RegisterTickTackToeHandlerFromEndpoint(ctx, mux, "localhost:12345", opts)
	if err != nil {
		log.Fatalf("registering the endpoint failed; %v", err)
	}

	if err := http.ListenAndServe(":8080", wsproxy.WebsocketProxy(mux,
		wsproxy.WithLogger(&wsproxyLogger{}))); err != nil {
		log.Fatalf("listening to the port failed; %v", err)
	}
}
