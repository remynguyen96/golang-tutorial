package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	// "github.com/julienschmidt/httprouter"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		_, _ = fmt.Fprint(w, "<h1>Welcome to mywebsites</h1>")
	} else if r.URL.Path == "/contact" {
		_, _ = fmt.Fprint(w, ""+
			"To get in touch, please send an email to "+
			"<a href=\"mailto:support@lensloc.com\">Support</a>."+
			"")
	} else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprint(w,
			"<h2>Sorry, We could not find the page "+
				r.URL.Path+" you were looking for :(</h2>"+
				"<p>Please email us if you keep being sent to an invalid page</p>.")
	}
}

/*func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hola, %s!\n", ps.ByName("name"))
}*/

func homepage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprint(w, "<h1>WelcomHelloe to mywebsites</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprint(w, ""+
		"To get in touch, please send an email to "+
		"<a href=\"mailto:support@lensloc.com\">Support</a>."+
		"")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprint(w,
		"<h2>Sorry, We could not find the page "+
			r.URL.Path+" you were looking for :(</h2>"+
			"<p>Please email us if you keep being sent to an invalid page</p>.")
}

var templates *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	// r.HandleFunc("/", homepage)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/contact", contact)
	// http.Handle("/", r)
	_ = http.ListenAndServe(":3000", r)

	// router := httprouter.New()
	// router.GET("/hello/:name/spanish", Hello)
	// _ = http.ListenAndServe(":3000", router)

	// mux := &http.ServeMux{}
	// mux.HandleFunc("/", handlerFunc)
	// _ = http.ListenAndServe(":3000", mux)

	// http.HandleFunc("/", handlerFunc)
	// _ = http.ListenAndServe(":3000", nil)
}
