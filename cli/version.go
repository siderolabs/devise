package cli

import (
	"bytes"
	"fmt"
	"runtime"
	"text/template"
)

var (
	// Release is set at build time.
	Release string
	// SHA is set at build time.
	SHA string
	// Built is set at build time.
	Built string
)

const versionTemplate = `Devise:
		Release:     {{ .Release }}
		SHA:         {{ .SHA }}
		Built:       {{ .Built }}
		Go version:  {{ .GoVersion }}
		OS/Arch:     {{ .Os }}/{{ .Arch }}
`

// Version contains verbose version information.
type Version struct {
	Release   string
	SHA       string
	Built     string
	GoVersion string
	Os        string
	Arch      string
}

// PrintLongVersion prints verbose version information.
func PrintLongVersion() {
	v := Version{
		Release:   Release,
		SHA:       SHA,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Built:     Built,
	}

	var wr bytes.Buffer
	tmpl, err := template.New("version").Parse(versionTemplate)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(&wr, v)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(wr.String())
}

// PrintShortVersion prints the release and sha.
func PrintShortVersion() {
	fmt.Println(fmt.Sprintf("Devise %s-%s", Release, SHA))
}
