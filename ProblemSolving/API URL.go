package Problemsolving

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/services", services)
	http.HandleFunc("/help", help)
	// http.HandleFunc("/overview", overview)

	// 2. Start the web server
	fmt.Println("Server running on http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page! ")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a simple Go web application.")
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact us at support@example.com.")
}

func services(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "We provide fast and reliable web services.")
}

func help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Need assistance? Check our FAQs here.")

	// func  overview(w http.ResponseWriter,r*http.Request){
	// 	fmt.Fprint(w," overview ")
	// }
}