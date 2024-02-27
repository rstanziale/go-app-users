# GO app-users

This test application was created to explore the basic features by the Go programming language.

As a test scenario, a web server was implemented using the [gorilla/mux](https://github.com/gorilla/mux) package. 

## Configuration

There is a `.env` file where you can find the app properties.

## How to run

There are two ways to run the application:
* *Local*
* *Docker*

### Local

In this mode, you must have an active and running PostgreSQL instance.

To launch the application, run:

```bash
go run main.go
```

### Docker

In this mode, you simply use the command to create the docker image:

```bash
docker build -t app-users .
```

And finally run the command:

```bash
docker-compose up
```