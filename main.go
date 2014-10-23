package main

import (
  "time"
  "net/http"
)

func main() {
  p("Merkato", version(), "started at", config.Address)

  // handle static assets
  mux := http.NewServeMux()
  files := http.FileServer(http.Dir(config.Static))
  mux.Handle("/static/", http.StripPrefix("/static/", files))
  
  //
  // all route patterns matched here
  // route handler functions defined in other files
  //
  
  // defined in route_auth.go
  mux.HandleFunc("/", index)
  mux.HandleFunc("/login", login)
  mux.HandleFunc("/logout", logout)
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/signup_account", signupAccount)
  mux.HandleFunc("/authenticate", authenticate)



  // starting up the server
  server := &http.Server{
    Addr:           config.Address,
    Handler:        mux,
    ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
    WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
    MaxHeaderBytes: 1 << 20,
  }
  server.ListenAndServe()      
}