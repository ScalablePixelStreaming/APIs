## Deploying to the cloud

Once you have built your instance customisation plugin it needs to be deployed to the cloud. The most common way is to deploy it to the same cluster that is running Scalable Pixel Streaming, however, you can deploy it anywhere that is publicly accessible.

### Building a container 

In order to deploy your plugin to the Scalable Pixel Streaming cluster, it first needs to be containerised. You can follow the docker guide [here](https://docs.docker.com/language/golang/build-images/) or check out the Dockerfile we have in examples [here](examples/golang/instance-customisation-plugin)

### Deploying to the cluster

Scalable Pixel Streaming is installed onto a [kubernetes](https://kubernetes.io/) cluster. The cloud platform you have installed Scalable Pixel Streaming onto will inform how you access your cluster. Refer to the [Scalable Pixel Streaming documentation](http://docs.beta.scalablestreaming.io/) for information on how to access your cluster.

There is a helm chart provided in this repository available [here](charts/instance-customisation-plugin). This chart includes the necessary resources to deploy an instance customisation plugin. By default the helm chart comes with the [example plugin](examples/golang/instance-customisation-plugin) already configured.

To change the helm chart to use your own containerised instance customisation plugin:

```
helm install instance-customisation-plugin charts/instance-customisation-plugin --set image=your_registry/your_repo/image:tag
```

Where the image `your_registry/your_repo/image:tag` is a valid image tag to your containersed instance customisation plugin.


If your container is stored in a private repository, additional values are required when installing the instance customisation plugin helm chart.

**imageCredentials.registry**

The URL for your container registry, for example, if you are using Docker this would be `https://index.docker.io/v2/`

**imageCredentials.username**

The username that you use to authenticate with your container registry

**imageCredentials.password**

The password that you use to authenticate with your container registry


An example of what a `helm install` command might look like for a private repository:
```
helm install instance-customisation-plugin charts/instance-customisation-plugin --set imageCredentials.registry=MY_REGISTRY --set imageCredentials.username=MY_USERNAME --set imageCredentials.password=MY_PASSWORD  --set image=your_registry/your_repo/image:tag
```
