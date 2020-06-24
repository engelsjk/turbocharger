package main

import (
	"flag"
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

	inPtr := flag.String("i", ".", "input filepath")
	outPtr := flag.String("o", ".", "output dir")

	flag.Parse()

	turbo := New(name, banner)
	turbo.IO(*inPtr, *outPtr)
	turbo.Charge()
}
