package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	// "strconv"
)

//Home handler function which serves the home page
func home(w http.ResponseWriter, r *http.Request) {
	//Check if URL path matches exactly
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
	}
	//If it's a POST request, parse and hold callsign
	if r.Method == http.MethodPost {
		//check for an error in parsing
		err := r.ParseForm()
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
		//hold the user name in a variable
		callsign := r.Form.Get("userName")
		//set cookie to hole callsign
		cookie := http.Cookie{
			Name:    "callsign",
			Value:   callsign,
			Expires: time.Now().AddDate(0, 0, 1),
			Path:    "/",
		}
		http.SetCookie(w, &cookie)
		//Redirect user to nav page
		http.Redirect(w, r, "/navigation", http.StatusSeeOther)

	}
	// //If it's a GET request, display "Enter username"
	// if r.Method == http.MethodGet {
	// }

	//Otherwise, throw a 405 error
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//Navigation handler function which serves the navigation page
func showNavigation(w http.ResponseWriter, r *http.Request) {
	//get info from cookie
	var cookie, err = r.Cookie("callsign")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error: Couldn't get callsign", 500)
		return
	}
	callsign := cookie.Value
	log.Println(callsign)
	// w.Write([]byte(callsign))

	//Check if URL path matches exactly
	if r.URL.Path != "/navigation" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/navigation.page.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.ExecuteTemplate(w, "navigation", callsign)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//Navigation handler function which serves the trade page
func showTrade(w http.ResponseWriter, r *http.Request) {
	//Check if URL path matches exactly
	if r.URL.Path != "/trade" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/trade.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

//Navigation handler function which serves the chat page
func showChat(w http.ResponseWriter, r *http.Request) {
	//Check if URL path matches exactly
	if r.URL.Path != "/chat" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/chat.page.tmpl",
	}

	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
