package handlers

// This file will contain all the structs definitions that the program will need so it caould run.

// The response struct will have a property called "Message" that is going to say things like command
// executed successfully, and another called the error will have like more than 0 if the program glitches.
type response struct {

  Message string `json:"message"`
  Error int `json:"error"`
}
