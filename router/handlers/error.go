package handlers

import (

  "fmt"
  "net/http"
  "encoding/json"
)

// When called this function sends an error response to the client.
func error(w http.ResponseWriter, err string) {

  fmt.Println("Error: " + "peo")

  // Set response status code to and internal server error.
  w.WriteHeader(http.StatusInternalServerError)

  // Create response.
  resp := response{

    Message: err,
    Error: 1,
  }

  // Send response.
  json.NewEncoder(w).Encode(resp)
}
