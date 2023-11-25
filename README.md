# fold search engine
> part of assignment

## bootstrapping

### requirements

1. docker
2. docker-compose
3. make

### instructions

1. `make build` to all build local images
   1. `pgsync` an elt pipeline application in python
   2. `app` server application written in `golang`
2. `make up` to bootstrap all applications

## endpoints

### create project
```shell
curl -X POST --location "http://localhost:8080/projects" \
    -H "Content-Type: application/json" \
    -H "Accept: application/json" \
    -d "{
          \"name\": \"p1\",
          \"description\": \"awesome description\",
          \"users\": [
            \"john doe\",
            \"jane doe\"
          ],
          \"hashtags\": [\"h1\", \"h2\"]
        }"
```

### list project
```shell
curl 'http://localhost:8080/projects' \
  -H "Content-Type: application/json" \
  -H "Accept: application/json"
```

### update project
```shell
curl -X PATCH --location "http://localhost:8080/projects" \
    -H "Content-Type: application/json" \
    -H "Accept: application/json" \
    -d "{
          \"id\": 51,
          \"name\": \"p1\",
          \"description\": \"awesome description\",
          \"users\": [
            \"john doe\",
            \"jane doe\"
          ],
          \"hashtags\": [\"h1\", \"h2\"]
        }"
```

### search project
```shell
curl -X POST --location "http://localhost:8080/projects/search" \
    -H "Content-Type: application/json" \
    -H "Accept: application/json" \
    -d "{
          \"description\": \"dumy\",
          \"fuzziness\": 2
        }"
```

## explanation
1. `internal/api`: business logic
2. `internal/env`: env parsing
3. `internal/es`: managing elastic search client
4. `internal/models`: database models
5. `internal/routes`: configuring api (above) and http
6. `internal/utils`: helper functions
7. [pgsync](https://github.com/toluaina/pgsync/) reads from pg wal, 
   transforms it as prescribed in `pg/schema.json` and finally dumps it into elasticsearch
