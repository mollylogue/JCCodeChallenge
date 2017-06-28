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
    "flag"
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

    // Make a channel that can send and receive signals
    signal_channel := make(chan os.Signal, 1)

    // Send only interrupt signals to the signal_channel when they occur
    signal.Notify(signal_channel, os.Interrupt)

    // Get port from command line arguments, default to 8080
    portPtr := flag.String("port", "8080", "a string")
    flag.Parse()
    logger.Println("port:", *portPtr)
    port := ":" + *portPtr

    // Create a new server
    s := &http.Server{
                        Addr:           port,
                        Handler:        http.HandlerFunc(handler),
                        }

    // Start a goroutine
    go func(){
        logger.Printf("Listening on port %s...", port)
        s.ListenAndServe()
    }()

    <- signal_channel

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Use Shutdown module for a graceful shutdown
    s.Shutdown(ctx)

    }
