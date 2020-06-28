package apiserver

import "fmt"

// ErrChecker method for check error and show message
func ErrChecker(err error) {
	// If got error
	if err != nil {
		// Show error in logs
		fmt.Println(err)
		return
	}
}
