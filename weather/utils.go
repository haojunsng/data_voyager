package main

import "log"

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
