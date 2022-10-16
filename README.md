# Simple Gift Card Management Service

This is a simple gift card management service.

## Structure
The Domain Driven Design pattern has been used for the implementation.

- __config package:__ Is where the configs are stored and loaded.
- __domain package:__ Is where the logic and repositories are stored.
- __helper package:__ Is some helper functions.
- __infrastructure package:__ Is to connect to the database and initialize the database driver.
- __mocks package:__ Is used for tests.

## How to Use
First you have to run the docker-compose file to run a local database.
```bash
docker-compose -f docker-compose.yaml up -d
```

Then you have to install the dependencies.
```bash
go mod vendor
```

Then you are good to go and just have to start the program.
```bash
go run main.go
```

## APIs
You can see the APIs in the postman link below:
```
https://www.getpostman.com/collections/6b55a51562af19c3e6fa
```

There are examples added for better understanding the APIs.