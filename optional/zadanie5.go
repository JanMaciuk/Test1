package optional

import (
	"fmt"
	"net/http"
)

//HelloWorldHTTP returns Hello World through http
func HelloWorldHTTP(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello World")
}
