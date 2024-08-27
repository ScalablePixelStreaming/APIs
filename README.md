# Scalable Pixel Streaming APIs

This repository contains definitions and Go packages for consuming the public APIs exposed by Scalable Pixel Streaming that are versioned independently of SPS itself. 

The REST API is currently versioned alongside SPS and is not included in this repository. Refer to the [REST API section of our documentation](http://docs.scalablestreaming.io/tools-and-software/rest-api) for more information.

The following APIs are currently available:

- **Authentication Plugin Plugin API** for providing authentication for Pixel Streaming applications.

- **Instance Customisation Plugin API** for dynamically modifying the runtime options of individual instances of a Pixel Streaming application.

## Instance Customisation Plugin API Docs

This repository contains the following supporting documentation for the instance customisation plugin:

- [Overview of instance and application customisation](/docs/instance-customisation/overview.md)
- [Plugin option and runtime options](/docs/instance-customisation/what-can-i-customise.md)
- [Implementing the API in an application](/docs/instance-customisation/usage.md)
- [Deploying the plugin to the cloud](/docs/instance-customisation/deploy.md) 
- [Configure your application to use the plugin](/docs/instance-customisation/application.md)

## Legal

Copyright &copy; 2024, TensorWorks Pty Ltd. Licensed under the MIT License, see the [license](./LICENSE) file for details.
