package handlers

import (

  "fmt"
  "net/http"
  "encoding/json"
)

// This function is just a test handler to check if the api is runing.
func HomeTest(w http.ResponseWriter, r *http.Request) {

  fmt.Println("Home test request from: " + r.RemoteAddr)

  // Create the json response.
  var resp response
  resp.Message = "Hello!"
  resp.Error = 0

  // Send response
  err := json.NewEncoder(w).Encode(resp)
  fmt.Println(err)
}
