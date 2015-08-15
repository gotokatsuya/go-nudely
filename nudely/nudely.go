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

func color2rgb(c color.Color) (r, g, b, a uint8) {
	r16, g16, b16, a16 := c.RGBA()
	return uint8(r16 >> 8), uint8(g16 >> 8), uint8(b16 >> 8), uint8(a16 >> 8)
}

// Ycbcr ...
type Ycbcr struct {
	y  float32
	cb float32
	cr float32
}

func rgb2ycbcr(r, g, b uint8) Ycbcr {
	fr := float32(r)
	fg := float32(g)
	fb := float32(b)
	return Ycbcr{16 + ((65.738*fr + 129.057*fg + 25.064*fb) / 256),
		128 + ((-37.945*fr - 74.494*fg + 112.439*fb) / 256),
		128 + ((112.439*fr - 94.154*fg - 18.285*fb) / 256)}
}

func image2ycbcr(src image.Image) (ycbcrs []Ycbcr) {
	srcBounds := src.Bounds()
	for i := 0; i < srcBounds.Max.Y; i++ {
		for j := 0; j < srcBounds.Max.X; j++ {
			r, g, b, _ := color2rgb(src.At(j, i))
			ycbcr := rgb2ycbcr(r, g, b)
			ycbcrs = append(ycbcrs, ycbcr)
		}
	}
	return ycbcrs
}

func countNude(ycbcrs []Ycbcr) (counter int) {
	for i := 0; i < len(ycbcrs); i++ {
		if 80 <= ycbcrs[i].cb && ycbcrs[i].cb <= 120 {
			if 133 <= ycbcrs[i].cr && ycbcrs[i].cr <= 173 {
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

	ycbcrs := image2ycbcr(dst)
	sumTotalNude := countNude(ycbcrs)

	rating := float32(sumTotalNude) / float32(len(ycbcrs))
	fmt.Println(fmt.Sprintf("Rating : %f", rating))
	if rating > threshHold {
		fmt.Println("I think this is nude.")
		return true
	}

	fmt.Println("No nude.")
	return false
}
