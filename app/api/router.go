package router

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a message from server. Hello, %s!", r.URL.Path[1:])
}
