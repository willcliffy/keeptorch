# Kubernetes API Template & Practice

Used while learning how to set up infrastructure and deploy a kubernetes cluster to different environments

- Local
- Dev (on-prem/manually configured server)
- Staging/Prod (hosting provider - most likely AWS because i bought a domain for keeptorch.com)

## General History

- Make basic golang API for testing'

- Containerize API
- Configure dockerhub, make repo, push image to repo
- Configure kubernetes and kubectl locally
- Set up DO kubernetes cluster
    - cli tool vs dashboard
    - configure cluster
- configure loadbalancer/server
- ingress in DO ([using helm](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-on-digitalocean-kubernetes-using-helm))
- configure domain/DNS

