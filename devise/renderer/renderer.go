package renderer

import (
	"bytes"
	"context"
	"log"
	"text/template"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	"github.com/Masterminds/sprig"
	"github.com/autonomy/devise/api"
	"github.com/autonomy/devise/devise/modifiers"
	"github.com/autonomy/devise/devise/storage/datastore"
)

// Renderer is an object that discovers use to render templates.
type Renderer struct {
	modifiers *modifiers.Modifiers
}

// NewRenderer instantiates and returns a Renderer.
func NewRenderer(m *modifiers.Modifiers) *Renderer {
	return &Renderer{modifiers: m}
}

// Template implements api.DeviseServer.
func (r *Renderer) Template(in *api.TemplateReply) ([]byte, error) {
	r.modifiers.Vault.Client.SetToken(in.VaultToken)
	defer r.modifiers.Vault.Client.ClearToken()
	var wr bytes.Buffer
	tmpl, err := template.New("base").Funcs(sprig.TxtFuncMap()).Parse(string(in.Template))
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&wr, r.modifiers)
	if err != nil {
		return nil, err
	}

	return wr.Bytes(), nil
}

// Render executes the template.
func (r *Renderer) Render(name string, ip string, d datastore.Datastore) {
	log.Printf("Dialing IP %s", ip)
	err := d.Put(&datastore.Entry{Key: name, Value: []byte(ip)})
	if err != nil {
		log.Printf("Failed to insert entry: %v", err)
	}

	conn, err := grpc.Dial(ip+":50000", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer func() {
		deferErr := conn.Close()
		if deferErr != nil {
			return
		}
	}()
	client := api.NewDeviseClient(conn)
	t, err := client.GetTemplate(context.Background(), &api.TemplateRequest{})
	if err != nil {
		grpclog.Fatalf("failed to get template: %v", err)
	}

	rendered, err := r.Template(t)
	if err != nil {
		return
	}
	u, err := client.SendRenderedTemplate(context.Background(), &api.RenderRequest{Rendered: rendered})
	if err != nil {
		grpclog.Fatalf("failed to get template: %v", err)
	}
	grpclog.Printf("%v", u)
}
