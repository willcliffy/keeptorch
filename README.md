# Kubernetes API Template & Practice

Steps to set up and deploy a Kubernetes cluster to Digital Ocean, starting with any containerized app.

0. **Prerequisites:**
    - Containerized target application that you wish to deploy
    - Knowledge:
        - General understanding of Docker and containerization
        - Familiarity with Kubernetes and `kubectl` a plus
    - Accounts:
        - Digital Ocean
        - Dockerhub
    - Software:
        - WSL2 (I'm using Ubuntun 20.04 here)
        - Docker Desktop + Kubernetes with WSL2 integration enabled
        - `kubectl`
    - Watch [this video](https://www.youtube.com/watch?v=PvfBCE-xgBY&t=119s) carefully

1. **Create Kubernetes cluster in Digital Ocean.** This can be done either [from the dev dashboard](https://cloud.digitalocean.com/kubernetes/clusters), or from the command line using `doctl`:
    - Install `doctl`
        - The video above, That Devops Guy does everything within the official `doctl` docker image, but I prefer to just [install `doctl` locally](https://docs.digitalocean.com/reference/doctl/how-to/install/)
            - If you're familiar with snap: `sudo snap install doctl`. Note that snap doesn't work in WSL
    - Add PAT from dev dashboard (under Account -> API), then `doctl auth init`
    - Check out your deployment options in DO (check out both)
        - `doctl`
            - `doctl kubernetes options`
            - `doctl kubernetes options regions`
            - `doctl kubernetes options sizes`
            - `doctl kubernetes options versions`
        - [Digital Ocean docs](https://docs.digitalocean.com/products/kubernetes/)
            - [Availability matrix](https://docs.digitalocean.com/products/platform/availability-matrix/)
            - [Pricing](https://docs.digitalocean.com/products/droplets/#plans-and-pricing)
    - `doctl kubernetes cluster create --help` 
        ```
        doctl kubernetes cluster create <cluster_name> \
            --version <k8s_version> \
            --count <initial_num_pods> \
            --size <pod_size> \
            --region <region_slug>
        ```
    - After the cluster finishes initializing, you may want to create a namespace - `kubectl create ns keeptorch`

2. **Deploy target application to cluster**
    - Create a public Dockerhub repository if you haven't yet. Note that configuring private Dockerhub repositories is beyond the scope of this repo.
    - Build container image from target application, push to Dockerhub. This can be done through Docker Desktop, or through Docker CLI:
        - `docker login --username=<docker_username>`
            - if you haven't logged into Dockerhub in your terminal
        - With your container running locally, `docker ps` to get the container ID
        - `docker commit <container_id> <dockerhub_username>/<dockerhub_repo>`
        - `docker push <docker_username>/<dockerhub_repo>`
    - Configure secrets, configmaps, services, ingresses, and deployments. Examples can be found in this repo under `_kubernetes/base`. This is where actual Kubernetes work begins
        - *Secrets* may include certs, signing keys, api keys, client secrets, etc...
        - *ConfigMaps* are sort of like environment variables for a cluster
        - *Services* define how traffic is permitted to flow into and out of a cluster
        - *Ingresses* are designed to safely expose a cluster to the public
        - *Deployments* define how to deploy and update the cluster, and how it should behave when deployed.
    - Deploy Dockerhub image to Kubernetes cluster
        - `kubectl apply -n keeptorch -f _kubernetes`
        - Deployment might take a few minutes. Use `kubectl get`/`kubectl describe` to begin debugging if there are issues.

3. **Next Steps**
    - Configure ingress/service. The `yml` files in `_kubernetes/base` will configure a basic Kubernetes Service of type LoadBalancer. Exposing an application through Services is acceptable for small projects and setup, but production-level code should be behind `Ingresses` which point to the underlying service. This is beyond the scope of this repo, and left for the reader to research.
    - Configure DNS. Steps may vary depending on where you purchased and how you configure your domain. Understanding how DNS works and how to configure it is beyond the scope of this repo.
