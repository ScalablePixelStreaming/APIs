//go:build never

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	archiver "github.com/mholt/archiver/v3"
	module "github.com/tensorworks/go-build-helpers/pkg/module"
	network "github.com/tensorworks/go-build-helpers/pkg/network"
	system "github.com/tensorworks/go-build-helpers/pkg/system"
	protoc "github.com/tensorworks/go-build-helpers/pkg/tools/protoc"
	validation "github.com/tensorworks/go-build-helpers/pkg/validation"
)

// Alias validation.ExitIfError() as check()
var check = validation.ExitIfError

func main() {

	// Parse our command-line flags
	doRelease := flag.Bool("release", false, "builds executables for all target platforms")
	flag.Parse()

	// Disable CGO
	os.Setenv("CGO_ENABLED", "0")

	// Create a build helper for the Go module
	mod, err := module.ModuleInCwd()
	check(err)

	// Install the Go tools that we require for code generation
	check(mod.InstallGoTools([]string{
		"google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0",
	}))

	// If we haven't already installed the protobuf compiler then download and install it
	protocZip := &network.DownloadedFile{URL: protoc.ReleaseForHost("3.15.8"), Filename: network.AutoDetectFilename, Module: mod}
	if protocZip.Exists() == false {

		// Download the protoc release archive
		log.Println("Downloading and extracting:", protocZip.URL)
		check(protocZip.Download())

		// Extract the protoc archive
		check(archiver.Unarchive(protocZip.LocalPath(), mod.DownloadsDir()))

		// Move the protoc executable to our codegen tools directory
		protocExe := fmt.Sprint("protoc", system.GOEXE)
		check(os.Rename(
			filepath.Join(mod.DownloadsDir(), "bin", protocExe),
			filepath.Join(mod.CodegenToolsDir(), protocExe),
		))
	}

	// Perform code generation
	check(mod.Generate())

	// Determine if we're building our executables for just the host platform or for the full matrix of release platforms
	if *doRelease == false {
		buildOptions := module.BuildOptions{Scheme: module.Undecorated}
		check(mod.BuildBinariesForHost(module.DefaultBinDir, buildOptions))
	} else {
		buildOptions := module.BuildOptions{Scheme: module.PrefixedDirs, AdditionalFlags: []string{}}
		check(mod.BuildBinariesForMatrix(
			module.DefaultBinDir,
			buildOptions,
			module.BuildMatrix{
				Platforms:     []string{"windows", "linux"},
				Architectures: []string{"amd64"},
				Ignore:        []string{},
			},
		))
	}
}
