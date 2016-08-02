package server

import (
	"fmt"
	"net/http"

	"github.com/luci/luci-go/grpc/discovery"
	"github.com/luci/luci-go/grpc/prpc"
	"github.com/luci/luci-go/server/router"
	"github.com/nodirt/prpc-example/helloworld/proto"
)

// init registers HTTP routes.
func init() {
	// pRPC uses httprouter that implements http.Handler.
	r := router.New()

	middleware := router.MiddlewareChain{}

	// Configure pRPC server.
	var server prpc.Server
	server.Authenticator = prpc.NoAuthenticator
	helloworld.RegisterGreeterServer(&server, &greeterService{})
	discovery.Enable(&server)
	server.InstallHandlers(r, middleware)

	r.GET("/", middleware, index)

	// Plug the router into std HTTP stack.
	http.DefaultServeMux.Handle("/", r)
}


var indexPage = `<html>
<head><title>Helloworld</title></head>
<body>

Use <code>rpc</code> tool to talk to this server,
as described in
<a href="http://nodir.io/post/138899670556/prpc">pPRC blog post</a>.

</body>
</html>
`

func index(c *router.Context) {
	c.Writer.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(c.Writer, indexPage)
}
