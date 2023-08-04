# Messenger

A go implementation of the Hexagonal Architecture, taken from the article by: [Antony Shikubu](https://www.golinuxcloud.com/hexagonal-architectural-golang/)

## Environment Config
    export PGSQL_HOST="ipOrFqdn"
    export PGSQL_USER="dbuser"
    export PGSQL_PW="dbuserrpassword"
    export PGSQL_DB="dbname"

## Application Layout
```text
├── LICENSE
├── README.md
├── bin
│   └── messenger
├── go.mod
├── go.sum
├── internal
│   ├── adapters
│   │   ├── handler
│   │   │   └── http.go
│   │   └── repository
│   │       ├── postgres.go
│   │       └── redis.go
│   └── core
│       ├── domain
│       │   └── model.go
│       ├── ports
│       │   └── ports.go
│       └── services
│           └── services.go
└── main.go
```

