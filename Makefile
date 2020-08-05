#!make
include .env

MONGO_DEV_CONFIG = {"_id":"rs0","members":[{"_id":0,"host":"mongo0:27017"},{"_id":1,"host":"mongo1:27017"},{"_id":2,"host":"mongo2:27017"}]}

all: mongo
	docker build -t medods-test .
	docker run --rm --name medods-auth --network app -p 0.0.0.0:3000:3000 medods-test

test: mongo
	go test github.com/Sippata/medods-test/test

mongo:
	docker build -t rs_configured_mongo - < Dockerfile-mongo
	docker-compose up -d