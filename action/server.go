package action

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/miclle/lisa/msg"
)

// Server : Serving Static Files with HTTP
func Server(addr, dir string) {

	msg.Info(fmt.Sprintf("Serving HTTP on 0.0.0.0 port %s ...", addr))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Print(os.Args[0])

		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		if _, err = os.Stat(dir + r.URL.Path); err != nil {
			// Return a 404 if the template doesn't exist
			if os.IsNotExist(err) {
				msg.Err(err.Error())
				http.NotFound(w, r)
				return
			}
			msg.Err(err.Error())
		}

		msg.Info(fmt.Sprintf("%s\t%s\t%s\t%s", r.RemoteAddr, r.Method, r.URL.Path, r.URL.RawQuery))
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
