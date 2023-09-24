package cli

import (
	"fmt"
	j2g "github.com/AuroraV/json-to-go"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	tmpFilePrefix     = "tmp-file-g2j-"
	goFilePackageHead = "package main"

	exitCodeOK = iota
	exitCodeFalsyErr
	exitCodeFlagParseErr
	exitCodeCompileErr
	exitCodeNoValueErr
	exitCodeDefaultErr
)

type cli struct{}

type flagOpts struct {
	Format  bool   `short:"f" long:"format" description:"output format(by gofmt)"`
	Inline  bool   `short:"i" long:"inline" description:"output inline"`
	Stream  string `short:"s" long:"stream" description:"input stream"`
	Version bool   `short:"v" long:"version" description:"display version information"`
	Help    bool   `short:"h" long:"help" description:"display this help information"`
}

func (cli *cli) run(args []string) int {
	if err := cli.runInternal(args); err != nil {
		fmt.Println(err.Error())
		return exitCodeDefaultErr
	}
	return exitCodeOK
}

func (cli *cli) runInternal(args []string) (err error) {
	var opts flagOpts
	args, err = parseFlags(args, &opts)
	if err != nil {
		return err
	}
	meta := j2g.Parse(opts.Stream)
	result := j2g.GenerateStruct(meta, opts.Inline)

	if opts.Format {
		return gofmtOutput(result)
	}
	return normalOutput(result)
}

func genFile() string {
	return fmt.Sprintf("%s%d", tmpFilePrefix, time.Now().UnixNano())
}

func normalOutput(result []string) error {
	for _, item := range result {
		fmt.Println(item)
	}
	return nil
}

func gofmtOutput(result []string) error {
	f, err := os.CreateTemp("", genFile())
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()
	if len(result) > 0 {
		_, err := f.WriteString(goFilePackageHead + "\n")
		if err != nil {
			return err
		}
	}
	for _, item := range result {
		_, err := f.WriteString(item + "\n")
		if err != nil {
			return err
		}
	}
	fmt.Println(f.Name())
	cmd := exec.Command("gofmt", "-w", f.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprintf("gofmt output: %s", string(out)))
		return err
	}
	body, err := os.ReadFile(f.Name())
	if err != nil {
		return err
	}
	content, _ := strings.CutPrefix(string(body), goFilePackageHead)
	content, _ = strings.CutPrefix(content, "\n")
	fmt.Println(content)
	return nil
}
