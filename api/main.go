package main

import (
	"event-driven/api/handler"
	"event-driven/api/service"
	"event-driven/api/usecase"
	"event-driven/api/websocket"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	jobService := service.NewJobService()
	jobUsecase := usecase.NewJobUsecase(jobService)
	jobHandler := handler.NewJobHandler(jobUsecase)
	jobStatusHandler := handler.NewJobStatusHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// CORS設定の追加
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}
	r.Use(cors.Handler(corsOptions))

	r.Route("/job", func(r chi.Router) {
		r.Post("/", jobHandler.HandleJob)
		r.Get("/status/{jobID}", jobHandler.HandleJobStatus)
		r.Post("/status/update", jobStatusHandler.UpdateJobStatus)
	})

	r.HandleFunc("/ws", websocket.HandleConnections)

	go websocket.HandleMessages()

	log.Println("API server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
