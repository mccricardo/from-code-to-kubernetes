# Expose service
```sh
# Create Service to expose webservice
~ kubectl expose deployment webservice --port=8090 --target-port=8090

# Check `Selector` and `Endpoints`
~ kubectl describe service webservice

# Check `Labels`
~ kubectl describe deployments webservice

# Check `Labels`
~ kubectl describe pod <pod-id>

# Get full Service definition
~ kubectl get service webservice -o yaml

# Test your service
~ kubectl port-forward service/webservice 8090:8090

~ curl localhost:8090/hello

~ curl localhost:8090/headers
```