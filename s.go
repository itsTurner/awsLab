package main
import ."net/http"
func main() { HandleFunc("/", func(w ResponseWriter, r *Request) { ServeFile(w, r, "/home/ubuntu/deploy/i.html") }); ListenAndServe(":80", nil) }
