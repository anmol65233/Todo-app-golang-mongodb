package main

import (
	"anmol/todo/routehandlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(commonMiddleware)
	fmt.Println("👉🏻 server started at 3000")
	r.Get("/", routehandlers.GetHomefunc)
	r.Get("/get-tasks", routehandlers.GetAllTask)
	r.Post("/create-task", routehandlers.CreateTask)
	r.Get("/getone/{id}", routehandlers.GetOne)
	r.Post("/task-update/{id}/{status}", routehandlers.TaskUpdateStatus)
	r.Delete("/getone/{id}", routehandlers.DeleteOne)
	http.ListenAndServe(":3000", r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
