package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Isaias-Developer/cripto-api/cripto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

var r *chi.Mux

type Texto struct {
	Message string `json:"message"`
}

func InitServer(){
	r = chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    })

	r.Use(corsMiddleware.Handler)


	r.Post("/criptografar", func(w http.ResponseWriter, r *http.Request) {

		var mensagem Texto

		if err := json.NewDecoder(r.Body).Decode(&mensagem); err != nil {
			http.Error(w,err.Error(),500)
			return
		}

		w.Header().Set("Content-Type","application/json")

		response := map[string]interface{}{"texto_criptografado":cripto.Encrypt(mensagem.Message)}

		json.NewEncoder(w).Encode(response)
	})

	r.Post("/descriptografar",func(w http.ResponseWriter, r *http.Request) {
		var mensagem Texto

		if err := json.NewDecoder(r.Body).Decode(&mensagem); err != nil {
			http.Error(w,err.Error(),500)
			return
		}

		w.Header().Set("Content-Type","application/json")

		response := map[string]interface{}{"texto_descriptografado":cripto.Decrypt(mensagem.Message)}

		json.NewEncoder(w).Encode(response)
	})

	log.Fatal(http.ListenAndServe(":8080",r))
}