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
$ go get github.com/aws/aws-lambda-go/lambda   (onlu if you have not installed it)
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
- Nota: Para proyectos desde cero es mejor hacerlo con la lib de aws (https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)
pero en existentes y si impacta el desarrollo se pueden usar wrappers como: (para este proyecto se usó apexGateway, ya que contabamos con un desarrollo ya hecho y una arquitectura definida)
https://github.com/apex/gateway (https://www.ocelotconsulting.com/2019/02/25/the-right-abstraction-for-lambdas.html)

## AWS configuration:
- See "aws_conf_folder" where all process is documented via images



## Docker conf
# 1) Create docker file
# 2) build for create the image go-ms-client:1.0
sudo docker build --tag go-ms-client:1.0 .

# 3) run the image created in previous step

list stopped containers:
docker ps -a

copy the name or the container id of the container you want to attach to, and start the container with:
docker start -i <name/id>

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


con entrypoint 
# official image 
FROM golang:1.14.12-alpine3.11 AS build
# Set working directory---donde se incluirá el código de la aplicacion
WORKDIR /go/src/ms-client
COPY go.mod .
COPY go.sum .
RUN go mod download 
COPY . .
# compilar el app(go build main.go) dejando el binario (-o de out) en el directorio bin del GOPATH (/go/bin)
RUN go build -o /go/bin/ms-client main.go
# crear una nueva imagen mínima (scratch) 
FROM scratch
# haciendo que el entry poin sea el binario generadi 
COPY --from=build /go/bin/ms-client /go/bin/ms-client
ENTRYPOINT ["./ms-client"]

#CORRERLO 

>Construir la imagen basado en  e docker file 
sudo docker build -t ms-client-image .

Crea un container 
sudo docker run --name ms-client -p 8080:8000 ms-client-image
Crear un contenedor y borrarlo cuando se haga stop de éste 
sudo docker run --rm --name ms-client -p 8000:8080 ms-client-image  

:::::::::::::::::::..con esta configuracion funciona peero queda la imagen pesando 1GB :O::::::::::::::::
# official image 
FROM golang:latest
# Set working directory---donde se incluirá el código de la aplicacion
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download 
#copiar todo al wordir
COPY . .

ENV DATABASE_DRIVER=mongo
ENV DATABASE_CONN="mongodb+srv://xxx:xxxx@cluster-tulsoft.5r2lz.mongodb.net/test?retryWrites=true&w=majority"
ENV CLIENTAPI_SERVER_HOST=0.0.0.0
ENV CLIENTAPI_SERVER_PORT=8080
# compilar el app(go build main.go) dejando el binario (-o de out) en el directorio bin del GOPATH (/go/bin)
RUN go build
# crear una nueva imagen mínima (scratch) 
# haciendo que el entry poin sea el binario generadi 
CMD ["./ms-client"]
::::::::::::::::::::::::::::::::::::::::::::::::.