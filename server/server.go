package main

import (
    "net/http"
    "crypto/sha512"
    "encoding/base64"
    "time"
    "log"
    "os"
    "os/signal"
    "context"
        )

func encode_password(password string) (string){
    // Hash password & encode
    hasher := sha512.New()
    bv := []byte(password) 
    hasher.Write(bv)
    sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
    return sha
}

func handler(w http.ResponseWriter, r *http.Request) {
    // Wait 5 seconds before responding
    time.Sleep(5 * time.Second)
 
    logger := log.New(os.Stdout, "", 0)
    r.ParseForm()
    password := r.Form.Get("password")
    password = encode_password(password)

    // Would want to store this in a database, but printing for the
    // purposes of this challenge. 
    logger.Println("Password encoded ", password)
        
}

func main() {

    logger := log.New(os.Stdout, "", 0)
 
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt)
    port := ":8080"

    // Create a new server
    s := &http.Server{
                        Addr:           port,
                        Handler:        http.HandlerFunc(handler),
                        }

    // Start a goroutine
    go func(){
        logger.Printf("Listening on port %s...", port)
        if err := s.ListenAndServe() ; err != nil {
            logger.Println("Shutting down server...")
        }
    }()

    <- stop

    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

    // Use Shutdown module for a graceful shutdown
    s.Shutdown(ctx)

    }
