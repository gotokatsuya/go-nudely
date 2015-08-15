package nudely

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"

	"github.com/disintegration/gift"
)

// DecodeImageByPath ...
func DecodeImageByPath(path string) image.Image {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	src, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return src
}

// DecodeImageByFile ...
func DecodeImageByFile(reader io.Reader) image.Image {
	src, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return src
}

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
			y, u, v := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			yCbCrs = append(yCbCrs, color.YCbCr{y, u, v})
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
func Detect(src image.Image) bool {

	dst := resizeImage(src, denominator)

	yCbCrs := image2YCbCrs(dst)
	sumTotalNude := countNude(yCbCrs)

	rating := float32(sumTotalNude) / float32(len(yCbCrs))
	fmt.Println(fmt.Sprintf("Rating : %f", rating))
	if rating > threshHold {
		fmt.Println("I think this is nude.")
		return true
	}

	fmt.Println("No nude.")
	return false
}
