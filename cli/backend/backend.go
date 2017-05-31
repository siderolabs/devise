package backend

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Masterminds/sprig"
	"github.com/autonomy/devise/api"
	"github.com/autonomy/devise/cli/backend/modifiers"
	"github.com/autonomy/devise/storage"
	"github.com/autonomy/devise/storage/datastore"
	"golang.org/x/net/context"
)

// Server represents the Devise server.
type Server struct {
	Storage   datastore.Datastore
	Modifiers *modifiers.Modifiers
}

// Template implements api.DeviseServer.
func (s *Server) Template(ctx context.Context, in *api.TemplateRequest) (*api.TemplateReply, error) {
	var wr bytes.Buffer
	tmpl, err := template.New("base").Funcs(sprig.FuncMap()).Parse(string(in.Template))
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&wr, s.Modifiers)
	if err != nil {
		return nil, err
	}

	return &api.TemplateReply{Rendered: wr.Bytes()}, nil
}

// Start starts the gRPC server
func Start(port, datastore string) {
	s := grpc.NewServer()
	server := &Server{
		Storage:   storage.New(datastore),
		Modifiers: modifiers.NewModifiers(),
	}
	// Register Devise service on gRPC server.
	api.RegisterDeviseServer(s, server)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go s.Serve(lis)
}
