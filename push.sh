#!/bin/bash

# Your actual repository URI is different
ECR_REGISTRY=public.ecr.aws/s5g3t0e9
ECR_REPOSITORY=janus-test/example-tool

aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws
docker build -t $ECR_REPOSITORY:latest . --platform linux/amd64 --pull
docker tag $ECR_REPOSITORY:latest $ECR_REGISTRY/$ECR_REPOSITORY:latest
docker push $ECR_REGISTRY/$ECR_REPOSITORY:latest