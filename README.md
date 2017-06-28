# JCCodeChallenge

## How to run:

From within project directory:
```bash
go build server.go && ./server
```
Optional port (default 8080)
```bash
./server -port=8888
```

Testing:
```bash
curl -d "password=angryMonkey" http://localhost:8080
```

 Ctrl-c waits for existing requests to complete before exiting and rejects new connections
