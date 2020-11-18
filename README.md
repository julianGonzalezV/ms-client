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
    eg
	GOPATH="/home/jag/Documents/study/goWorkspace"
	GOROOT="/usr/local/go"

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


## Upgrade Golang version:
- If you have a previous version of Go installed, be sure to remove it before installing another.



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
Crear un contenedor y que se borre automáticamente cuando se haga stop de éste 
sudo docker run --rm --name ms-client -p 8000:8080 ms-client-image  

::::::::..con esta configuracion funciona, pero queda la imagen pesando 705mb:: :
# official image 
FROM golang:1.15.5-alpine3.12
# Set working directory---donde se incluirá el código de la aplicacion
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download 
#copiar todo al wordir
COPY . .
# compilar el app(go build main.go) dejando el binario (-o de out) en el directorio bin del GOPATH (/go/bin)
RUN go build

RUN ls
# crear una nueva imagen mínima (scratch) 
# haciendo que el entry poin sea el binario generadi 
CMD ["./ms-client"]
::::::::::::::::::::::::::::::::::::::::::::::::.


::::..Con esta CONF arranca el server pero no reconoce el string de conexión mongo::::::.
# official image 
FROM golang:1.15.5-alpine3.12 AS build
# Set working directory---donde se incluirá el código de la aplicacion
WORKDIR /go/src/ms-client
COPY go.mod .
COPY go.sum .
RUN go mod download 
#copiar todo al workdir
COPY . .
# compilar el app actual(go build main.go) dejando el binario (-o de out) en el directorio bin del GOPATH (/go/bin)
# bRUN go build -o /go/bin/msclient 
RUN  CGO_ENABLED=0  go build -o /go/bin/ms-client 

# crear una nueva imagen mínima (scratch) 
FROM scratch
# copiando el build que quedó en /go/bin/msclient a go/src/msclient
COPY --from=build /go/bin/ms-client /go/bin/ms-client

# haciendo que el entry point sea el binario generado al ejecutar rgo build
# este es el comando que se ejecutará al incial el contenedory
ENTRYPOINT ["/go/bin/ms-client"]
::::::::::::::::::::::::::::::::::::::::::::::::.

:::::::::::::finalmente la que funciona es (ejemplo ):::::::::::::...
author https://github.com/jeremyhuiskamp/golang-docker-scratch (mejorando el ejemplo en mi repo a solo tener lo que se requiere)
::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
FROM golang:1.15.5-alpine3.12 as build
WORKDIR /go/src/ms-client
COPY . .
# Static build required so that we can safely copy the binary over.
# CGO_ENABLED permite la interoperatividad de Go programs con C
RUN CGO_ENABLED=0 go build -o /go/bin/msclient 

FROM golang:1.15.5-alpine3.12 as alpine
# --no-cache permitr no cachear el index localmente(docker container), 
# ayudando a que el contenedor sea lo más pequeño posible 
# Es igual que colocar  apk update AL INICIO y  rm -rf /var/cache/apk/* AL FINAL
RUN apk --no-cache add ca-certificates

#Creando imagen desde cero para que no pese ese voleo de MB o GB
FROM scratch
# copiando el build generado 
COPY --from=build /go/bin/msclient /go/bin/msclient
# agregadndo los Certificados  tls para conexión al exterior(eg https)
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/go/bin/msclient"]
