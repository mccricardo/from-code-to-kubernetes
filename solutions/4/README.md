# Launch Kubernetes cluster with kind

## Launch a local [Kubernetes](https://kubernetes.io/) cluster with [kind](https://kind.sigs.k8s.io)
```sh
# Create cluster
~ kind create cluster --config kind/config.yaml

# Ensure cluster was created
~ kind get clusters

# Check containers that kind launched
~ docker ps

# Use kubectl to check Kubernetes nodes
~ kubectl get nodes

# Use kubectl to check for Pods
~ kubectl get pods

# Use kubectl to check for all Pods
~ kubectl get pods -A
```