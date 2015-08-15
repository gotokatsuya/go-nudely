package main

import (
	"flag"
	"fmt"
	"image"

	"github.com/gotokatsuya/go-nudely/nudely"
)

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}
	path := flag.String("path", "", "A path of image will be read")
	flag.Parse()

	if *path == "" {
		fmt.Println("Specify a path of image, please")
		return
	}

	var src image.Image
	if src = nudely.DecodeImageByPath(*path); src == nil {
		return
	}
	nudely.Detect(src)
}
