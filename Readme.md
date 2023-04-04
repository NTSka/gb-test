# Run
``
go run cmd/calc/main.go
``

# Config
## client.eth.URI
URI of eth node (using dial by go-ethereum)

## calc.parallel
Count of parallel goroutines to regulate limits (using in errgroup)

## calc.count
Count of last blocks (100 by condition)