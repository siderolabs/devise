package cli

import (
	"bytes"
	"fmt"
	"runtime"
	"text/template"
)

var (
	// Tag is set at build time.
	Tag string
	// GitCommit is set at build time.
	GitCommit string
	// Built is set at build time.
	Built string
)

const versionTemplate = `Devise:
		Tag:         {{ .Tag | printf "%s" }}
		Git commit:  {{ .GitCommit | printf "%s" }}
		Built:       {{ .Built | printf "%s" }}
		Go version:  {{ .GoVersion | printf "%s" }}
		OS/Arch:     {{ .Os | printf "%s" }}/{{ .Arch | printf "%s" }}
`

// Version contains verbose version information.
type Version struct {
	Tag       string
	GitCommit string
	Built     string
	GoVersion string
	Os        string
	Arch      string
}

// PrintLongVersion prints verbose version information.
func PrintLongVersion() {
	v := Version{
		Tag:       Tag,
		GitCommit: GitCommit,
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

// PrintShortVersion prints the tag and git commit.
func PrintShortVersion() {
	fmt.Println(fmt.Sprintf("Devise %s-%s", Tag, GitCommit))
}
