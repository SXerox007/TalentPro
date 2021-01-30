# TalentPro
Assignment


### Run docker-compose 
// this will run the mongodb and go server
```
docker-compose build && docker-compose up
```

## How to run Docker file
```
//build docker file
docker build -t talentpro  -f Dockerfile .

//list of available image
docker image ls

//run docker
docker run -d -p 50051:50051 talentpro
```

## How to run Binary File
```
//go build 
go build -o talentpro Assignment-1/myapplication.go Assignment-1/talentpro.go

// run binary fike
./talentpro

```

## Go Server:
```
server run on localhost:50051
```

## Mongodb:
```
MongoDB run on mongodb://localhost:27017
```

