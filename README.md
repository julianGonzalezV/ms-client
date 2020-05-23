# ms-client
Microservice for  managing of data related to clients


In your local host you have to create the next folders structrure:
xxworkspace
    bin
    src
    pkg

Then, create the GOPATH environment variable, point to xxworkspace folder
Additionally, you  have to create  GOPATH/bin into your PATH env variable (This is a pending step, please read https://golang.org/doc/gopath_code.html before )
    

# Herramientas 

# Referencia para la realización del presente ejemplo:
https://blog.friendsofgo.tech/posts/como_crear_una_api_rest_en_golang/ 


# Commands to execute if you want to run this project
-To install gorilla mux
$ go get -u github.com/gorilla/mux

# Install the MongoDB Go Driver
go get go.mongodb.org/mongo-driver


$ go run cmd/clientapi/main.go
