package main

import (
  "encoding/json"
  "log"
  "net/http"
)

type Comment struct {
  Comment string `json:"comment"`
}

var comments = make([]Comment, 0)

func AddComments(w http.ResponseWriter, r *http.Request)  {
  log.Println("add a comment")
  var comment Comment
  _ = json.NewDecoder(r.Body).Decode(&comment)
  comments = append(comments, comment)
  w.WriteHeader(200)
}


func GetComments(w http.ResponseWriter, _ *http.Request)  {
  w.WriteHeader(200)
  json.NewEncoder(w).Encode(comments)
}