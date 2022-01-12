# Share a reusable definition of your service
```sh
# Boostrap chart template
~ helm create webservice

# Remove unnecessary resources. Necessary resources: Deployment and Service. Make them as reusable as possible
# Don't forget
# imagePullSecrets:
#   - name: regcred

# Apply chart 
~ helm upgrade -i webservice2 --set name=webservice2 webservice

# Check releases
~ helm ls

# Check Deployment
~ kubectl get deployment

# Check Service
~ kubectl get service

# Check Pods
~ kubectl get pods

# Test your service
~ kubectl port-forward service/webservice2 8090:8090

~ curl localhost:8090/hello

~ curl localhost:8090/headers
```
