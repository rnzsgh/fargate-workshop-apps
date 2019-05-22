#!/bin/bash

TAG=latest

REPOSITORY=rnzdocker1/fargate-intro-api

docker build --tag $REPOSITORY:$TAG .

docker push $REPOSITORY:$TAG
