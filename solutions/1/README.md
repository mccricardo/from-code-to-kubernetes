# Build and Run the service

## How to run the application
```sh
~ go run main.go
```

## How to build and run the build service
```sh
~ go build -o server
~ ./server
```

## What endpoints are available to be called and tested
Available endpoints: /hello and /headers
```sh
~ curl localhost:8090/hello

~ curl localhost:8090/headers
```