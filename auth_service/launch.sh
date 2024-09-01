#!/bin/bash
source config.sh

# Check if the container is running
if [ "$(docker ps -q -f name=$CONTAINER_NAME)" ]; then
    echo "Stopping and removing the existing container..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
fi

# Check if the container exists but is not running
if [ "$(docker ps -aq -f status=exited -f name=$CONTAINER_NAME)" ]; then
    echo "Removing the existing stopped container..."
    docker rm $CONTAINER_NAME
fi

# Run the container
docker run -it --network=ecofy-net -p 8081:8080 -p 50052:50051 --restart always \
--name $CONTAINER_NAME -v $DIR:/app \
--env-file .env $IMAGE_NAME:$IMAGE_TAG sh -c "cd cmd/user_service/ && bash"
