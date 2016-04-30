package action

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/miclle/lisa/msg"
)

type server struct {
	addr, dir, absolute string
}

// Server : Serving Static Files with HTTP
func Server(addr, dir string) {
	s := &server{
		addr: addr,
		dir:  dir,
	}

	var err error
	s.absolute, err = filepath.Abs(dir)

	if err != nil {
		msg.Err(err.Error())
	}

	http.HandleFunc("/", s.handleFunc)

	msg.Info(fmt.Sprintf("Serving HTTP on 0.0.0.0 port %s ...", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func (s *server) handleFunc(w http.ResponseWriter, r *http.Request) {
	dir := s.absolute + r.URL.Path

	if _, err := os.Stat(dir); err != nil {

		if os.IsNotExist(err) {
			msg.Err(s.requestInfo(r, 404))
			http.NotFound(w, r)
			return
		}

		msg.Err(s.requestInfo(r, 500))
		http.Error(w, err.Error(), 500)
		return
	}

	msg.Info(s.requestInfo(r, 200))
	http.ServeFile(w, r, dir)
}

func (s *server) requestInfo(r *http.Request, code int) string {
	return fmt.Sprintf("%s\t%s\t%d\t%s\t%s", r.RemoteAddr, r.Method, code, r.URL.Path, r.URL.RawQuery)
}
