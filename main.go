package main

import ("fmt"; "net/http")
func health(w http.ResponseWriter,r *http.Request){w.Header().Set("Content-Type","application/json");fmt.Fprint(w,`{"status":"ok"}`)}
func main(){http.HandleFunc("/health",health);http.ListenAndServe(":8080",nil)}
