# Build, Run and Publish containerized service

## Build a container
```sh
~ docker build -t ghcr.io/mccricardo/webservice:0.1 .
~ docker images
```

## Run the built container and ensure the same functionality as before
```sh
~ docker run --name server -p 8090:8090 ghcr.io/mccricardo/webservice:0.1
~ curl localhost:8090/hello
~ curl localhost:8090/headers
```

## How to publish the container Github container registry
Follow [Creating a personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) and [Working with the Container registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) to allow docker to publish images to the Github container image. To publish it:
```sh
~ docker push ghcr.io/mccricardo/webservice:0.1
```


