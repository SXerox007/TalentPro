package main

import (
	"encoding/json"
	"log"
	"net/http"
	"talentpro/base/db/mongodb"
	db "talentpro/base/db/postgres"
	env "talentpro/base/environment"
	"talentpro/base/router"
	"talentpro/base/router/server"
)

// common response
type CommonResponse struct {
	Success bool        `param:"success" json:"success"`
	Message string      `param:"message" json:"message"`
	Code    int         `param:"code" json:"code"`
	Data    interface{} `param:"data" json:"data"`
}

// init
func Init() {
	environment := env.GetEnv()
	port := env.GetPort()
	//mongodbSetup()
	PgSetup()
	setupRouter(environment, port)

}

func main() {
	Init()
}

func PgSetup() {
	db.DBConnecting()
}

// mongodb setup
func mongodbSetup() {
	if err := mongodb.InitDB(); err != nil {
		log.Println("Error in Init MongoDB:", err)
		return
	}
}

func setupRouter(env, port string) {
	//initilize the router
	router.InitRouter()
	talentpro := router.SubRouter("/talentpro")
	talentpro.HandleFunc("/{version}/word-counter", GetWordCounterPage()).Methods("GET")
	talentpro.HandleFunc("/{version}/url-crawl",GetUrlCrawl()).Methods("GET")
	talentpro.HandleFunc("/{version}/prime-number", GetPrimeNumber()).Methods("GET")
	talentpro.HandleFunc("/{version}/last-day", GetLastDayOfMonth()).Methods("POST")

	user := router.SubRouter("/talentpro/{version}/user")
	user.HandleFunc("/list", GetUserList()).Methods("GET")
	user.HandleFunc("/add", AddUser()).Methods("POST")
	user.HandleFunc("/delete", DeleteUser()).Methods("DELETE")
	user.HandleFunc("/edit", EditUserDetails()).Methods("PATCH")
	

	log.Println("Server serve at: ", "localhost:"+port)
	server.StartServer(port)
}

// respondWithJSON - Serializes the payload to JSON and writes to ResponseWriter.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-XSS-Protection", "1; mode=block")
	w.WriteHeader(code)
	w.Write(response)
}
