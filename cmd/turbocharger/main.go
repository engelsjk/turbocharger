package main

import (
	"flag"
	"fmt"

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

	input := flag.String("i", "", "input filepath")
	output := flag.String("o", ".", "output dir")
	palette := flag.String("p", "turbo", "palette")
	list := flag.Bool("l", false, "list palettes")

	flag.Parse()

	if *input == "" && *output == "." && *palette == "turbo" && !*list {
		fmt.Println(banner)
		return
	}

	turbo := turbocharger.New()

	if *list {
		turbo.ListPalettes()
		return
	}

	err := turbo.IO(*input, *output, *palette)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = turbo.Charge()
	if err != nil {
		fmt.Println(err)
		return
	}
}
