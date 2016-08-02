// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/luci/luci-go/grpc/prpc"
	"golang.org/x/net/context"

	"github.com/nodirt/prpc-example/helloworld/proto"
)

var (
	server   = flag.String("server", "", "host of the helloworld service")
	name     = flag.String("name", "", "user name")
	insecure = flag.Bool("insecure", false, "use HTTP instead of HTTPS")
)

func main() {
	flag.Parse()

	client := &prpc.Client{Host: *server, Options: prpc.DefaultOptions()}
	if *insecure {
		client.Options.Insecure = true
	}
	greeter := helloworld.NewGreeterPRPCClient(client)

	ctx := context.Background()
	res, err := greeter.SayHello(ctx, &helloworld.HelloRequest{
		Name: *name,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(res.Message)
}
