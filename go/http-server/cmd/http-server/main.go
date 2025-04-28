package main

import (
	"flag"
	"os"
	"path/filepath"

	http "github.com/EldoranDev/experiments/tree/main/go/http-server/internal"
)

var directoryFlag = flag.String("directory", ".", "directory that contains the files to serve")

func main() {
	flag.Parse()

	l := http.NewListener()
	defer l.Close()

	l.Add("GET", "/", func(req *http.Request, res *http.Response) {
		res.Status = http.StatusOk
	})

	l.Add("GET", "/echo/([a-zA-Z]+)", func(req *http.Request, res *http.Response) {
		res.Status = http.StatusOk
		res.Write([]byte(req.Params[1]))
		res.SetContentType("text/plain")
	})

	l.Add("GET", "/files/(.+)", func(req *http.Request, res *http.Response) {
		path := filepath.Join(*directoryFlag, req.Params[1])

		if _, err := os.Stat(path); err != nil {
			res.Status = http.StatusNotFound
			return
		}

		dat, err := os.ReadFile(path)
		if err != nil {
			res.Status = http.StatusInternalServerError
			return
		}

		res.Status = http.StatusOk
		res.SetContentType("application/octet-stream")
		res.Write(dat)
	})

	l.Add("POST", "/files/(.+)", func(req *http.Request, res *http.Response) {
		path := filepath.Join(*directoryFlag, req.Params[1])

		err := os.WriteFile(path, req.Body.Bytes(), 0644)

		if err != nil {
			res.Status = http.StatusInternalServerError
			return
		}

		res.Status = http.StatusCreated
	})

	l.Add("GET", "/user-agent", func(req *http.Request, res *http.Response) {
		ua := req.Header.Get("User-Agent")

		res.Status = http.StatusOk

		if ua != nil {
			res.SetContentType("text/plain")
			res.Write([]byte(*ua))
			return
		}
	})

	l.Listen()
}
