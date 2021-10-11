package router

import (

  "net/http"
  "./handlers"
)

func Router() *http.ServeMux{

  // Create serve mux.
  SM := http.NewServeMux()

  // Register all handlers to the serve mux.
  SM.HandleFunc("/", handlers.Routes["/"])
  SM.HandleFunc("/login", handlers.Routes["/login"])

  // Return SM.
  return SM
}
