package main

import (
	"flag"
	"fmt"

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

	img, err := nudely.DecodeImageByPath(*path)
	if err != nil {
		fmt.Println(err)
		return
	}

	detected, rating := nudely.Detect(img)
	fmt.Println(fmt.Sprintf("Rating : %f", rating))
	if detected {
		fmt.Println("I think this is nude.")
		return
	}

	fmt.Println("No nude.")
}
