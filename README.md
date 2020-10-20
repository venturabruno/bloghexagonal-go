# BlogHexagonal in Go

## Build
`make build`

## Run
```
make dependencies
make up
make run
```

## Create post
```
curl -X "POST" "http://localhost:8080/v1/posts" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json' \
     -d $'{
	"title": "title",
	"subtitle": "subtitle",
	"content": "content"
}'
```

## Show posts
```
curl "http://localhost:8080/v1/posts" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

## Get post
```
curl "http://localhost:8080/v1/posts/78fe2a3e-fe4d-4d0b-90e8-82d3608cdc11" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```

## Publish post
```
curl -X "POST" "http://localhost:8080/v1/posts/78fe2a3e-fe4d-4d0b-90e8-82d3608cdc11/publish" \
     -H 'Content-Type: application/json' \
     -H 'Accept: application/json'
```