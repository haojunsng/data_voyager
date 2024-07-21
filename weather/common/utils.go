package common

import "log"

func HandleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
