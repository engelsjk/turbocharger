package main

import (
	"flag"

	"github.com/engelsjk/turbocharger"
)

const (
	name   = "turbocharger"
	banner = `┏━━━━┳┓╋┏┳━━━┳━━┓┏━━━┳━━━┳┓╋┏┳━━━┳━━━┳━━━┳━━━┳━━━┓
┃┏┓┏┓┃┃╋┃┃┏━┓┃┏┓┃┃┏━┓┃┏━┓┃┃╋┃┃┏━┓┃┏━┓┃┏━┓┃┏━━┫┏━┓┃
┗┛┃┃┗┫┃╋┃┃┗━┛┃┗┛┗┫┃╋┃┃┃╋┗┫┗━┛┃┃╋┃┃┗━┛┃┃╋┗┫┗━━┫┗━┛┃
╋╋┃┃╋┃┃╋┃┃┏┓┏┫┏━┓┃┃╋┃┃┃╋┏┫┏━┓┃┗━┛┃┏┓┏┫┃┏━┫┏━━┫┏┓┏┛
╋╋┃┃╋┃┗━┛┃┃┃┗┫┗━┛┃┗━┛┃┗━┛┃┃╋┃┃┏━┓┃┃┃┗┫┗┻━┃┗━━┫┃┃┗┓
╋╋┗┛╋┗━━━┻┛┗━┻━━━┻━━━┻━━━┻┛╋┗┻┛╋┗┻┛┗━┻━━━┻━━━┻┛┗━┛
turbocharging images with the turbo colormap
try "turbocharger --help" to learn more`
)

func main() {

	inPtr := flag.String("in", ".", "input filepath")
	outPtr := flag.String("out", ".", "output dir")

	flag.Parse()

	turbo := turbocharger.New(name, banner)
	turbo.IO(*inPtr, *outPtr)
	turbo.Charge()
}
