package action

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/miclle/lisa/msg"
	"github.com/skratchdot/open-golang/open"
)

type server struct {
	port                int
	bind, dir, absolute string
}

// Server : Serving Static Files with HTTP
func Server(port int, bind, dir string) {
	s := &server{
		bind: bind,
		port: port,
		dir:  dir,
	}

	if s.port == 0 {
		s.port = 8080
	}

	var err error
	s.absolute, err = filepath.Abs(dir)

	if err != nil {
		msg.Err(err.Error())
		return
	}

	http.HandleFunc("/", s.handleFunc)

	msg.Info("Serving HTTP on %s port %d ...", s.bind, s.port)

	// open URI using the OS's default browser
	if err := open.Run(fmt.Sprintf("http://%s:%d", s.bind, s.port)); err != nil {
		msg.Err(err.Error())
	}

	addr := fmt.Sprintf("%s:%d", s.bind, s.port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		msg.Err(err.Error())
		return
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
