# Leverage Docker multi-stage builds to reduce image size

## Make image smaller, maintaining the same functionality.
```sh
# Check images
~ docker images

# Build new smaller image
~ docker build -t ghcr.io/mccricardo/webservice:0.1 -f Dockerfile-multistage .

# Check images
~ docker images

# Run new image
~ docker run --name server -p 8090:8090 ghcr.io/mccricardo/webservice:0.1

# Check that functionality is maintained
~ curl localhost:8090/hello

~ curl localhost:8090/headers

# Push new image
~ docker push ghcr.io/mccricardo/webservice:0.1
```