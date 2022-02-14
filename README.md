![Go](https://img.shields.io/github/go-mod/go-version/vigo/statoo)

# yemekSepeti
golang in-memory

## To build and run
```
go build
./yemekSepeti
```

## After build, to run the tests 
```
go test ./internal/handlers/commands
go test ./internal/handlers/queries
```

## To run with dynamic interval
```
TIME_INTERVAL_IN_MINUTES=30 go run main.go
```

# API DOC
## Set Key
### Request
<code>POST /api/keys</code>
  
```json
{
    "testKey": "testValue"
}
```
### Response
```json
{
    "testKey": "testValue"
}
```

## Get Key
### Request
<code>GET /api/keys/:key</code>
### Response
```json
{
    "testKey": "testValue"
}
```

## Flush Memory
### Request
<code>DELETE /api/keys</code>
### Response
```json
{
    "message": "Store flushed successfully!"
}
```
## Get List (Not requested)
### Request
<code>GET /api/keys</code>
### Response
```json
[
    {
        "key0": "value0"
    },
    {
        "key1": "value1"
    }
]
```
# Heroku base url
[yemek-sepeti](https://yemek-sepeti-berkay-ozdemir.herokuapp.com)

# Docker
```
docker pull berkay94/yemek-sepeti:latest
```
