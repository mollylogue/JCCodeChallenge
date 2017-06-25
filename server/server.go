package main

import (
    "fmt"
    "net/http"
    "crypto/sha512"
    "encoding/base64"
        )

func encode_password(password string) (string){
    hasher := sha512.New()
    bv := []byte(password) 
    hasher.Write(bv)
    sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
    return sha
}

func handler(w http.ResponseWriter, r *http.Request) {
        
        r.ParseForm()
        password := r.Form.Get("password")
        password = encode_password(password)

        fmt.Fprintf(w, "Password encoded %s", password)
        
}

func main() {

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
    
    }
