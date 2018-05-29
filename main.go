package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const tpl = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>liveimg</title>
  <script src="http://localhost:35730/livereload.js"></script>
</head>
<body>
  <img src="%s" />
</body>
</html>`

func main() {
	addr := flag.String("addr", ":8080", "")
	flag.Parse()

	http.HandleFunc("/_liveimg/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, tpl, strings.TrimPrefix(r.URL.Path, "/_liveimg"))
	})
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Starting server on %s, preview path is `/_liveimg/:filename`", *addr)
	log.Println(http.ListenAndServe(*addr, nil))
}
