# Gopher Spree API

Experimental implementation of the Spree API on steroids.

## Getting started

### Prerequisites

- An up and running spree app
- A go installation

### Configuration

There are sane defaults for running this app, but if you need to have
specific configurations, the following variables will override them:

```
# Database
DATABASE_DEBUG  = true                             # Logging enabled
DATABASE_ENGINE = postgres                         # Engine name
DATABASE_URL    = dbname=spree_dev sslmode=disable # Connection string/URL
MAX_IDLE_CONNS  = 2                                # Max iddle connections
MAX_OPEN_CONNS  = 5                                # Max open connections

# Spree
SPREE_NAMESPACE = ''                               # Mounted at location
SPREE_URL       = http://localhost:5100            # Host and port
```

## Dependencies

To install project dependencies (packages):

    $ go get ./...

## Build

To build run the entire project:

    $ go build ./...

## Run

We use [forego](http://github.com/ddollar/forego) to run the app,
install it by simply running:

    $ go get github.com/ddollar/forego

Then, to run:

    $ forego start

**NOTE**: Use a custom `Procfile` or `Envfile` by specifying the
following flags:

    $ forego -f <My Procfile> -e <My Envfile>

## Tests

  We use the builtin `testing` package. To run the entire test suite:

    $ go test ./...

  To run some specific package tests:

    $ go test ./<package name>

  To avoid GIN logging in tests:

    $ GIN_MODE=test go test


## Guidelines and project structure

- We try to follow [the clean
architecture](http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html) please, raise your hand if you have doubts about it.
- Some helpful reading for applying it in go can be found
[here](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications).

## Contributing

1. Create your feature branch (`git checkout -b feature/my-new-feature`)
2. Commit your changes (`git commit -am 'Add some feature'`)
3. Push to the branch (`git push origin feature/my-new-feature`)
4. Create a new Pull Request
