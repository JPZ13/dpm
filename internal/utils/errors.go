package utils

import "log"

// HandleFatalError logs a fatal error if not nil
func HandleFatalError(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
