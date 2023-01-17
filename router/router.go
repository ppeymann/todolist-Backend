package router

import (
	"todoList/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	//* ROUTING

	router.HandleFunc("/api/getTodo",controller.GetTodos).Methods("GET")
	router.HandleFunc("/api/createTodo",controller.CreateOneTodo).Methods("POST")
	router.HandleFunc("/api/updateTodo/{id}",controller.MarkTodo).Methods("PUT")
	router.HandleFunc("/api/deleteTodo/{id}",controller.DeleteTodo).Methods("DELETE")
	router.HandleFunc("/api/deleteTodos",controller.DeleteTodos).Methods("DELETE")

	return router
}