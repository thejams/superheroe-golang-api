# superheroes-golang-api

a simple golang api that provides information about super heroes

# instalation with docker and docker compose

- make sure you have docker and docker-compose installed

- in the root of the proyect, execute docker-compose up and compose will carry all the hard work to activate the app

# instalation with minikube and kubernetes

- make sure you have minikube and kubernetes already installed

- execute the following steps

- create the docker image for the api with: docker build -t superheroe-golang-api -f Dockerfile .

- generate all the kubernetes resources running: kubectl apply -f k8s/

- the resources created by the previous command generate the following k8s resources: 

- 1 deployment with 1 POD and 1 replicaset.
- 1 service that allows the comunication with the POD.
- 1 configmap for passing enviornment variables. 
- 1 ingress to expose the APi outside the k8s cluster and a nginx-ingress-controller for supporting the ingress.

- run minikube service ingress-nginx-controller --url -n ingress-nginx to see the ip and port generated to interact with the API

