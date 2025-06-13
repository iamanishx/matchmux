package authentication

import (
	"fmt"
	"io"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "auth")
	ctx , _ := io.ReadAll(r.Body)
	fmt.Println(string(ctx))
}
