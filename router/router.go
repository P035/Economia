package router

import (

  "net/http"
  "./handlers"
)

func Router() *http.ServeMux{

  // Create serve mux.
  SM := http.NewServeMux()

  // Register all handlers to the serve mux.
  SM.HandleFunc("/", handlers.HomeTest)

  // Return SM.
  return SM
}
