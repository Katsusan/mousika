package mousika

import (
	"time"

	"github.com/gorilla/schema"
	"github.com/valyala/fasthttp"
)

const (
	DEFAULT_MAX_BODY_SIZE   = 4 * 1024 * 1024
	DEFAULT_MAX_CONNECTIONS = 1 << 16
)

type App struct {
	srv      *fasthttp.Server
	routes   *[]Route  //router list
	settings *Settings //global server settings
}

// Settings hold server settings
type Settings struct {
	//Determine whether be strict when routing. When enabled, "/foo" and "/foo/" are treated as different routes.
	StrictRouting bool //default: false
	//Determine whether url is case-sensitive when parsed. When enbaled, "/Foo" and "/foo" are treated as different routes.
	CaseSensitive bool //default: false
	//Enables header of "Server: xx"
	ServerHeader string //default: ""
	//Enables etag header
	ETag bool //default: false
	//accpted http body size(byte)
	BodyLimitSize int //default: DEFAULT_MAX_BODY_SIZE
	//Maximum connections at the same time
	ConnectionsLimit int //default: DEFAULT_MAX_CONNECTIONS

	//Determine whether server will close connection after first response is sent.
	DisableKeepAlive bool //default: false
	//Determine whether server will sent "Date" Header unber the default behavior.
	DisableDefaultDate bool //default: false
	//Determine whether server will sent "Content-Type" Header under the default behavior.
	DisableDefaultContentType bool //default: false

	//where template files exist
	TemplateFolder string //default: "./template"
	//one of the template engines: html, amber, handlebars , mustache or pug
	TemplateEngine func(raw string, bind interface{}) (string, error) //default: nil
	//Extension for template files
	TemplateExtension string //default: ""

	//maximum time duration spent on reading full request including body
	ReadTimeout time.Duration //default: 0(unlimited)
	//maximum time duration spent on writing response
	WriteTimeout time.Duration //default: 0(unlimited)
	//maximum time duration waiting for next request coming when keepalive is enabled.
	IdleTimeout time.Duration //default: 0(unlimited)
}

// Group will serve with different url prefixes
type Group struct {
	prefix string
	app    *App
}

var formDecoder = schema.NewDecoder()
var queryDecoder = schema.NewDecoder()

//New creates mousika instance with your own settings
//only the first setting will be taken
func New(customSettings ...*Settings) *App {
	formDecoder.SetAliasTag("form")
	formDecoder.IgnoreUnknownKeys(true)
	queryDecoder.SetAliasTag("query")
	formDecoder.IgnoreUnknownKeys(true)

	app := new(App)

	app.settings = new(Settings)
	app.settings.BodyLimitSize = DEFAULT_MAX_BODY_SIZE

	if len(settings) > 0 {
		app.settings = customSettings[0]

		if app.settings.BodyLimitSize <= 0 {
			app.settings.BodyLimitSize = DEFAULT_MAX_BODY_SIZE
		}

		if app.settings.ConnectionsLimit <= 0 {
			app.settings.ConnectionsLimit = DEFAULT_MAX_CONNECTIONS
		}
	}

	return app
}
