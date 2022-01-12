# From Code to Kubernetes
This workshops aims to provided a high level overview about how you can go from have a service running locally to having it running in Kubernetes. It will introduce a few tools and practices that can be used in this journey.

The workshop will be a mix of [Learning-by-doing](https://en.wikipedia.org/wiki/Learning-by-doing) and [Project-based learning](https://en.wikipedia.org/wiki/Project-based_learning), two popular and effective learning methodologies. You'll progress through a series of [tasks](#tasks) that will highlight different tools and practices.

## Requirements
  - [Github account](https://github.com/)
  - [git](https://git-scm.com/)
  - [Go](https://go.dev/learn/)
  - [Docker](https://www.docker.com/get-started)
  - [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
  - [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
  - [Helm](https://helm.sh/docs/intro/install/)
  - [curl](https://curl.se/) (optional)

## Tasks
1. [Build and Run the service](tasks/1/README.md)
2. [Build, Run and Publish containerized service](tasks/2/README.md)
3. [Leverage Docker multi-stage builds to reduce image size](tasks/3/README.md)
4. [Launch Kubernetes cluster with kind](tasks/4/README.md)
5. [Run service in Kubernetes](tasks/5/README.md)
6. [Scale service to multiple replicas](tasks/6/README.md)
7. [Expose service](tasks/7/README.md)
8. [Share a reusable definition of your service](tasks/8/README.md)