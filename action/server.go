package action

import (
	"fmt"
	"net/http"

	"github.com/miclle/lisa/msg"
)

// Server : Serving Static Files with HTTP
func Server(addr, dir string) {

	msg.Info(fmt.Sprintf("Serving HTTP on 0.0.0.0 port %s ...", addr))

	handler := http.FileServer(http.Dir(dir))

	if err := http.ListenAndServe(addr, handler); err != nil {
		panic(err)
	}
}
