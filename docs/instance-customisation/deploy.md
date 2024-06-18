## Deploying to the cloud

Once you have built your instance customisation plugin, it needs to be deployed to the cloud. The most common way is to deploy it to the same cluster that is running Scalable Pixel Streaming, although, you can deploy it anywhere that is publicly accessible.

### Building a container 

In order to deploy your plugin to the Scalable Pixel Streaming cluster, it first needs to be containerised. You can follow the docker guide [here](https://docs.docker.com/language/golang/build-images/) or check out [our Dockerfile example](../../examples/golang/instance-customisation-plugin).

### Deploying to the cluster

Scalable Pixel Streaming is installed onto a [kubernetes](https://kubernetes.io/) cluster. The cloud platform you have installed SPS to will inform you of how to access the cluster. Refer to the [Scalable Pixel Streaming documentation](http://docs.scalablestreaming.io/) for information on how to access your cluster.

We provide [a helm chart](../../charts/instance-customisation-plugin) in this repository that includes the necessary resources to deploy an instance customisation plugin. By default, the helm chart comes with the [example plugin](../../examples/golang/instance-customisation-plugin) already configured.

Change the helm chart to use your own containerised instance customisation plugin:

```
helm install instance-customisation-plugin charts/instance-customisation-plugin --set image=<your-registry/your-repo/your-image:tag>
```

Where the image `<your-registry/your-repo/your-image:tag>` is a valid image tag to your containerised instance customisation plugin.

If your container is stored in a private repository, you will need to supply your container registry index, username, and password via the respective flags: `--set imageCredentials.registry=`, `--set imageCredentials.username=` and `--set imageCredentials.password=`.

Here is what a `helm install` command would look like for a private repository:

```
helm install instance-customisation-plugin charts/instance-customisation-plugin --set imageCredentials.registry=<your-registry-index> --set imageCredentials.username=<your-username> --set imageCredentials.password=<your-password>  --set image=<your-registry/your-repo/your-image:tag>
```
