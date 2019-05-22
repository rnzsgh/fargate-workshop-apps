package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/golang/glog"
)

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		log.Info("Backend called")

		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "Hello World - new!\n")

		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			io.WriteString(w, pair[0]+"="+pair[1]+"\n")
		}

	})
	http.ListenAndServe(":7080", nil)
}
