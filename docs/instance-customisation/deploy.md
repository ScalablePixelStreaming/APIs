## Deploying to the cloud

Once you have built your instance customisation plugin it needs to be deployed to the cloud. The most common way is to deploy it to the same cluster that is running Scalable Pixel Streaming, however, you can deploy it anywhere that is publically accessible.

### Building a container 

In order to deploy your plugin to the Scalable Pixel Streaming cluster, it first needs to be put into a container. You can follow the docker guide [here](https://docs.docker.com/language/golang/build-images/) or check out the Dockerfile we have in examples [here](TODO:Link to example)

### Deploying to the cluster

Scalable Pixel Streaming is installed onto a [kubernetes](https://kubernetes.io/) cluster. Depending on what cloud platform you've installed Scalable Pixel Streaming onto will depend on how you access your cluster. Refer to the [Scalable Pixel Streaming documentation](http://docs.beta.scalablestreaming.io/) for information on how to access your cluster.

