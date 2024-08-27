## Usage

The following example demonstrates how to implement the instance customisation API into a Go application, however it is possible to generate the necessary code in other languages.

### Generating code

This repository contains the generated code necessary to use the instance customisation plugin API in a Go application inside the `customisation` package. If you want to use the API in any other language, refer to [how to generate code from protocol buffers](https://protobuf.dev/reference/) for your language of choice.

### Adding customisation package to your Go application

You will need to add the SPS APIs package to your Go application:

```
go get github.com/scalablepixelstreaming/apis@latest
```

### Importing the package

Once the SPS APIs package has been added to your Go application, you will then be able to import it:

```
"github.com/scalablepixelstreaming/apis/pkg/customisation"
```

### Implementing the API in Go

Once imported, there are two main components to the API:

1) The server, which you can define as:

```go
type InstanceCustomisationPluginServer struct {
   customisation.UnimplementedInstanceCustomisationPluginServer
}
```

2) The pointer receiver to implement the actual gRPC call that is made from Scalable Pixel Streaming:

```go
func (server *InstanceCustomisationPluginServer) UpdateRuntimeOptions(ctx context.Context, req *customisation.UpdateRuntimeOptionsRequest) (*customisation.UpdateRuntimeOptionsResponse, error) {
   // my implementation here
}
```
