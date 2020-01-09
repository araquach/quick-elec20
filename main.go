package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/mailgun/mailgun-go/v3"
)

var (
	tplHome *template.Template
	tplAbout *template.Template
	tplDomestic *template.Template
	tplCommercial *template.Template
	tplTestimonials *template.Template
	tplContact *template.Template
)

type ContactMessage struct {
	Name 		string
	Email 		string
	Message 	string
}

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplHome.Execute(w, nil); err != nil {
		panic(err)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplAbout.Execute(w, nil); err != nil {
		panic(err)
	}
}

func domestic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplDomestic.Execute(w, nil); err != nil {
		panic(err)
	}
}

func commercial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplCommercial.Execute(w, nil); err != nil {
		panic(err)
	}
}

func testimonials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplTestimonials.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplContact.Execute(w, nil); err != nil {
		panic(err)
	}
}

func apiSendMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data ContactMessage
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"))

	sender := "info@basehairdressing.co.uk"
	subject := "New Message for Base"
	body := data.Message
	recipient := "adam@jakatasalon.co.uk"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	return
}

func main() {

	var err error

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplHome = template.Must(template.ParseFiles(
	"views/layouts/main.gohtml",
	"views/pages/index.gohtml"))
	if err != nil {
		panic(err)
	}

	tplAbout = template.Must(template.ParseFiles(
		"views/layouts/main.gohtml",
		"views/pages/about.gohtml"))
	if err != nil {
		panic(err)
	}

	tplDomestic = template.Must(template.ParseFiles(
		"views/layouts/main.gohtml",
		"views/pages/domestic.gohtml"))
	if err != nil {
		panic(err)
	}

	tplCommercial = template.Must(template.ParseFiles(
		"views/layouts/main.gohtml",
		"views/pages/commercial.gohtml"))
	if err != nil {
		panic(err)
	}

	tplTestimonials = template.Must(template.ParseFiles(
		"views/layouts/main.gohtml",
		"views/pages/testimonials.gohtml"))
	if err != nil {
		panic(err)
	}

	tplContact = template.Must(template.ParseFiles(
		"views/layouts/main.gohtml",
		"views/pages/contact.gohtml"))
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/about", about).Methods("GET")
	r.HandleFunc("/commercial", commercial).Methods("GET")
	r.HandleFunc("/domestic", domestic).Methods("GET")
	r.HandleFunc("/testimonials", testimonials).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	// api
	r.HandleFunc("/api/sendMessage", apiSendMessage).Methods("POST")

	// Styles
	assetHandler := http.FileServer(http.Dir("./dist/"))
	assetHandler = http.StripPrefix("/dist/", assetHandler)
	r.PathPrefix("/dist/").Handler(assetHandler)

	// JS
	jsHandler := http.FileServer(http.Dir("./dist/"))
	jsHandler = http.StripPrefix("/dist/", jsHandler)
	r.PathPrefix("/public/js/").Handler(jsHandler)

	//Images
	imageHandler := http.FileServer(http.Dir("./public/images/"))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imageHandler))

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":" + port, r)
}
