package helper

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"
)

var NetClient = &http.Client{
	Timeout: time.Second * 60,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
	CheckRedirect: func(r *http.Request, via []*http.Request) error {
		r.URL.Opaque = r.URL.Path
		return nil
	},
}

func PrintVars(w io.Writer, writePre bool, vars ...interface{}) {
	if writePre {
		io.WriteString(w, "<pre>\n")
	}
	for i, v := range vars {
		fmt.Fprintf(w, "» item %d type %T:\n", i, v)
		j, err := json.MarshalIndent(v, "", "    ")
		switch {
		case err != nil:
			fmt.Fprintf(w, "error: %v", err)
		case len(j) < 3: // {}, empty struct maybe or empty string, usually mean unexported struct fields
			w.Write([]byte(html.EscapeString(fmt.Sprintf("%+v", v))))
		default:
			w.Write(j)
		}
		w.Write([]byte("\n\n"))
	}
	if writePre {
		io.WriteString(w, "</pre>\n")
	}
}

func VarDump(vars ...interface{}) {
	PrintVars(os.Stdout, false, vars)
}
