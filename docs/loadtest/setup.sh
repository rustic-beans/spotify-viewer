#!/usr/bin/env bash
set -e

# Configuration
IMAGE_NAME="graphql-websocket-client"
IMAGE_TAG="latest"
FULL_IMAGE_NAME="${IMAGE_NAME}:${IMAGE_TAG}"
NUM_CONTAINERS=100
# GRAPHQL_ENDPOINT="ws://localhost:8080/query"
GRAPHQL_ENDPOINT="wss://spotify-viewer-app-7n6u4.ondigitalocean.app/query"

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}GraphQL WebSocket Client Docker Setup${NC}"
echo "This script will:"
echo "1. Build the Docker image"
echo "2. Start $NUM_CONTAINERS Docker containers with network access to the host"

# Check for Docker
echo -e "\n${YELLOW}Checking for Docker...${NC}"
if ! command -v docker &> /dev/null; then
  echo "Docker is required but not installed. Please install it and try again."
  exit 1
fi
echo "Docker found ✓"

# Build Docker image
echo -e "\n${YELLOW}Building Docker image ${FULL_IMAGE_NAME}...${NC}"
docker build -t ${FULL_IMAGE_NAME} .

# Start Docker containers
echo -e "\n${YELLOW}Starting $NUM_CONTAINERS Docker containers...${NC}"
for i in $(seq 1 $NUM_CONTAINERS); do
  CONTAINER_NAME="${IMAGE_NAME}-$i"
  echo "Starting container $i: $CONTAINER_NAME"
  
  docker run -d \
    --name $CONTAINER_NAME \
    --network=host \
    --restart=always \
    -e GRAPHQL_ENDPOINT=$GRAPHQL_ENDPOINT \
    ${NETWORK_OPT} \
    ${FULL_IMAGE_NAME}
done

echo -e "\n${GREEN}All containers started!${NC}"
echo -e "\nTo see logs from a specific container:"
echo "  docker logs -f ${IMAGE_NAME}-<number>"
echo -e "\nTo see logs from all containers:"
echo "  docker logs \$(docker ps -q --filter name=${IMAGE_NAME}) -f"
echo -e "\nTo stop and remove all containers:"
echo "  ./destroy.sh"
