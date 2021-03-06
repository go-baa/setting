// Package setting inherit from safeie/goconfig, rewrite method with runtime.
package setting

import (
	"os"

	"github.com/go-baa/baa"
	"github.com/safeie/goconfig"
)

type cfg struct {
	config *goconfig.ConfigFile
}

// Set add a new option
func (c *cfg) Set(key, value string) {
	c.config.AddOption(baa.Env, key, value)
}

// Remove remove an exists option
func (c *cfg) Remove(key string) {
	c.config.RemoveOption(baa.Env, key)
}

// Handle returns origin config handler
func (c *cfg) Handle() *goconfig.ConfigFile {
	return c.config
}

// GetPrefix ...
func (c *cfg) GetPrefix(prefix string, keep bool) (map[string]string, error) {
	return c.config.GetPrefix(baa.Env, prefix, keep)
}

// GetString ...
func (c *cfg) GetString(key string) (string, error) {
	val, err := c.config.GetString(baa.Env, key)
	if err != nil {
		val, err = c.config.GetString("default", key)
	}
	return val, err
}

// GetInt ...
func (c *cfg) GetInt(key string) (int, error) {
	val, err := c.config.GetInt(baa.Env, key)
	if err != nil {
		val, err = c.config.GetInt("default", key)
	}
	return val, err
}

// GetInt64 ...
func (c *cfg) GetInt64(key string) (int64, error) {
	val, err := c.config.GetInt64(baa.Env, key)
	if err != nil {
		val, err = c.config.GetInt64("default", key)
	}
	return val, err
}

// GetFloat ...
func (c *cfg) GetFloat(key string) (float64, error) {
	val, err := c.config.GetFloat(baa.Env, key)
	if err != nil {
		val, err = c.config.GetFloat("default", key)
	}
	return val, err
}

// GetBool ...
func (c *cfg) GetBool(key string) (bool, error) {
	val, err := c.config.GetBool(baa.Env, key)
	if err != nil {
		val, err = c.config.GetBool("default", key)
	}
	return val, err
}

// MustPrefix ...
func (c *cfg) MustPrefix(prefix string, keep bool) map[string]string {
	vals, err := c.config.GetPrefix(baa.Env, prefix, keep)
	if err != nil {
		return map[string]string{}
	}
	return vals
}

// MustString ...
func (c *cfg) MustString(key string, value string) string {
	val, err := c.config.GetString(baa.Env, key)
	if err != nil || val == "" {
		val = c.config.MustString("default", key, value)
	}
	return val
}

// MustInt ...
func (c *cfg) MustInt(key string, value int) int {
	val, err := c.config.GetInt(baa.Env, key)
	if err != nil || val == 0 {
		val = c.config.MustInt("default", key, value)
	}
	return val
}

// MustInt64 ...
func (c *cfg) MustInt64(key string, value int64) int64 {
	val, err := c.config.GetInt64(baa.Env, key)
	if err != nil || val == 0 {
		val = c.config.MustInt64("default", key, value)
	}
	return val
}

// MustFloat ...
func (c *cfg) MustFloat(key string, value float64) float64 {
	val, err := c.config.GetFloat(baa.Env, key)
	if err != nil || val == 0.0 {
		val = c.config.MustFloat("default", key, value)
	}
	return val
}

// MustBool ...
func (c *cfg) MustBool(key string, value bool) bool {
	val, err := c.config.GetBool(baa.Env, key)
	if err != nil {
		val = c.config.MustBool("default", key, value)
	}
	return val
}

// Config config file handler
var Config *cfg

func init() {
	var err error
	Config = new(cfg)
	file := "conf/app.ini"
	dir := os.Getenv("BAA_ROOT")
	if IsExist(file) == false && dir != "" {
		file = dir + "/" + file
	}

	confFile := os.Getenv("BAA_CONFIG_FILE")
	if confFile != "" {
		file = confFile
	}

	Config.config, err = goconfig.ReadConfigFile(file)
	if err != nil {
		// log.Printf("cannot load config file: %s, err: %v\n", file, err)
		// init an empty config
		Config.config = goconfig.NewConfigFile()
	}
}
