package mousika

import (
	"net/http"
	"sync"

	"github.com/valyala/fasthttp"
)

//Ctx represent a Request&Response Context
type Ctx struct {
	app      *App   //reference to app server
	route    *Route //reference to route
	index    int    //index of stack
	method   string
	path     string
	param    []string //parameter values
	Fasthttp *fasthttp.RequestCtx
	err      error
}

// http.Cookie
type Cookie struct {
	http.Cookie
}

var poolCtx = sync.Pool{
	New: func() interface{} {
		return new(Ctx)
	},
}
