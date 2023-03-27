## Usage

The following usage guide is intended for users looking to implement the above API into a Go application. Although it is possible to generate the necessary code to implement the Instance Customisation Plugin API in other languages, we will not be covering that.

### Generating code

This repository contains the generated code necessary to use the Instance Customisation Plugin API in a Go application inside the `customisation` package

If you want to use the API in any other language refer to how to generate code from Protocol Buffers (https://protobuf.dev/reference/) for your language of choice

### Adding customisation package to your Go application

You will need to add the Scalable Pixel Streaming APIs package to your Go application

```
go get github.com/scalablepixelstreaming/apis@latest
```

### Importing the package

Once the Scalable Pixel Streaming APIs package has been added to your Go application, you will then be able to import it

```
"github.com/scalablepixelstreaming/apis/pkg/customisation"
```

### Implementing the API in Go

Once imported, there are two main components to the API

The server, which you can define as

```go
type InstanceCustomisationPluginServer struct {
   customisation.UnimplementedInstanceCustomisationPluginServer
}
```

And the pointer receiver to implement the actual gRPC call that’s made from Scalable Pixel Streaming

```go
func (server *InstanceCustomisationPluginServer) UpdateRuntimeOptions(ctx context.Context, req *customisation.UpdateRuntimeOptionsRequest) (*customisation.UpdateRuntimeOptionsResponse, error) {
   // my implementation here
}
```
