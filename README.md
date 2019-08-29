## Golang Wasm Hello World
A hello world implementation of golang compiled to web assembly.

1. To build:
```
$ GOOS=js GOARCH=wasm go build -o public/main.wasm
```
2. Start the server that serves the public folder:

3. 
```
$ go run utilities/http-server.go
```

4. Open `localhost:12000/index.html` in your browser. 

With these steps in place, you should see "Hello from Web Assembly" in the developer console of your browser. If you do, you have compiled and ran your first go program in the browser. 