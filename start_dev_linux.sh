#!/bin/bash
echo "::::::init:::::::"
export DATABASE_DRIVER=mongo
export DATABASE_CONN="mongodb+srv://juligove:PASSWPRD@cluster-XXX.5r2lz.mongodb.net/test?retryWrites=true&w=majority"
export CLIENTAPI_SERVER_HOST=0.0.0.0
export CLIENTAPI_SERVER_PORT=8080
go run main.go
echo "::::::end:::::::"
