package cli

import "os"

func Run() int {
	return (&cli{}).run(os.Args[1:])
}
