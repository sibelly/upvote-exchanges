# Upvote Exchanges
API with Golang using gRPC with stream pipes that exposes an upvote service endpoints for exchanges, digital platforms that offers the purchase, sale and exchange of cryptocurrencies.

# Requirements
- [Golang 1.8](https://go.dev/)
- [gRPC](https://grpc.io/)
- [Docker](https://www.docker.com/) -> optional
- [Mongo](https://www.mongodb.com/) -> optional

## Usage

### Setting your environment variables

Copy and edit the `.env.example` to `.env`. Set the `MONGOURI` variable to your MongoDB uri.
The uri is going to be defined depending on the way that the mongodb runs on the machine, using docker, local installation or mongo atlas.

## Creating mongo container with docker
``` docker run -d -p 27017:27017 --name local-mongo mongo:latest ```

## Install go dependencies
```go get #add dependencies to current module and install them``` 

```go install #compile and install packages and dependencies```

```go mod tidy #add missing and remove unused modules```

## Generate .pb.go files from proto

```make gen```

## Start server

``` make server ```

## Start client

```make client```

## Testing the server with evans
More expressive universal gRPC client:

``` https://github.com/ktr0731/evans ```

# Improvements
- Add User and [middleware](https://grpc.io/blog/grpc-web-interceptor/) to authenticate and authorize using the *UnaryInterceptor* and *StreamInterceptor*;
- Improve the logic to seed the database to execute only once and not repeat the same values;
- Consume the [coin api](https://www.coinapi.io/) to seed the database;
- Deploy to a free cloud service;
- Improve tests.

# References

- https://github.com/junereycasuga/gokit-grpc-demo
- https://levelup.gitconnected.com/working-with-mongodb-using-golang-754ead0c10c
- http://gokit.io/
- https://github.dev/purwokertodev/go-ddd-grpc
- https://github.com/go-kit/examples
- https://github.dev/techschool/pcbook-go
- https://github.dev/kainbr3/klever.io_challenge
- https://github.dev/gabrielfvale/klever-grpc