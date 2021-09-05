# Kubernetes API Template & Practice

Learning how to set up infrastructure and deploy Kubernetes clusters to different environments.

- Local - Kubernetes cluster through Docker Desktop
- Prod - Digital Oceans Kubernetes cluster

## General History

- Make basic golang API for testing
    - See `main.go`. It's <75 lines and documented with psudocode if you dont know Golang
- Containerize API - see `Dockerfile` and `docker-compose.yml`
- Configure dockerhub, make repo, push image to repo
    - [Create dockerhub account](https://hub.docker.com/signup/)
    - Create repo
    - Create and tag image - TODO: document this process
    - `docker push <account name>/<repo name>:<tag name>`
- Configure kubernetes and kubectl locally
- Set up DO kubernetes cluster
    - cli tool vs dashboard
    - deployments, secrets, configs, services, ingresses
- ingress in DO ([using helm](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-on-digitalocean-kubernetes-using-helm))
- configure domain/DNS
