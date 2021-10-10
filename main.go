package main

import (

  "fmt"
  "net/http"
  "log"
  "./router"
)

func main() {

  fmt.Println("Api listening on port 3000")

  // Start listener on port 3000.
  log.Fatal(http.ListenAndServe("localhost:3000", router.Router()))
}
