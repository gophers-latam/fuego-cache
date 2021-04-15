# fuego cache
Fuego cache is a concurrent hashed key-value pair written 100% in Golang. Could run in 3 modes:
 - TPC process
 - HTTP server
 - CLI

### Introduction
Just need to deploy the fuego instance locally or in a cloud provider and connect with a tcp client or as a web server.
Need different "modes" if you need a TCP plain connection or a web server, just add it in the json config file, located in the root.

### Installation
No further installation needed, just `go get github.com/tomiok/fuego-cache`

### Build
`make build`
  
### Run
`make run`

### Docker
#### Build CLI
`make docker-build-cli`

#### Build HTTP
`make docker-build-http`

#### Run CLI Mode
make docker-run-cli

#### Run HTTP Mode
make docker-run-http

**_NOTE:_**
Since configuration changes from a json file you need to build different images based on your configuration 

### TEST
this is a test
------------------------------------------------------------
EOY commit
