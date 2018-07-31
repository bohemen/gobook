// Fetch a url and print it's body
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	app := "fetch"
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", app, err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", app, err)
			os.Exit(1)
		}
		resp.Body.Close()
	}

}
