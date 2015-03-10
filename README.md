# Gopher Spree API

Experimental implementation of the Spree API on steroids.

## Pre-requisites

- An up and running spree app
- A go installation
- A memcached server
- Any database server
- An elasticsearch server

## Getting started

First of all install [gpm](https://github.com/pote/gpm), for osx:

    $ brew install gpm

For *nix:

    $ git clone https://github.com/pote/gpm.git && cd gpm
    $ git checkout v1.3.2 # You can ignore this part if you want to install HEAD.
    $ ./configure
    $ make install

In order to allow `gpm` to install our package dependencies
you will need a personal github token configured for your workstation,
if you already did it for another project that will do it, if not, just
follow [gpm instructions](https://github.com/pote/gpm#private-repos).

### Install memcached

To install memcached.

    $ brew install memcached

To start Memcached Server

    $ memcached -vv

### Install Elastic Search

Follow the instructions [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/_installation.html).

## Workflow

### Adding new package dependencies

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

### Installing package dependencies

Install dependencies by running:

    $ gpm install       # All defined dependencies in the Godeps file
    $ go install ./...  # All of this project's subpackages


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

# Cross Domain Policy
X_ORIGIN              = '*'                                                # X-Domain origin header
X_METHODS             = '*'                                                # X-Domain methods header

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

  You should create a test database. So the following commands should be executed in the spree project:

    $ RAILS_ENV=test rake db:create
    $ RAILS_ENV=test rake db:migrate

  Also you should index some product test data into Elastic. First, create the 'test' index:

    curl -XPOST 'http://localhost:9200/test?pretty'

  Then index some data. We provide an example in `test/products.json`:

    curl -XPOST 'localhost:9200/test/product/_bulk?pretty' --data-binary @/full_project_path/test/products.json

Each test should create and roll back the data used to tests.

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
