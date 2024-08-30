#!/bin/bash

IMAGE_NAME=expense-management-user-service-${USER}
IMAGE_TAG=development
DOCKER_FILE=$1
CONTAINER_NAME=expense-management-user-service-backend-${USER}
DIR=$(pwd)