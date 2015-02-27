# Gopher Spree API

Experimental implementation of the Spree API on steroids.

## Prerequisites

- An up and running spree app
- A go installation

## Getting started

First of all install [gpm](https://github.com/pote/gpm),
[gvp](https://github.com/pote/gvp) and
[gpm-local](https://github.com/thecnosophos/gpm-local):

    $ brew install gpm
    $ brew install gvp
    $ brew install gpm-local # This is a gpm plugin

In order to allow `gpm` to install our package dependencies
you will need a personal github token configured for your workstation,
if you already did it for another project that will do it, if not, just
follow [gpm instructions](https://github.com/pote/gpm#private-repos).

## Workflow

### Always and for every shell session

Source gvp into your shell session:

    $ source gvp in

### Installing dependencies

Install dependencies by running:

    $ gpm install

### One time only

Tell gpm to link this project into the `.godeps` dir with `gpm-local`:

    $ gpm local name github.com/crowdint/gopher-spree-api

This will allow us to find this projects subpackages in the current
$GOPATH

### Adding new dependencies

Add dependencies by appending them to the Godeps file in the following
format:

    # Git repos
    github.com/nu7hatch/gotrail             v0.0.2
    github.com/replicon/fast-archiver       v1.02

    # Subpackages
    github.com/garyburd/redigo/redis        a6a0a737c00caf4d4c2bb589941ace0d688168bb

    # Bazaar Repo
    launchpad.net/gocheck                   r2013.03.03

    # Mercurial Repo
    code.google.com/p/go.example/hello/...  ae081cd1d6cc

And install dependencies again.

### Installign Memcached
To install memcached, you must have [port](http://www.macports.org/install.php) installed.

    $ sudo port install memcached

To start Memcached Server

    $ memcached -vv

### Configuration

There are sane defaults for running this app, but if you need to have
specific configurations, the following variables will override them:

```
# Database
DATABASE_DEBUG        = true                                               # Logging enabled
DATABASE_ENGINE       = postgres                                           # Engine name
DATABASE_URL          = dbname=spree_dev sslmode=disable                   # Connection string/URL
TEST_DATABASE_URL     = dbname=spree_test sslmode=disable                  # Test database connection string/URL
MAX_IDLE_CONNS        = 2                                                  # Max iddle connections
MAX_OPEN_CONNS        = 5                                                  # Max open connections


# New Relic
NEWRELIC_API_KEY      = ''                                                 # API Key
NEWRELIC_APP_NAME     = ''                                                 # Application name

# Spree
SPREE_NAMESPACE       = ''                                                 # Mounted at location (without slashes)
SPREE_URL             = http://localhost:5100                              # Host and port
SPREE_ASSET_PATH      = ":host/spree/products/:asset_id/:style/:filename"  # Assets default path
SPREE_ASSET_HOST      = ""                                                 # Assets host
SPREE_DEFAULT_STYLES  = "mini,small,product,large"                         # Assets default styles
```

## Testing

You should create a test database. So the following commands should be executed in the spree project:

    $ RAILS_ENV=test rake db:create
    $ RAILS_ENV=test rake db:migrate

Each test should create and roll back the data used to tests.

## Build

To build run the entire project:

    $ go build ./...

## Run

We use [forego](http://github.com/ddollar/forego) to run the app and
its services by simply running:

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
