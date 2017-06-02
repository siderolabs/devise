package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/autonomy/devise/api"

	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50000"
	port    = ":8081"
)

type app struct {
	config []byte
}

type response map[string]interface{}

func main() {
	app, err := newApp()
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}
	var srv http.Server
	http2.VerboseLogs = true
	srv.Addr = port
	// This enables http2 support
	http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/config", app.templateHandler)
	log.Fatal(srv.ListenAndServe())
}

func (app *app) templateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{"config": string(app.config)})
}

func newApp() (*app, error) {
	b, err := ioutil.ReadFile("example.yaml")
	if err != nil {
		log.Printf("%v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := api.NewDeviseClient(conn)

	// Contact the server and get the rendered plan.
	r, err := c.Template(context.Background(), &api.TemplateRequest{Template: b, VaultToken: os.Getenv("VAULT_TOKEN")})
	if err != nil {
		return nil, err
	}

	return &app{config: r.Rendered}, nil
}
