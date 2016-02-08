# Hello World

A pRPC HelloWorld server accepts a user name and replies with a greeting.

A pRPC HelloWorld CLI client accepts a host and a user name as arguments, 
makes an RPC and prints the greeting to stdout.

This app is deployed to https://prpc-helloworld.appspot.com.
You can use `rpc` tool to discovery and call it:

    # Install rpc
    go get -u github.com/luci/luci-go/client/cmd/rpc
    # Discover services
    rpc show prpc-helloworld.appspot.com
    # Make RPC
    rpc call prpc-helloworld.appspot.com helloworld.Greeter.SayHello -name $USER
    
Read the full article at http://nodir.io/post/138899670556/prpc.
