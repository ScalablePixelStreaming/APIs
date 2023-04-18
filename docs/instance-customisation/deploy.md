## Deploying to the cloud

Once you have built your instance customisation plugin it needs to be deployed to the cloud. The most common way is to deploy it to the same cluster that is running Scalable Pixel Streaming, however, you can deploy it anywhere that is publicly accessible.

### Building a container 

In order to deploy your plugin to the Scalable Pixel Streaming cluster, it first needs to be containerised. You can follow the docker guide [here](https://docs.docker.com/language/golang/build-images/) or check out the Dockerfile we have in examples [here](examples/golang/instance-customisation-plugin)

### Deploying to the cluster

Scalable Pixel Streaming is installed onto a [kubernetes](https://kubernetes.io/) cluster. The cloud platform you have installed Scalable Pixel Streaming onto will inform how you access your cluster. Refer to the [Scalable Pixel Streaming documentation](http://docs.beta.scalablestreaming.io/) for information on how to access your cluster.

There is a helm chart provided in this repository available [here](charts/instance-customisation-plugin). This chart includes the necessary resources to deploy an instance customisation plugin. By default the helm chart comes with the [example plugin](examples/golang/instance-customisation-plugin) already configured.

To change the helm chart to use your own container image:

```
helm install instance-customisation-plugin charts/instance-customisation-plugin --set image=your_registry/your_repo/image:tag
```
