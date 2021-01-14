package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Love is but a song to sing Fear's the way we die You can make the mountains ring Or make the angels cry"
	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}
