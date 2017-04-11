package main

import (
	"fmt"

	"github.com/ivancduran/edgecast/live"
)

func main() {

	stream := live.New("hls")
	id := stream.Create()
	fmt.Println(id)
	fmt.Println(id.HLS)

}
