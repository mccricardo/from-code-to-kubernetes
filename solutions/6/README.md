# Scale service to multiple replicas
```sh
# Delete old Pod
~ kubectl delete pod webservice

# Create Deployment and edit to add 
# imagePullSecrets:
#   - name: regcred
~ kubectl create deployment webservice --image=ghcr.io/mccricardo/webservice:0.1 --dry-run=client -o yaml > webservice_deployment.yaml

# Apply deployment
~ kubectl apply - f webservice_deployment.yaml

# Check deployment
~ kubectl get deployment

# Check ReplicaSet
~ kubectl get rs 

# Check Pods
~ kubectl get pod

# Scale service to multiple replicas
~ kubectl scale deployments webservice --replicas=2

# Check created Pods
~ kubectl get pod
```