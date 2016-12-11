package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"

	yaml "gopkg.in/yaml.v2"

	"github.com/autonomy/devise/discoverers"
	"github.com/autonomy/devise/modifiers"
	"github.com/urfave/cli"
)

type Devise struct {
	Modifiers   *Modifiers
	Discoverers *Discoverers
}

type Modifiers struct {
	Env *modifiers.Env
	K8s *modifiers.K8s
	AWS *modifiers.AWS
}

type Discoverers struct {
	DNS *discoverers.DNS
}

type Plan struct {
	Templates map[string]*Template `yaml:"templates"`
	Script    string               `yaml:"script"`
}

type Template struct {
	Destination string      `yaml:"destination"`
	Permissions os.FileMode `yaml:"permissions"`
}

func NewDevise(k bool) (*Devise, error) {
	env := modifiers.NewEnv()

	var k8s *modifiers.K8s
	if k {
		_k8s, err := modifiers.NewK8s()
		if err != nil {
			return nil, err
		}

		k8s = _k8s
	}

	aws := modifiers.NewAWS()

	dns := discoverers.NewDNS()

	modifiers := &Modifiers{
		Env: env,
		K8s: k8s,
		AWS: aws,
	}

	discoverers := &Discoverers{
		DNS: dns,
	}

	d := Devise{
		Modifiers:   modifiers,
		Discoverers: discoverers,
	}

	return &d, nil
}

func Run(d *Devise, p *Plan) error {
	command := exec.Command("bash", "-c", p.Script)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Start()
	err := command.Wait()
	if err != nil {
		return err
	}

	return nil
}

func MakeTemplates(d *Devise, p *Plan) error {
	for source, stats := range p.Templates {
		config, err := ioutil.ReadFile(path.Join("./templates", source))
		if err != nil {
			return err
		}

		var doc bytes.Buffer
		tmpl, err := template.New("").Funcs(template.FuncMap{
			"inc": func(i int) int {
				return i + 1
			},
			"hostname": func() (*string, error) {
				hostname, err := os.Hostname()
				if err != nil {
					return nil, err
				}
				return &hostname, nil
			},
			"contains": func(s string, substr string) bool {
				return strings.Contains(s, substr)
			},
			"not_contains": func(s string, substr string) bool {
				return !strings.Contains(s, substr)
			},
			"has_prefix": func(s string, prefix string) bool {
				return strings.HasPrefix(s, prefix)
			},
			"not_has_prefix": func(s string, prefix string) bool {
				return !strings.HasPrefix(s, prefix)
			},
		}).Parse(string(config))
		if err != nil {
			return err
		}

		err = tmpl.Execute(&doc, d)
		if err != nil {
			return err
		}

		err = os.MkdirAll(path.Dir(stats.Destination), 0755)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(stats.Destination, doc.Bytes(), stats.Permissions)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	app := cli.NewApp()

	app.Name = "devise"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "kubernetes",
			Usage: "enable kubernetes, defaults to false",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "implement",
			Aliases: []string{"i"},
			Usage:   "implement a plan",
			Action: func(c *cli.Context) error {
				devise, err := NewDevise(c.Bool("kubernetes"))
				if err != nil {
					return err
				}

				plan := &Plan{}
				data, err := ioutil.ReadFile("plan.yml")
				if err != nil {
					return err
				}

				err = yaml.Unmarshal([]byte(data), plan)
				if err != nil {
					return err
				}

				err = MakeTemplates(devise, plan)
				if err != nil {
					return err
				}

				err = Run(devise, plan)
				if err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:    "discover",
			Aliases: []string{"d"},
			Usage:   "discover a service",
			Subcommands: []cli.Command{
				{
					Name:  "dns",
					Usage: "discover a service via DNS",
					Action: func(c *cli.Context) error {
						dns := discoverers.DNS{}
						dns.Discover("kafka", "prod")
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
