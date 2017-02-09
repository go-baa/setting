// Package setting provider runtime setting env.
package setting

import "gopkg.in/baa.v1"

var (
	AppName    string
	AppVersion string
	AppURL     string
	Debug      bool
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
	AppName = Config.MustString("app.name", "appName")
	AppVersion = Config.MustString("app.version", "unknown")
	AppURL = Config.MustString("app.url", "")
}
