
# official image 
FROM golang:1.14.12-alpine3.11 AS build
# Set working directory---donde se incluirá el código de la aplicacion
WORKDIR /go/src/ms-client
COPY go.mod .
COPY go.sum .
RUN go mod download 
#copiar todo al workdir
COPY . .
# compilar el app actual(go build main.go) dejando el binario (-o de out) en el directorio bin del GOPATH (/go/bin)
# bRUN go build -o /go/bin/msclient 
RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/ms-client 

# crear una nueva imagen mínima (scratch) 
FROM scratch
# copiando el build que quedó en /go/bin/msclient a go/src/msclient
COPY --from=build /go/bin/ms-client /go/bin/ms-client

# haciendo que el entry poin sea el binario generado al ejecutar rgo build
# este es el comando que se ejecutará al incial el contenedory
CMD ["./go/bin/ms-client"]