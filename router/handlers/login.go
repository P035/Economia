package handlers

import (

  "fmt"
  "net/http"
  "encoding/json"
)

// This function will be the handler to all post requests to the /login route.
func Login(w http.ResponseWriter, r *http.Request) {

  fmt.Println("\nLogin request from: " + r.RemoteAddr)

  // Set headers.
  w.Header().Set("Content-Type", "application/json")

  // Create response.
  resp := response{

    Message: "Login!",
    Error: 0,
  }

  // Send response to client.
  err := json.NewEncoder(w).Encode(resp)
  if err != nil {

    error(w, err.Error())
  }
}
