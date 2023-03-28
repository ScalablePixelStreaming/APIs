package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/scalablepixelstreaming/apis/pkg/customisation"
	"google.golang.org/grpc"
)

// The default port for the custom plugin for kubernetes to listen on
// This port is important as it will be the port you refer to when adding this as a plugin
// to the Scalable Pixel Streaming Application (e.g. my-plugin.default.svc.cluster.local:55774)
const PLUGIN_PORT = 55774

// Implement the generated gRPC plugin server
type InstanceCustomisationPluginServer struct {
	customisation.UnimplementedInstanceCustomisationPluginServer
}

type MyCustomData struct {

	// The name of the map the user has selected
	Map string `json:"map"`

	// The name of the user requesting the instance
	User string `json:"user"`
}

// Implement the UpdateRuntimeOptions gRPC call
// This call is made from Scalable Pixel Streaming via gRPC (https://grpc.io/docs/what-is-grpc/introduction/)
// @req - The incoming request from Scalable Pixel Streaming containing all data that's associated with current request for an instance (see proto/customisation.proto)
func (server *InstanceCustomisationPluginServer) UpdateRuntimeOptions(ctx context.Context, req *customisation.UpdateRuntimeOptionsRequest) (*customisation.UpdateRuntimeOptionsResponse, error) {

	// One of the key features to instance customisation will be ability to send data from your frontend implementation through to action at the time of requesting an instance
	// One of the best ways to achieve this is to have your data formatted as JSON. This comes through from the frontend instance request message as `PluginOptions`

	// Lets assume the following value had already come in from the frontend (in this example the frontend data will be defined here)
	req.PluginOptions = `{"map":"Level1","user":"User1"}`

	// convert the incoming frontend data from json to our struct type
	options := &MyCustomData{}
	json.Unmarshal([]byte(req.PluginOptions), options)

	// now we can do whatever we want with the plugin options
	// lets make sure there is a user
	if options.User == "" {
		return &customisation.UpdateRuntimeOptionsResponse{RuntimeOptions: req.RuntimeOptions}, errors.New("username is required")
	}

	// Dynamically adding a volume mount to the requested instance.
	req.RuntimeOptions.VolumeMounts = append(req.RuntimeOptions.VolumeMounts, &customisation.RuntimeOptions_VolumeMounts{
		Name:      "my-custom-mount",
		MountPath: "/home/ue4/my-custom-mount",
		ReadOnly:  false,
	})

	// Here is where you might do something with loading a map that was selected by the user
	log.Println("Here is where we might load the following map: " + options.Map)

	// Appending a custom argument to the Unreal application at runtime
	req.RuntimeOptions.Args = append(req.RuntimeOptions.Args, "-MyCustomArgument")

	// Adding environment variables to the container
	if req.RuntimeOptions.EnvironmentVariables == nil {
		req.RuntimeOptions.EnvironmentVariables = map[string]string{}
	}
	req.RuntimeOptions.EnvironmentVariables["MyCustomEnvVar"] = "MyCustomValue"

	// Changing the resolution
	req.RuntimeOptions.Resolution.X = 1280
	req.RuntimeOptions.Resolution.Y = 720

	// Setting the max FPS
	req.RuntimeOptions.PixelStreaming.WebRTC.MaxFPS = 30

	// return the gRPC call with updated runtime options
	return &customisation.UpdateRuntimeOptionsResponse{RuntimeOptions: req.RuntimeOptions}, nil
}

func main() {

	// Attempt to listen on the default port number for Instance Customisation Plugins
	sock, err := net.Listen("tcp", fmt.Sprintf(":%d", PLUGIN_PORT))
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Create our gRPC server and start serving requests
	log.Println("Listening on " + fmt.Sprint(PLUGIN_PORT))
	grpcServer := grpc.NewServer()
	customisation.RegisterInstanceCustomisationPluginServer(grpcServer, &InstanceCustomisationPluginServer{})

	if err := grpcServer.Serve(sock); err != nil {
		log.Fatalln(err.Error())
	}
}
