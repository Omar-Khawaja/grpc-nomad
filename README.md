Testing load balancing gRPC services in Nomad with fabio. Job files are provided
if you want to run these in your own Nomad cluster.

### Deploying in Nomad

If you've deployed the services correctly, you should be able to hit your
gateway service and get results as show below:

```shell
$ curl greeting-gateway.service.consul:9090?name=Omar
Whats up Omar!
```

### Running locally

You can run these services and test locally as well. See below for examples.

Type the following to run the server:

```shell
$ go run greeting-server/server.go
```

In another window, type the following to run the client/gateway:

```shell
address=localhost:8080 go run greeting-gateway/client.go 
```

You should be able to run the following now:

```shell
$ curl localhost:9090/?name=omar
Whats up omar!
```
