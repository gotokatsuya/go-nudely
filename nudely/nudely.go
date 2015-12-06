package nudely

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/disintegration/gift"
)

// DecodeImageByPath ...
func DecodeImageByPath(path string) (image.Image, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	src, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return src, nil
}

// DecodeImageByFile ...
func DecodeImageByFile(reader io.Reader) (image.Image, error) {
	src, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return src, nil
}

// TODO Too min images should not resized
func resizeImage(src image.Image, n int) image.Image {
	srcBounds := src.Bounds()
	gi := gift.New(gift.Resize(srcBounds.Max.Y/n, srcBounds.Max.Y/n, gift.LanczosResampling))
	dst := image.NewRGBA(gi.Bounds(srcBounds))
	gi.Draw(dst, src)
	return dst
}

func image2YCbCrs(src image.Image) (yCbCrs []color.YCbCr) {
	srcBounds := src.Bounds()
	for i := 0; i < srcBounds.Max.Y; i++ {
		for j := 0; j < srcBounds.Max.X; j++ {
			r, g, b, _ := src.At(j, i).RGBA()
			y, cb, cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			yCbCrs = append(yCbCrs, color.YCbCr{y, cb, cr})
		}
	}
	return yCbCrs
}

const (
	nudelyMinCb = 80
	nudelyMaxCb = 120
	nudelyMinCr = 133
	nudelyMaxCr = 173
)

func countNude(yCbCrs []color.YCbCr) (counter int) {
	for i := 0; i < len(yCbCrs); i++ {
		if nudelyMinCb <= yCbCrs[i].Cb && yCbCrs[i].Cb <= nudelyMaxCb {
			if nudelyMinCr <= yCbCrs[i].Cr && yCbCrs[i].Cr <= nudelyMaxCr {
				counter++
			}
		}
	}
	return counter
}

const (
	denominator = 2
	threshHold  = 0.5
)

// Detect ...
func Detect(src image.Image) (bool, float32) {

	dst := resizeImage(src, denominator)

	yCbCrs := image2YCbCrs(dst)
	sumTotalNude := countNude(yCbCrs)

	rating := float32(sumTotalNude) / float32(len(yCbCrs))
	if rating > threshHold {
		return true, rating
	}

	return false, rating
}
