package main

import (
	"db"
	"encoding/json"
	"log"
	"model/blog"
	"net/http"
	"service/blogService"
	"time"

	"github.com/gorilla/mux"
)

type Info struct {
	Message string `json:"message"`
}

func init() {
	log.Println("Connecting to database...")
	time.Sleep(2 * time.Second)
	db.ConnectDB()
	time.Sleep(1 * time.Second)
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

func createPost(response http.ResponseWriter, request *http.Request) {
	setupCorsResponse(&response, request)
	if (*request).Method == "OPTIONS" {
		return
	}
	var post blog.Posts
	_ = json.NewDecoder(request.Body).Decode(&post)
	var message = blogService.InsertPost(&post)
	json.NewEncoder(response).Encode(Info{message})
}

func getRecentPost(response http.ResponseWriter, request *http.Request) {
	setupCorsResponse(&response, request)
	if (*request).Method == "OPTIONS" {
		return
	}
	val, err := blogService.GetPost()
	if err != nil {
		log.Println(err)
		json.NewEncoder(response).Encode(err.Error())
		return

	}
	json.NewEncoder(response).Encode(&val)
}

func getAllPosts(response http.ResponseWriter, request *http.Request) {
	setupCorsResponse(&response, request)
	if (*request).Method == "OPTIONS" {
		return
	}
	val, err := blogService.GetAllPosts()
	if err != nil {
		log.Println(err)
		json.NewEncoder(response).Encode(err)
	}
	json.NewEncoder(response).Encode(&val)
}

func main() {
	log.Println("Starting server...")
	time.Sleep(2 * time.Second)
	router := mux.NewRouter()
	router.HandleFunc("/create-post", createPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/get-recent-post", getRecentPost).Methods("GET", "OPTIONS")
	router.HandleFunc("/get-all-post", getAllPosts).Methods("GET", "OPTIONS")
	log.Println("Server started successfully :)")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Failed to start server", err)
	}
}
