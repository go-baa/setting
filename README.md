# setting
setting and config file manage for baa

> Note: config file must be `./conf/app.ini`

## Example

```
// conf/app.ini
[default]
# app
app.name = baaBlog
app.version = 0.1
app.url = ""
debug = false

# http
http.address = 0.0.0.0
http.port = 80
http.access_open = off

# output log to os.Stderr
log.file = os.Stderr
# 0 off, 1 fatal, 2 panic, 5 error, 6 warn, 10 info, 11 debug
log.level = 11

# development mode overwrite default config
[development]
debug = true

# production mode overwrite default config
[production]
debug = false
```

## Usage

```
package main

import "github.com/go-baa/setting"

func main() {
    appName := setting.AppName
    httpAddress := setting.Config.MustString("http.address", "127.0.0.1")
    httpPort := setting.Config.MustInt("http.port", 1323)
    httpAccessOpen := setting.Config.MustBool("http.access_open", false)
}
``

