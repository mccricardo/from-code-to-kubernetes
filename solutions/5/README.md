# Run service in Kubernetes

## Run the service in Kubernetes
```sh
~ kubectl run webservice --image=ghcr.io/mccricardo/webservice:0.1

~ kubectl get pods

~ kubectl describe pod webservice

~ kubectl delete pod webservice
```

Kubernetes is not able to pull the image because the repository is private and it has no access to it. You have to options: make the repository public or enable Kubernetes to access. You're going for the latter. For that we'll create a [Secret](https://kubernetes.io/docs/concepts/configuration/secret/) that can be used to pull the image.
```sh
~ kubectl create secret generic regcred \
    --from-file=.dockerconfigjson=/home/mccricardo/.docker/config.json \
    --type=kubernetes.io/dockerconfigjson

~ kubectl get secrets
```

You now need to configure our Pod to use that secret. For that you'll generate a base definition.
```sh
~ kubectl run webservice --image=ghcr.io/mccricardo/webservice:0.1 --dry-run=client -o yaml > webservice.yaml
```
And add:
```yaml
 imagePullSecrets:
  - name: regcred
```

And you can now relaunch the service.
```sh
~ kubectl apply -f webservice.yaml
~ kubectl get pods
```

## Ensure it provides the same fucntionality as before
To ensure the service provides the same functionality you start by looking at it's logs.
```sh
~ kubectl logs webservice
```

You can then port-forward to the service, meaning that you'll be able to send requests as if the service was running locally.
```sh
~ kubectl port-forward pod/webservice 8090:8090
~ curl localhost:8090/hello
~ curl localhost:8090/headers
```