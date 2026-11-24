package main
import("log";"net/http";"releasegate/api")
func main(){mux:=http.NewServeMux();mux.HandleFunc("/api/",api.Handler);mux.Handle("/",http.FileServer(http.Dir("public")));log.Println("ReleaseGate listening on :8080");log.Fatal(http.ListenAndServe(":8080",mux))}
