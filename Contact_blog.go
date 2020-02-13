package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Name    string
	Email   string
	Subject string
	Message string
}
type Blog struct {
	Content string
	Body    string
}

//defining the template details
func main() {

	http.HandleFunc("/", FeedbackHandler)
	http.HandleFunc("/blog", BlogHandler)

	http.ListenAndServe(":8000", nil)

}

func BlogHandler(w http.ResponseWriter, r *http.Request) {

	blog :=[] Blog{
		{ Content: "This is the content of my blog", Body:"Body items"},
		{Content:"title", Body:"I am here"},
		{Content:"heyy" , Body:"sunshine"},
	}
	jsonContent, _ := json.Marshal(blog)

	fmt.Fprintf(w, "%s", jsonContent)

}
func FeedbackHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	if r.Method != http.MethodGet {
		tmpl.Execute(w, struct{ Success bool }{false})
		return

	}
	details := ContactDetails{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	//do something with the details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})
}
