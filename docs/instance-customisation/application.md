## Configuring your application

The **application** running on Scalable Pixel Streaming needs to be told to use the instance customisation plugin via the SPS dashboard. Refer to the [official SPS documentation](http://docs.scalablestreaming.io/) for information on accessing and using the dashboard.

### Dashboard

Any **application** can have a plugin added to its configuration during application creation or editing:

![Plugin configuration in the Scalable Pixel Streaming Dashboard](./assets/dashboard-plugins.png)

The example plugin configured in the above image is deployed on the same cluster as Scalable Pixel Streaming is installed to. This endpoint can point to any location that Scalable Pixel Streaming has access to.

The endpoint format for same cluster deployment follows kubernetes [DNS Pod Service](https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/) convention:

```
service-name.namespace.svc.cluster-domain
```

Where:

- `service-name` is **sps-instance-customisation-plugin**
- `namespace` is **default** (although, this may change depending on the cloud platform you are using)
- `svc` is **svc**
- `cluster-domain` is **cluster.local**
