package turbocharger

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/engelsjk/colormap"
	plt "github.com/engelsjk/colormap/palette"
	"github.com/schollz/progressbar/v3"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

// TurboCharger ...
type TurboCharger struct {
	InputPath   string
	OutputPath  string
	Palette     plt.Palette
	PaletteName string
}

// New instantiates a new TurboCharger instance
func New() *TurboCharger {
	return &TurboCharger{}
}

// IO runs quality checks and processes the specified input filepath and output dirpath
func (t *TurboCharger) IO(input, output, palette string) error {

	var err error

	err = t.setInputFilepath(input)
	if err != nil {
		return err
	}

	err = t.setPalette(palette)
	if err != nil {
		return err
	}

	err = t.setOutputDir(output)
	if err != nil {
		return err
	}

	return nil
}

// Charge applies the turbo colormap styling to the image specified by an input filepath
func (t *TurboCharger) Charge() error {

	infile, err := os.Open(t.InputPath)
	if err != nil {
		return err
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		return err
	}

	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	tImg := image.NewRGBA(rect)

	bar := progressbar.Default(int64(size.X * size.Y))

	cm := colormap.Colormap{Palette: t.Palette}

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			grayPixel := img.At(x, y)
			p := color.GrayModel.Convert(grayPixel).(color.Gray).Y
			turboPixel := cm.ToRGBA(p, 255)
			tImg.Set(x, y, turboPixel)
			bar.Add(1)
		}
	}

	outfile, _ := os.Create(t.OutputPath)

	err = png.Encode(outfile, tImg)
	if err != nil {
		return err
	}

	fmt.Printf("turbocharged image created at '%s'\n", t.OutputPath)

	return nil
}

func (t *TurboCharger) ListPalettes() {
	p := t.palettes()
	for k := range p {
		fmt.Println(k)
	}
}

func (t *TurboCharger) palettes() map[string]plt.Palette {
	return map[string]plt.Palette{
		"cividis": plt.Cividis{},
		"crest":   plt.Crest{},
		"flare":   plt.Flare{},
		"icefire": plt.Icefire{},
		"inferno": plt.Inferno{},
		"magma":   plt.Magma{},
		"mako":    plt.Mako{},
		"plasma":  plt.Plasma{},
		"rocket":  plt.Rocket{},
		"turbo":   plt.Turbo{},
		"viridis": plt.Viridis{},
		"vlag":    plt.Vlag{},
	}
}

func (t *TurboCharger) setInputFilepath(input string) error {

	fileinfo, err := os.Stat(input)
	if os.IsNotExist(err) {
		return fmt.Errorf("input filepath '%s' does not exist", input)
	}
	if fileinfo.IsDir() {
		return fmt.Errorf("input filepath '%s' must be a file", input)
	}

	inExt := filepath.Ext(input)
	allowedExt := map[string]bool{".png": true, ".jpg": true, ".jpeg": true, ".webp": false}
	if supported, recognized := allowedExt[inExt]; recognized {
		if !supported {
			return fmt.Errorf("input extension '%s' not supported", inExt)
		}
	} else {
		return fmt.Errorf("input extension '%s' not recognized", inExt)
	}

	t.InputPath = input

	return nil
}

func (t *TurboCharger) setPalette(palette string) error {

	m := t.palettes()

	if _, ok := m[palette]; !ok {
		return fmt.Errorf("palette %s not recognized", palette)
	}

	t.Palette = m[palette]
	t.PaletteName = palette

	return nil
}

func (t *TurboCharger) setOutputDir(output string) error {

	dirinfo, err := os.Stat(output)
	if os.IsNotExist(err) {
		return fmt.Errorf("output dir '%s' does not exist", output)
	}
	if !dirinfo.IsDir() {
		return fmt.Errorf("output dir '%s' must be a folder", output)
	}

	inExt := filepath.Ext(t.InputPath)

	fn := strings.TrimSuffix(filepath.Base(t.InputPath), inExt)
	fp := filepath.Join(output, fn)
	outFilePath := fmt.Sprintf("%s-%s%s", fp, t.PaletteName, inExt)

	t.OutputPath = outFilePath

	return nil
}
