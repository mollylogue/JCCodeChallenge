package main

import (
    "fmt"
    "net/http"
    "crypto/sha512"
    "encoding/base64"
        )

func handler(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
        password := r.Form.Get("password")
        hasher := sha512.New()
        bv := []byte(password) 
        hasher.Write(bv)
        sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
        fmt.Fprintf(w, "Password normal %s", password)
        fmt.Fprintf(w, "Password encoded %s", sha)
        
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
    
    }
