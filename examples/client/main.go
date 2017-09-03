package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/autonomy/devise/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type app struct {
	config []byte
}

func (a *app) OpenTemplate(ctx context.Context, t *api.OpenTemplateRequest) (reply *api.OpenTemplateReply, err error) {
	vaultToken := os.Getenv("VAULT_TOKEN")
	b, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Printf("%v", err)
	}

	reply = &api.OpenTemplateReply{Template: b, VaultToken: vaultToken}

	return
}

func (a *app) RenderTemplate(ctx context.Context, t *api.RenderTemplateRequest) (reply *api.RenderTemplateReply, err error) {
	grpclog.Printf("%v", t)
	a.config = t.Rendered
	reply = &api.RenderTemplateReply{Success: true}

	return
}

func main() {
	app := &app{}
	s := grpc.NewServer()
	// Register Devise service on gRPC server.
	api.RegisterDeviseServer(s, app)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "50000"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		return
	}
}
