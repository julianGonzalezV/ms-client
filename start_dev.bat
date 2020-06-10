SET DATABASE_DRIVER="mongo"
SET DATABASE_CONN="mongodb://mongoadmin:admin@localhost:27017/ms_client?authSource=admin"
SET CLIENTAPI_SERVER_HOST="0.0.0.0"
SET CLIENTAPI_SERVER_PORT=8080
go run main.go
