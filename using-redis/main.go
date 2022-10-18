package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aseemwangoo/golang-programs/cache"
	"github.com/aseemwangoo/golang-programs/structs"
	"github.com/gorilla/mux"
)

func main() {
	ConnectDB()

	redis, err := cache.NewRedis()
	if err != nil {
		log.Fatalf("Could not initialize Redis client %s", err)
	}

	conn, err := PGXConnection()
	if err != nil {
		log.Fatalf("❌❌❌ newDB: %v", err)
		return
	}

	defer func() {
		_ = conn.Close(context.Background())
		err := db.Close()
		if err != nil {
			log.Fatalf("❌❌❌ Cannot close DB : %v", err)
		}
		log.Print("✅ ✅ DB Closed")
	}()

	router := mux.NewRouter()
	routes(router, redis)

	fmt.Println("Starting server :8081")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func routes(router *mux.Router, redis *cache.Client) {
	renderJSON := func(w http.ResponseWriter, val interface{}, statusCode int) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(val)
	}

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		user, err := GetDBUsers()
		if err != nil {
			renderJSON(w, &structs.Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		renderJSON(w, &user, http.StatusOK)
	})

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		val := redis.GetUser(id)

		if val != nil {
			val.Source = "cache"
			renderJSON(w, &val, http.StatusOK)
			return
		}

		user, err := GetUserByID(id)
		if err != nil {
			renderJSON(w, &structs.Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}

		redis.SetUser(id, user)

		user.Source = "API"
		renderJSON(w, &user, http.StatusOK)
	})
}
