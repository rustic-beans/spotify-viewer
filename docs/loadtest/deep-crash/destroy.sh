#!/usr/bin/env bash
set -e

# Configuration
IMAGE_NAME="graphql-websocket-client"

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${RED}GraphQL WebSocket Client Docker Cleanup${NC}"
echo "This script will:"
echo "1. Stop and remove all graphql-websocket-client containers"
echo "2. Optionally remove the Docker image"

# Stop and remove containers
echo -e "\n${YELLOW}Stopping and removing containers...${NC}"
CONTAINERS=$(docker ps -a -q --filter name=${IMAGE_NAME})
if [ -n "$CONTAINERS" ]; then
  docker stop $CONTAINERS
  docker rm $CONTAINERS
  echo -e "${GREEN}All ${IMAGE_NAME} containers removed.${NC}"
else
  echo -e "${YELLOW}No ${IMAGE_NAME} containers found.${NC}"
fi

# Ask if user wants to remove the Docker image
read -p "Do you want to remove the ${IMAGE_NAME} Docker image? (y/N): " REMOVE_IMAGE
if [[ "$REMOVE_IMAGE" =~ ^[Yy]$ ]]; then
  echo -e "\n${YELLOW}Removing Docker image ${IMAGE_NAME}...${NC}"
  docker rmi ${IMAGE_NAME}
  echo -e "${GREEN}Docker image removed.${NC}"
else
  echo -e "\n${YELLOW}Keeping Docker image ${IMAGE_NAME}.${NC}"
fi

echo -e "\n${GREEN}Cleanup complete!${NC}"
