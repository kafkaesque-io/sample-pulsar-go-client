# Pulsar Go Sample Client

The sample is based on the [Pulsar Go client library's reference implmentation](https://pulsar.apache.org/docs/en/client-libraries-go/).

## Set up

Install the library package.
```bash
$ go get -u github.com/apache/pulsar/pulsar-client-go/pulsar
```

Replace correct pulsar URL, token and topic name in both consumer.go and producer.go files.

Start consumer
```bash
$ cd src
$ go run consumer.go
```

Start producer
```bash
$ go run producer.go
```
