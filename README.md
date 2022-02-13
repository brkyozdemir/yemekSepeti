# yemekSepeti
golang in-memory

## To Build And Run
```
go build
./yemekSepeti
```

## After Build, to run the tests 
```
go test ./internal/handlers/commands
go test ./internal/handlers/queries
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
