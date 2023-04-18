## What can I customise?
The Instance Customisation Plugin API exposes a range of configurable options for your Scalable Pixel Streaming applications. Scalable Pixel Streaming will send what is currently configured as part of your Application to your instance customisation plugin and you have the option to modify these options as you see fit.

### Plugin Options
One of the key features to instance customisation is the ability to send data from your frontend implementation through to your instance customisation plugin at the time of requesting an instance.

This is a string that can be formatted to send any number of arbitrary options from your frontend to your instance customisation plugin and it might look something like the following:

```json
{
   "user": "UserA",
   "map": "Level_1",
   "region": ""
}
```

Keeping in mind that this data can represent **anything that you want**. It is up to you what you do with this in your instance customisation plugin.

### Runtime Options
Runtime Options allow your instance customisation plugin to access the configurations you have set for a given Version (e.g. resolution). These configurations will be sent back to Scalable Pixel Streaming when requesting an instance, so any modifications your instance customisation plugin makes to these options will be reflected at runtime (e.g. where instance 1 has a resolution of 1920x1080 and instance 2 has a resolution of 1280x720).

**Resolution**

The resolution of the instance

**Pixel Streaming Max FPS**

The maximum FPS that will be streamed back to the client. Beneficial on low bandwidth connections

**Arguments**

Any additional arguments you wish to add to your application. Please note that these arguments will be appended to the front of the argument list

**Volume Mounts**

Mount any additional volumes to your instance, for example, unique user data.

**Environment Variables**

Any environment variables that you want to modify

To see how you might update runtime options in your instance customisation plugin, you can refer to the [examples](examples/golang/instance-customisation-plugin) in this repository.
