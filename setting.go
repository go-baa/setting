// Package setting provider runtime setting env.
package setting

import "github.com/go-baa/baa"

var (
	// AppName app name
	AppName string
	// AppVersion app version
	AppVersion string
	// AppURL app url
	AppURL string
	// Debug if open debug mode
	Debug bool
	// Maintenance if open maintenance mode
	Maintenance bool
)

func init() {
	if baa.Env != baa.PROD {
		Debug = true
	}
	v, err := Config.GetBool("debug")
	if err == nil {
		Debug = v
	}
	v, err = Config.GetBool("app.debug")
	if err == nil {
		Debug = v
	}
	Maintenance = Config.MustBool("app.maintenance", false)
	AppName = Config.MustString("app.name", "appName")
	AppVersion = Config.MustString("app.version", "unknown")
	AppURL = Config.MustString("app.url", "")
}
