# Sample dev project

Simple project that takes in a stock ticker and date ranges and returns highest and lowest prices. Includes a Go server, a Python frontend, and a PostgreSQL database. Containerized these components, deploy them using Docker Compose and Kubernetes (Minikube), and ensured they can communicate with each other.

## Docker setup
### Building Docker Images:
Built Docker images for your Go server and Python frontend applications using docker build commands.

### Docker Compose:
Used a Docker Compose file to define services for your Go server, Python frontend, and PostgreSQL database. You defined a network for them to communicate.

### Docker Network:
Created a Docker network for your services to communicate within.

## Kubernetes (Minikube) Setup
### Minikube Installation:
Installed Minikube to create a local Kubernetes cluster.

### Deployment Manifests:
Created Kubernetes deployment manifest files (deployment.yaml) for Go server and Python frontend.

### Service Manifests:
Created Kubernetes service manifest files (service.yaml) for Go server and Python frontend.

### Applying Manifests:
You used kubectl apply to deploy Kubernetes resources.

### Communication Between Services:
After deploying to Kubernetes, frontend should send requests to the service names of your backend components.