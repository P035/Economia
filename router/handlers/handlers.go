// This package is going to have all the handler functions of the server routes.
package handlers

import (

  "fmt"
  "net/http"
  "encoding/json"
)

// This function is just a test handler to check if the api is runing.
func HomeTest(w http.ResponseWriter, r *http.Request) {

  fmt.Println("Home test request from: " + r.RemoteAddr)

  // Set response headers.
  w.Header().Set("Content-Type", "application/json")

  // Create the json response.
  var resp response
  resp.Message = "Hello!"
  resp.Error = 0

  // Send response
  err := json.NewEncoder(w).Encode(resp)
  fmt.Println(err)
}

// This map will contain all the handler functions and whith what route are they associated.
var Routes map[string]func(w http.ResponseWriter, r *http.Request) = map[string]func(w http.ResponseWriter, r *http.Request){

  "/": HomeTest,
  "/login": Login,
}
