## What is Instance Customisation
Scalable Pixel Streaming provides you with the option to configure and **deploy your own instance customisation plugins**, providing you with the ability to customise certain aspects of your Scalable Pixel Streaming application **as they are requested**. The benefit to deploying your own instance customisation plugin is that you can change your Scalable Pixel Streaming application at **runtime**.

The instance customisation plugin API sends the ****Instance ID****, ****Player ID****, ****Plugin Options**** and ****Runtime Options**** as part of the `UpdateRuntimeOptionsRequest` request and responds with a version of the `RuntimeOptions`.

### What is the difference between Application and Instance customisation?
Applications can be configured with many of the same options that are available through an instance customisation plugin. The difference is that an Application’s options are applied to **every instance of that Application**, whereas instance customisation plugin's **allow instance configuration on a per user level** and even **allow an end user to influence your instance configuration**. 

It is entirely up to you how you choose to customise your instances, these could range from things like map selection, attaching storage volumes to your instance based on a user account, etc.
