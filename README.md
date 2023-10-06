# Tofu Kubed

> A Kubernetes CRD and Controller to handle Tofu operations by generating k8s pods catered to perform Tofu workflows


## WIP

This repo is a fork of [terraform-operator](https://github.com/galleybytes/terraform-operator) and is still a work in progress

## What is tofu-kubed?

This project is:

- A way to run Tofu in Kubernetes by defining Tofu deployments as Kubernetes manifests
- A controller that configures and starts [Tofu Workflows](http://tf.galleybytes.com/docs/architecture/workflow/) when it sees changes to the Kubernetes manifest
- Workflow runner pods that execute Tofu plan/apply and other user-defined scripts

This project is not:

- A Tofu Module or Registry


## Join the Community

Currently, I'm experimenting with a Discord channel. It may be tough when taking into account juggling a full time job and full time parenting, but I'll see what comes of it. Join the channel https://discord.gg/J5vRmT2PWg

