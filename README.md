# Upvote Exchanges
API with Golang using gRPC with stream pipes that exposes an upvote service endpoints for exchanges, digital platforms that offers the purchase, sale and exchange of cryptocurrencies.

## Usage

### Setting your environment variables

Copy and edit the `.env.example` to `.env`. Set the `MONGOURI` variable to your MongoDB uri.

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

# References

- https://github.com/junereycasuga/gokit-grpc-demo
- https://levelup.gitconnected.com/working-with-mongodb-using-golang-754ead0c10c
- http://gokit.io/
- https://github.dev/purwokertodev/go-ddd-grpc
- https://github.com/go-kit/examples
- https://github.dev/techschool/pcbook-go
- https://github.dev/kainbr3/klever.io_challenge
- https://github.dev/gabrielfvale/klever-grpc