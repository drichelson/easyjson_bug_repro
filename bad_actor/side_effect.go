package bad_actor

import (
	"fmt"
)

// Global vars and init functions are largely considered not ok in Go, but they are both totally legal and many of us
// early on used them out of convenience.
// This example mimics a library we were using that emitted a log line when initializing a global var:
// 	github.com/aws/aws-xray-sdk-go v1.0.0-rc.9
// The module has since been updated so this is no longer a problem but this remains an issue for easyjson
func init() {
	fmt.Println("Output from a dependent package's init function")
}
