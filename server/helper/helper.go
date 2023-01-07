package helper

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"os"
)

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
