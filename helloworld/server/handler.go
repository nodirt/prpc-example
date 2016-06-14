package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luci/luci-go/server/auth"
	"github.com/luci/luci-go/server/discovery"
	"github.com/luci/luci-go/server/middleware"
	"github.com/luci/luci-go/server/prpc"
	"golang.org/x/net/context"

	"github.com/nodirt/prpc-example/helloworld/proto"
)

// init registers HTTP routes.
func init() {
	// pRPC uses httprouter that implements http.Handler.
	router := httprouter.New()

	// Configure pRPC server.
	var server prpc.Server
	server.Authenticator = auth.Authenticator{} // omit authentication.
	helloworld.RegisterGreeterServer(&server, &greeterService{})
	discovery.Enable(&server)
	server.InstallHandlers(router, base)

	router.GET("/", index)

	// Plug the router into std HTTP stack.
	http.DefaultServeMux.Handle("/", router)
}

// base is the root of the middleware chain.
// This is the place where you can add a hook for all methods
// or configure the context.
func base(h middleware.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		h(context.Background(), w, r, p)
	}
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

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, indexPage)
}
