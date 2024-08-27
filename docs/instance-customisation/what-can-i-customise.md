## What can be customised?

The instance customisation plugin API exposes a range of configurable options for your Scalable Pixel Streaming applications. SPS will send the current configuration of your application to your instance customisation plugin, where you will have the option to modify these options as you see fit.

### Plugin options

One of the key features of instance customisation is the ability to send data from your frontend implementation to your instance customisation plugin at the time of requesting an instance.

Here's an example of a string that can be formatted to send any number of options from your frontend to your instance customisation plugin:

```json
{
   "user": "UserA",
   "map": "Level_1",
   "region": ""
}
```

This data can represent any parameter, so it is entirely up to you how to configure your instance customisation plugin.

### Runtime options

Runtime options allow your instance customisation plugin to access the configuration you have set for a given version (e.g. stream resolution). This configuration will be sent back to Scalable Pixel Streaming when requesting an instance, so any modifications that your instance customisation plugin makes will be reflected at runtime; for example, two instances could be running with different stream resolutions.

**Resolution**

The stream resolution of the instance.

**Pixel Streaming Max FPS**

The maximum FPS that will be streamed back to the client. This limiter is beneficial on low-bandwidth connections.

**Arguments**

Any additional arguments you wish to add to your application. Note that these arguments will be appended to the front of the argument list.

**Volume Mounts**

Mount any additional volumes to your instance, for example, unique user data. Volume mounts require a [persistent volume claim](https://kubernetes.io/docs/concepts/storage/persistent-volumes/) to exist prior to attempting to add volume mounts to your instances. If volume mounts are incorrectly configured, the application instances will fail to start.

**Environment Variables**

Any environment variables that you want to modify.

**Examples**

To see how you might update runtime options in your instance customisation plugin, you can refer to the [examples](/examples/golang/instance-customisation-plugin) in this repository.
