package cli

import (
	"fmt"
	j2g "github.com/AuroraV/json-to-go"
)

const (
	exitCodeOK = iota
	exitCodeFalsyErr
	exitCodeFlagParseErr
	exitCodeCompileErr
	exitCodeNoValueErr
	exitCodeDefaultErr
)

type cli struct{}

type flagOpts struct {
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
	for _, item := range result {
		fmt.Println(item)
	}
	return nil
}
