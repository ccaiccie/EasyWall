package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

var isAuthenticated = false

// StartWebserver - Start the Listener for Pages, Assets and AJAX
func StartWebserver() {
	var err error
	var addr = ConfigGetValue("server", "address") + ":" + ConfigGetValue("server", "port")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", homeHandler)

	if ConfigGetValue("server", "ssl") == "true" {
		err = http.ListenAndServeTLS(addr, ConfigGetValue("server", "certfile"), ConfigGetValue("server", "keyfile"), nil)
	} else {
		err = http.ListenAndServe(addr, nil)
	}
	if err != nil {
		Error(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	result, _ := httputil.DumpRequest(r, true)
	Log(string(result))

	pusher, ok := w.(http.Pusher)
	if ok {
		err := pusher.Push("/static/stylesheet/bootstrap.min.css", nil)
		if err != nil {
			Error(err)
			return
		}
		pusher.Push("/static/images/favicon.ico", nil)
		pusher.Push("/static/stylesheet/font-awesome.min.css", nil)
		pusher.Push("/static/stylesheet/font-awesome.min.css", nil)
		pusher.Push("/static/stylesheet/easywall.min.css", nil)
		pusher.Push("/static/javascript/jquery.min.js", nil)
		pusher.Push("/static/javascript/popper.min.js", nil)
		pusher.Push("/static/javascript/bootstrap.min.js", nil)
		pusher.Push("/static/javascript/easywall.min.js", nil)
	}

	if r.URL.Path == "/" {
		io.WriteString(w, getTemplate("home"))
	} else {
		errorHandler(w, r, 404)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	template := getTemplate("error")
	template = strings.Replace(template, "[[status]]", strconv.Itoa(status), -1)

	io.WriteString(w, template)
}

func getTemplate(page string) string {
	assetPath := "./assets/html/"
	var html string

	content, err := ioutil.ReadFile(assetPath + "head.html")
	if err != nil {
		Error(err)
		return ""
	}
	s := string(content)
	html = html + s

	content, err = ioutil.ReadFile(assetPath + "nav.html")
	if err != nil {
		Error(err)
		return ""
	}
	s = string(content)
	html = html + s

	content, err = ioutil.ReadFile(assetPath + page + ".html")
	if err != nil {
		Error(err)
		return ""
	}
	s = string(content)
	html = html + s

	content, err = ioutil.ReadFile(assetPath + "footer.html")
	if err != nil {
		Error(err)
		return ""
	}
	s = string(content)
	return html + s
}
