# Kubernetes API Template & Practice

Steps to set up and deploy a Kubernetes cluster to Digital Ocean, starting with any containerized app.

0. Prerequisites:
    - Containerized target application that you wish to deploy
    - Knowledge:
        - Understanding of Docker and containerization
        - Familiarity with Kubernetes and `kubectl`
    - Accounts:
        - Digital Ocean
        - Dockerhub
    - Software:
        - Docker Desktop + Kubernetes
        - `kubectl`
    - Watch [this video](https://www.youtube.com/watch?v=PvfBCE-xgBY&t=119s) carefully

1. Create Kubernetes cluster in Digital Ocean. This can be done in two ways: 
    - From the command line using `doctl`
        - Install `doctl`
            - The video above, That Devops Guy does everything within the official `doctl` docker image, but I prefer to just [install `doctl` locally](https://docs.digitalocean.com/reference/doctl/how-to/install/)
                - If you're familiar with snap: `sudo snap install doctl`
        - Add PAT from dev dashboard (under Account -> API), then `doctl auth init`
        - Check out your deployment options in DO (check out both)
            - `doctl`
                - `doctl kubernetes options`
                - `doctl kubernetes options regions`
                - `doctl kubernetes options versions`
            - Dashboard
                - Availability matrix
                - Pricing
   - [From the dev dashboard](https://cloud.digitalocean.com/kubernetes/clusters)
        - This should honestly be pretty self-explanitory. DO has made it very easy to set things up.

2. Deploy target application to cluster
    - Create Dockerhub repository if you haven't yet
        - NOTE: private repositories require further setup than is required here
    - Build container image from target application, push to Dockerhub
        - `docker login --username=<docker_username>`
            - if you haven't logged into Dockerhub in your terminal
        - `docker images` 
            - this will show all of the images on your local machine
            - get the Image ID of the image you'd like to push
        - `docker tag <image_id> <docker_username>/<dockerhub_repo>/<image_name>:<version_tag>`
            - `docker tag` is similar to `git commit`
            - example: `docker tag 14a9f0ebbf02 devcliff/keeptorch/keeptorch:v0.0.1`
        - `docker push <docker_username>/<dockerhub_repo>`
    - Configure secrets, configmaps, services, ingresses, and deployments
        - *Secrets* may include certs, signing keys, api keys, 
        - *ConfigMaps* are _sort of_ like environment variables for a cluster
        - *Services* define how traffic is permitted to flow into and out of a cluster
        - *Ingresses* are designed to safely expose a cluster to the public
        - Examples can be found in this repo under `_kubernetes/base`
        - This is where actual Kubernetes work begins

3. Bonus: Configure DNS
    - steps may vary depending on where you purchased and how you configure your domain.
    - I have a domain registered through Amazon's Route53
        - `todo` - steps to hook up a route53 domain
