# Leverage Docker multi-stage builds to reduce image size
At this point you have a functional container image that can run your application successfully. Looking at the image size it looks a bit big:
```sh
~ docker images
REPOSITORY                      TAG           IMAGE ID       CREATED        SIZE
ghcr.io/mccricardo/webservice   0.1           3611977bdf0d   2 days ago     321MB
golang                          1.17-alpine   d8bf44a3f6b4   4 weeks ago    315MB
```

Can you make it smaller, maintaining the same functionality?

### Resources
  - [Use multi-stage builds](https://docs.docker.com/develop/develop-images/multistage-build/)

