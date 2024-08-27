## Instance customisation

Scalable Pixel Streaming provides users with an option to configure and **deploy their own instance plugins**, customising certain aspects of their Scalable Pixel Streaming at **runtime**.

The instance customisation plugin API sends the **Instance ID**, **Player ID**, **Plugin Options**, and **Runtime Options** as part of the `UpdateRuntimeOptionsRequest` request and responds with a version of the `RuntimeOptions`.

### What is the difference between application customisation and instance customisation?

Applications can be configured with many of the same options that are available through an instance customisation plugin. The difference is that application options are applied to **every instance of that application**, whereas instance customisation plugins allow instance configuration on a per-user level and can even allow an end user to influence the instance configuration. 

It is entirely up to you how you choose to customise your instances, for example, you could add map selection, or attach storage volumes to the instance based on a user account, etc.
