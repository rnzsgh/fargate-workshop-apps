#!/bin/bash

TAG=latest

REPOSITORY=rnzdocker1/fargate-intro-web

docker build --tag $REPOSITORY:$TAG .

docker push $REPOSITORY:$TAG
