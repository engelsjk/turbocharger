package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/engelsjk/cturbo"
	"github.com/schollz/progressbar/v3"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

// TurboCharger ...
type TurboCharger struct {
	InFilePath  string
	Name        string
	Banner      string
	OutFilePath string
	Error       bool
}

// New instantiates a new TurboCharger instance
func New(name, banner string) *TurboCharger {
	return &TurboCharger{
		Name:   name,
		Banner: banner,
	}
}

// IO runs quality checks and processes the specified input filepath and output dirpath
func (t *TurboCharger) IO(inFilePath, outFileDir string) {

	if inFilePath == "." {
		t.Error = true
		fmt.Println(t.Banner)
		return
	}

	fileinfo, err := os.Stat(inFilePath)
	if os.IsNotExist(err) {
		t.Error = true
		fmt.Printf("input filepath '%s' does not exist\n", inFilePath)
		return
	}
	if fileinfo.IsDir() {
		t.Error = true
		fmt.Printf("input filepath '%s' must be a file\n", inFilePath)
		return
	}

	inExt := filepath.Ext(inFilePath)
	allowedExt := map[string]bool{".png": true, ".jpg": true, ".jpeg": true, ".webp": false}
	if supported, recognized := allowedExt[inExt]; recognized {
		if !supported {
			t.Error = true
			fmt.Printf("input extension '%s' not supported\n", inExt)
			return
		}
	} else {
		t.Error = true
		fmt.Printf("input extension '%s' not recognized\n", inExt)
		return
	}

	t.InFilePath = inFilePath

	dirinfo, err := os.Stat(outFileDir)
	if os.IsNotExist(err) {
		t.Error = true
		fmt.Printf("output dir '%s' does not exist\n", outFileDir)
		return
	}
	if !dirinfo.IsDir() {
		t.Error = true
		fmt.Printf("output dir '%s' must be a folder\n", outFileDir)
		return
	}

	fn := strings.TrimSuffix(filepath.Base(inFilePath), inExt)
	fp := filepath.Join(outFileDir, fn)
	outFilePath := fmt.Sprintf("%s-turbo%s", fp, inExt)

	t.OutFilePath = outFilePath

	t.Error = false
}

// Charge applies the turbo colormap styling to the image specified by an input filepath
func (t *TurboCharger) Charge() {
	if t.Error {
		return
	}
	infile, err := os.Open(t.InFilePath)
	if err != nil {
		t.Error = true
		fmt.Println(err)
		return
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		t.Error = true
		fmt.Println(err)
		return
	}

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	tImg := image.NewRGBA(rect)

	bar := progressbar.Default(int64(size.X * size.Y))

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			grayPixel := img.At(x, y)
			p := color.GrayModel.Convert(grayPixel).(color.Gray).Y
			r, g, b, a := cturbo.Map(p, 255)
			turboPixel := color.RGBA{r, g, b, a}
			tImg.Set(x, y, turboPixel)
			bar.Add(1)
		}
	}

	outfile, _ := os.Create(t.OutFilePath)

	err = png.Encode(outfile, tImg)
	if err != nil {
		t.Error = true
		fmt.Println(err)
		return
	}

	fmt.Printf("turbocharged image created at '%s'\n", t.OutFilePath)
}
