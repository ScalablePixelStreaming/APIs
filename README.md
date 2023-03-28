# Scalable Pixel Streaming APIs

This repository contains definitions and Go packages for consuming the public APIs exposed by Scalable Pixel Streaming that are versioned independently of SPS itself. (Note that the REST API is currently versioned alongside SPS and is not included in this repository.)

The following APIs are currently available:

- **Authentication Plugin Plugin API:** used to provide authentication for Pixel Streaming applications.

- **Instance Customisation Plugin API:** used to dynamically modify the runtime options of individual instances of a Pixel Streaming application.

## Instance Customisation Plugin API Docs
- [Overview](/docs/instance-customisation/overview.md)
- [What can I customise](/docs/instance-customisation/what-can-i-customise.md)
  - [Plugin Options](/docs/instance-customisation/what-can-i-customise.md#plugin-options)
  - [Runtime Options](/docs/instance-customisation/what-can-i-customise.md#runtime-options)
- [Using the API](/docs/instance-customisation/usage.md)
  - [Generating code](/docs/instance-customisation/usage.md#generating-code)
  - [Adding customisation package](/docs/instance-customisation/usage.md#adding-customisation-package-to-your-go-application)
  - [Importing the package](/docs/instance-customisation/usage.md#importing-the-package)
  - [Implementing the API](/docs/instance-customisation/usage.md#implementing-the-api-in-go)
- [Deploying the plugin](/docs/instance-customisation/deploy.md)
  - [Building a container](/docs/instance-customisation/deploy.md#building-a-container)
  - [Deploying to the cluster](/docs/instance-customisation/deploy.md#deploying-to-the-cluster)
- [Configure Application](/docs/instance-customisation/application.md)
  - [Dashboard](/docs/instance-customisation/application.md#dashboard)


## Legal

Copyright &copy; 2022, TensorWorks Pty Ltd. Licensed under the MIT License, see the file [LICENSE](./LICENSE) for details.
