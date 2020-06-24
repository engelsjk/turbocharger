package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/engelsjk/cturbo"
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
try "turbocharger --help" to learn more
`
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {

	inPtr := flag.String("in", ".", "input filepath")
	outPtr := flag.String("out", ".", "output dir")

	flag.Parse()

	inFilePath := *inPtr
	outFileDir := *outPtr

	if inFilePath == "." {
		fmt.Println(banner)
		return
	}

	fileinfo, err := os.Stat(inFilePath)
	if os.IsNotExist(err) {
		fmt.Printf("-in=%s does not exist\n", inFilePath)
		return
	}
	if fileinfo.IsDir() {
		fmt.Printf("-in=%s must be a file\n", inFilePath)
		return
	}

	dirinfo, err := os.Stat(outFileDir)
	if os.IsNotExist(err) {
		fmt.Printf("-out=%s does not exist\n", outFileDir)
		return
	}
	if !dirinfo.IsDir() {
		fmt.Printf("-out=%s must be a dir\n", outFileDir)
		return
	}

	inExt := filepath.Ext(inFilePath)
	allowedExt := map[string]bool{".png": true, ".jpg": true, ".jpeg": true, ".webp": false}
	if supported, recognized := allowedExt[inExt]; recognized {
		if !supported {
			fmt.Printf("input extension %s not supported\n", inExt)
			return
		}
	} else {
		fmt.Printf("input extension %s not recognized\n", inExt)
		return
	}

	fn := strings.TrimSuffix(filepath.Base(inFilePath), inExt)
	fp := filepath.Join(outFileDir, fn)
	outFilePath := fmt.Sprintf("%s-turbo%s", fp, inExt)

	///

	infile, err := os.Open(inFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		log.Fatal(err)
	}

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	tImg := image.NewRGBA(rect)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			grayPixel := img.At(x, y)
			p := color.GrayModel.Convert(grayPixel).(color.Gray).Y
			r, g, b, a := cturbo.Map(p, 255)
			turboPixel := color.RGBA{r, g, b, a}
			tImg.Set(x, y, turboPixel)
		}
	}

	outfile, _ := os.Create(outFilePath)

	err = png.Encode(outfile, tImg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("turbocharged image created at %s\n", outFilePath)
}
