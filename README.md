## Golang Wasm Hello World
A hello world implementation of golang compiled to web assembly.

#### To build:
1. Build the main program (ensure you've changed your directory to the project folder)
```
$ GOOS=js GOARCH=wasm go build -o public/main.wasm
```
2. Start the server serving index.html, main.wasm, and wasm_exec.js:
```
$ go run utilities/http-server.go
```
3. Open `localhost:12000/index.html` in your browser. 

With these steps in place, you should see "Hello from Web Assembly" in the developer console of your browser every second. Check out main.go to see how it works.