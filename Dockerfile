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
