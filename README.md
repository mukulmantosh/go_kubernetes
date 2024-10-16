# Kubernetes with GoLand

### Prerequisites

Before starting up this project, make sure you have the necessary dependencies installed in your machine.

###  Installation

- [x] [Go](https://go.dev/) - Go is an open source programming language that makes it simple to build secure, scalable systems.

- [x] [Docker](https://www.docker.com/) - Docker helps developers bring their ideas to life by conquering the complexity of app development.

- [x] [PostgreSQL](https://www.postgresql.org/) - The World's Most Advanced Open Source Relational Database



#### Running Postgres Database

```bash
docker run --name goland-k8s-demo -p 5432:5432 -e POSTGRES_PASSWORD=mukul123 -d postgres
```


### Environment Variables

Be sure to place the `.env` file in the project root and update the information according to your settings. Refer to the example below.

```
DB_HOST=localhost
DB_USERNAME=**************
DB_PASSWORD=**************
DB_NAME=go_k8s
DB_PORT=5432
```