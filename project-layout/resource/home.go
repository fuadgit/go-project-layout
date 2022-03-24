package resource

import (
	"fmt"
	"net/http"
)

// HomeHandler defines a resource that renders the ??? home page. // TODO replace ??? by app name
func HomeHandler(homePage string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, homePage)
	}
}
