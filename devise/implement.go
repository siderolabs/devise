package devise

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"

	"github.com/autonomy/devise/api"
	"google.golang.org/grpc"
)

// Plan represents a plan.
type Plan struct {
	Templates map[string]*Template `yaml:"templates"`
	Script    *string              `yaml:"script,omitempty"`
}

// Template represents a template.
type Template struct {
	Destination string      `yaml:"destination"`
	Permissions os.FileMode `yaml:"permissions"`
}

// ImplementOptions is used to configure the client.
type ImplementOptions struct {
	Address    string
	Plan       string
	VaultToken string
}

// renderTemplates renders templates and writes them to disk.
func renderTemplates(templates map[string]*Template, opts *ImplementOptions) error {
	for template, stats := range templates {
		templateBytes, err := ioutil.ReadFile(path.Join("./templates", template))
		if err != nil {
			return err
		}

		// Set up a connection to the server.
		conn, err := grpc.Dial(opts.Address, grpc.WithInsecure())
		if err != nil {
			return err
		}
		defer func() {
			err = conn.Close()
			if err != nil {
				return
			}
		}()

		c := api.NewDeviseClient(conn)

		// Contact the server and get the rendered plan.
		r, err := c.Template(context.Background(), &api.TemplateRequest{Template: templateBytes, VaultToken: opts.VaultToken})
		if err != nil {
			return err
		}

		err = os.MkdirAll(path.Dir(stats.Destination), 0755)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(stats.Destination, r.Rendered, stats.Permissions)
		if err != nil {
			return err
		}
	}

	return nil
}

// executeScript executes the plan's script.
func executeScript(script *string) (err error) {
	if script != nil {
		command := exec.Command("bash", "-c", *script)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err = command.Start(); err != nil {
			return
		}
		if err = command.Wait(); err != nil {
			return
		}
	}

	return
}

// Implement executes the plan.
func Implement(opts *ImplementOptions) (err error) {
	absFilepath, err := filepath.Abs(opts.Plan)
	if err != nil {
		return
	}
	dir := filepath.Dir(absFilepath)
	if err = os.Chdir(dir); err != nil {
		return
	}
	data, err := ioutil.ReadFile(opts.Plan)
	if err != nil {
		return
	}
	plan := &Plan{}
	if err = yaml.Unmarshal(data, plan); err != nil {
		return
	}
	if err = renderTemplates(plan.Templates, opts); err != nil {
		return
	}
	if err = executeScript(plan.Script); err != nil {
		return
	}

	return
}
