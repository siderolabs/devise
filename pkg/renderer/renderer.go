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
	"github.com/autonomy/devise/pkg/modifier"
	"github.com/autonomy/devise/pkg/storage/datastore"
)

// Renderer is an object that discovers use to render templates.
type Renderer struct {
	modifiers *modifier.Modifiers
}

// NewRenderer instantiates and returns a Renderer.
func NewRenderer(m *modifier.Modifiers) *Renderer {
	return &Renderer{modifiers: m}
}

// Execute executes the template.
func (r *Renderer) Execute(name string, ip string, d datastore.Datastore) (err error) {
	log.Printf("Dialing IP %s", ip)
	err = d.Put(&datastore.Entry{Key: name, Value: []byte(ip)})
	if err != nil {
		log.Printf("Failed to insert entry: %v", err)
	}

	conn, err := grpc.Dial(ip+":50000", grpc.WithInsecure())
	if err != nil {
		grpclog.Printf("fail to dial: %v", err)
	}
	defer func() {
		deferErr := conn.Close()
		if deferErr != nil {
			return
		}
	}()
	client := api.NewDeviseClient(conn)
	in, err := client.OpenTemplate(context.Background(), &api.OpenTemplateRequest{})
	if err != nil {
		grpclog.Printf("failed to get template: %v", err)
	}

	r.modifiers.Vault.Client.SetToken(in.VaultToken)
	defer r.modifiers.Vault.Client.ClearToken()
	var wr bytes.Buffer
	tmpl, err := template.New("base").Funcs(sprig.TxtFuncMap()).Parse(string(in.Template))
	if err != nil {
		return err
	}

	err = tmpl.Execute(&wr, r.modifiers)
	if err != nil {
		return err
	}

	u, err := client.RenderTemplate(context.Background(), &api.RenderTemplateRequest{Rendered: wr.Bytes()})
	if err != nil {
		grpclog.Printf("failed to get template: %v", err)
	}
	grpclog.Printf("%v", u)

	return nil
}
