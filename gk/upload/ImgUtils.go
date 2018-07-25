package upload

import (
	"image"
	"os"
	"github.com/hunterhug/go_image/graphics"
	"github.com/cihub/seelog"
	_ "image/png"
	_ "image/jpeg"
	"image/png"
	"image/jpeg"
	"strings"
)

func loadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

func ImgResize(sourceFile, outFile string, newWidth int) (int, int) {
	src, err := loadImage(sourceFile)
	if err != nil {
		seelog.Error("图片缩放失败:", sourceFile, " ==> ", outFile, " ~~ ", err)
		return 0, 0
	}
	bound := src.Bounds()
	dx := bound.Dx()
	dy := bound.Dy()
	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newWidth*dy/dx))
	err = graphics.Scale(dst, src)
	if err != nil {
		seelog.Error(err)
	}
	out, err := os.Create(outFile)
	if err != nil {
		seelog.Error(err)
	}
	defer out.Close()
	if strings.LastIndex(strings.ToLower(sourceFile), ".png") > 0 {
		png.Encode(out, dst)
	} else {
		jpeg.Encode(out, dst, nil)
	}
	return dx, dy
}
