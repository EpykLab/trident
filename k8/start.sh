#!/bin/bash

#colima start
minikube start --driver=podman
skaffold config set --global local-cluster true
eval $(minikube docker-env)
