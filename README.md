# ms-client
Microservice for  managing of data related to clients

# DDD architecture (previous lecture):
- http://dddsample.sourceforge.net/architecture.html
- https://www.youtube.com/watch?v=y3MWfPDmVqo&feature=youtu.be
- Example of use -> https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5


In your local host you have to create the next folders structrure:
xxworkspace
    bin
    src
    pkg

>Then, create the GOPATH environment variable, point to xxworkspace folder--ok

>Additionally, you  have to create  GOPATH/bin into your PATH env variable (This is a pending step, please read https://golang.org/doc/gopath_code.html before )
    

# Herramientas 
Golang Context:
https://blog.friendsofgo.tech/posts/context-en-golang/

# Referencia para la realización del presente ejemplo:
https://blog.friendsofgo.tech/posts/como_crear_una_api_rest_en_golang/ 


# Commands to execute if you want to run this project
-To install gorilla mux
$ go get -u github.com/gorilla/mux

# Install the MongoDB Go Driver
https://blog.friendsofgo.tech/posts/driver-oficial-mongodb-golang/
go get -u go.mongodb.org/mongo-driver

## Linux
```bash
$ source .env && go run main.go
```

## Windows
```bash
$ start.bat
```

# Get zip to AWS lambda: 
Doc
https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html

Process
- Into your root project folder execute:
1) 
```bash
$ go get github.com/aws/aws-lambda-go/lambda   (if required)
```
2) Next command create an executable file called main as the .go name file
```bash
$ GOOS=linux go build main.go
```

3) 
```bash
$ zip ms-client.zip main
```
4)Upload zip to S3 via aws cli or manually

5) Update url lambda s3 into lamda module 
- Nota: Para proyectos desde cero es mejor hacerlo con la lib de aws 
pero en existentes se pueden usar wrappers como: (para este proyecto se usó apexGateway, ya que contabamos con una arquitectura ya hecha)
https://github.com/apex/gateway (https://www.ocelotconsulting.com/2019/02/25/the-right-abstraction-for-lambdas.html)

## AWS configuration:
- See "aws_conf_folder" where all process is documented via images

## Services Payloads:

- Create - Post:
{
	"idType":"C", 
	"idNumber":"12345345", 
	"gender":"M", 
	"firstName":"julian", 
	"secondName":"andres", 
	"firstLastName":"gonzalez",
	"secondLastName":"velez", 
	"birthdate":"1987-05-20",
	"contact":{
		"email":"j.andres2087@gmail.com", 
		"cellphone":"3166351736", 
		"address":"cll 20 # 2-40", 
		"city":"valle", 
		"country":"Colombia"
	}
}
