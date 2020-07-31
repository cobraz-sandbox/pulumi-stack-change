package main

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi/pkg/v2/backend"
	"github.com/pulumi/pulumi/pkg/v2/backend/display"
	"github.com/pulumi/pulumi/pkg/v2/backend/httpstate"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v2/go/common/workspace"
)

func main() {
	ctx := context.Background()
	// Authenticate
	var err error
	cloudURL, err := workspace.GetCurrentCloudURL()
	if err != nil {
		panic(err)
	}
	// Dump stack
	var be backend.Backend
	be, err = httpstate.Login(ctx, cmdutil.Diag(), cloudURL, display.Options{})
	if err != nil {
		panic(err)
	}
	fmt.Println(cloudURL, be)
	stckName, err := be.ParseStackReference("pulumi-test")
	if err != nil {
		panic(err)
	}
	stack, err := be.GetStack(ctx, stckName)
	if err != nil {
		panic(err)
	}
	fmt.Println(stack)
}
