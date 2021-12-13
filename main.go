package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func PrettyPrint(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(b)
	}
	return ""
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func debug(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	bytes, _ := httputil.DumpRequest(r, true)
	tmpl.Execute(w, map[string]interface{}{
		"HostDump": PrettyPrint(map[string]string{
			"Host": hostname,
			"Time": time.Now().String(),
		}),
		"EnvironemntDump": PrettyPrint(os.Environ()),
		"RequestDump":     string(bytes),
	})
}
func fetch(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(r.URL.Query().Get("url"))
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	fmt.Fprintf(w, "StatusCode: %d\n", resp.StatusCode)
	fmt.Fprintf(w, "Body: %s", string(body))
}

func main() {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/get", fetch)
	http.HandleFunc("/", debug)
	http.ListenAndServe(":"+port, nil)
}
