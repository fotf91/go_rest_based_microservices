# Run MySQL Database on Docker

- `cd <path>/resources/docker`
- `sudo docker-compose up` (sudo for linux)



This will start a container MySQL Database running on localhost:3306.

- DB_USER=root \
  DB_PASSWD=codecamp \
  DB_ADDR=localhost \
  DB_PORT=3306 \
  DB_NAME=banking \



Open a client (like DBeaver) and use credentials above.



# Project Structure

- `main.go`
- `app/app.go` ===> start server and define routes
- `app/handlers.go` ===> implementation of handling the service
- `domain/customer.go` ===> db data schema and interfaces of retrieving info from db
- `domain/CustomerRepositoryDb.go` ===> implementation of retrieving info from db
- `domain/CustomerRepositoryStub.go` ===> not used, is for mock data
- `service/customerService.go` ===> calls the database

