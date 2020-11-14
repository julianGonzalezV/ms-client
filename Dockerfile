
# official image 
FROM golang:1.14.12-alpine3.11 AS build
# Set working directory---donde se incluirá el código de la aplicacion
WORKDIR /go/src/ms-client
COPY . .
# compilar el app(go build main.go) dejando el binario (-o de out) en el directorio bin del GOPATH (/go/bin)
RUN go build -o /go/bin/ms-client main.go
# crear una nueva imagen mínima (scratch) 
FROM scratch
# haciendo que el entry poin sea el binario generadi 
COPY --from=build /go/bin/ms-client /go/bin/ms-client
ENTRYPOINT ["/go/bin/ms-client"]