package main

import (
	"weather/common"
)

func main() {
	err := consume()
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
