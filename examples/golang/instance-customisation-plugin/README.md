## Instance Customisation Plugin Example

From inside this directory, run the following to build the instance customisation plugin binary:

```go
go run ../../../build.go --release
```

Then to build the docker image using the supplied Dockerfile:

```
docker build -t YOUR_REGISTRY/YOUR_REPO/instance-customisation-plugin:latest .
```

To push your image to a registry of your choice:

```
docker push YOUR_REGISTRY/YOUR_REPO/instance-customisation-plugin:latest
```

You can then use the above image tag to install the example helm chart supplied at [/charts/instance-customisation-plugin](/charts/instance-customisation-plugin)
