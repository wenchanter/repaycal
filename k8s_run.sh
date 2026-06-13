#!/bin/bash

# Exit immediately if any command fails
set -e

echo "===================================================="
echo "      🚀 Starting RepayCal System on K8S 🚀         "
echo "===================================================="

# 1. Check for kubectl connectivity
echo ""
echo "[1/4] Checking Kubernetes connectivity..."
if ! docker info &> /dev/null; then
    echo "❌ Error: Docker daemon is not running!"
    echo "----------------------------------------------------"
    echo "💡 How to fix this:"
    echo "   1. Please open Docker Desktop (Mac/Windows) or start Docker service (Linux)."
    echo "   2. Wait until the Docker icon turns green (Running)."
    echo "   3. Re-run this script: ./k8s_run.sh"
    echo "----------------------------------------------------"
    exit 1
fi

if ! kubectl cluster-info &> /dev/null; then
    echo "X Error: Cannot connect to Kubernetes. Please ensure Docker Desktop/Minikube is running."
    exit 1
fi
echo "Ready: Connected to Kubernetes cluster."

# 2. Cleanup: Remove old resources to ensure a clean state
echo ""
echo "[2/4] Cleaning up existing Kubernetes resources..."
# Ignore errors if resources don't exist
kubectl delete -f k8s/ --ignore-not-found=true
echo "Ready: Old resources cleaned up."

# 3. Build Image locally (Shared with K8S in Docker Desktop/Minikube)
echo ""
echo "[3/4] Building Go RPC service image locally..."
docker build -t repaycal-calculator-rpc:latest .
echo "Ready: Image built locally."

# 4. Deploy all configurations
echo ""
echo "[4/4] Deploying to Kubernetes cluster..."
kubectl apply -f k8s/

echo "===================================================="
echo "Success! All services are deploying to K8S..."
echo "===================================================="

echo ""
echo "Monitoring Pod initialization (Press Ctrl+C to exit monitor):"
# Monitor all pods in the default namespace
kubectl get pods -w