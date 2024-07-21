package main

import (
	"weather/common"
)

func main() {
	err := consume()
	common.HandleError(err, "Failed to run consumer")
}
