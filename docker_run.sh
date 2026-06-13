#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

echo "===================================================="
echo "       Starting RepayCal Microservice System        "
echo "===================================================="

# 1. Environment Check: Verify if Docker and Docker Compose are installed
echo ""
echo "[1/4] Checking environment dependencies..."
if ! command -v docker &> /dev/null; then
    echo "X Error: Docker is not installed. Please download and install Docker Desktop first."
    exit 1
fi

if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
    echo "X Error: docker-compose is not installed. Please configure your Docker environment."
    exit 1
fi
echo "Ready: Docker environment check passed!"

# 2. Cleanup: Remove potential leftover containers from previous runs (Idempotency)
echo ""
echo "[2/4] Cleaning up existing container remnants..."
docker-compose down -v --remove-orphans > /dev/null 2>&1 || docker compose down -v --remove-orphans > /dev/null 2>&1
echo "Ready: Old environment cleaned up successfully."

# 3. Infrastructure: Pull images and start PostgreSQL & etcd
echo ""
echo "[3/4] Starting infrastructure components (PostgreSQL & etcd)..."
# Spin up postgres and etcd first to allow them to pull images and initialize
if command -v docker-compose &> /dev/null; then
    docker-compose up -d postgres etcd
else
    docker compose up -d postgres etcd
fi

# 4. Build & Run: Compile the Go microservice
echo ""
echo "[4/4] Compiling code and building Go RPC service image..."
if command -v docker-compose &> /dev/null; then
    docker-compose up --build -d calculator-rpc
else
    docker compose up --build -d calculator-rpc
fi

echo "===================================================="
echo "Success! All services are running in the background!"
echo "===================================================="

# Print current container running status
echo ""
echo "Current Container Status:"
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"

echo ""
echo "Quick Tips:"
echo "- Stream Go service real-time logs:  docker logs -f calculator_rpc"
echo "- please use 'docker exec -it postgres psql -U postgres -d loandb' & 'select * from pmt_log;' check DB data"
echo "- Stop and wipe all containers/data: docker-compose down -v"