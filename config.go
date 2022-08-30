package inkdrop

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Config struct {
	UserName    string
	Password    string
	BindAddress string
	Port        string
}

var conf *Config

func ConfigInstance() (*Config, error) {
	if conf != nil {
		return conf, nil
	}
	conf, err := initialize()
	return conf, err
}

func initialize() (*Config, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	// if mac
	if runtime.GOOS == "darwin" {
		path = filepath.Join(path, "Library", "Application Support", "inkdrop", "config.cson")
	} else {
		return nil, fmt.Errorf("config : not supported os - %s", runtime.GOOS)
	}

	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	conf := &Config{}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		text := strings.TrimLeft(scanner.Text(), " ")
		if strings.HasPrefix(text, "username:") {
			conf.UserName = trim(text, "username:")
		} else if strings.HasPrefix(text, "password:") {
			conf.Password = trim(text, "password:")
		} else if strings.HasPrefix(text, "bindAddress:") {
			conf.BindAddress = trim(text, "bindAddress:")
		} else if strings.HasPrefix(text, "port:") {
			conf.Port = trim(text, "port:")
		}
	}

	return conf, nil
}

func trim(s string, key string) string {
	s = strings.Split(s, key)[1]
	s = strings.Trim(s, " ")
	s = strings.Trim(s, "\"")
	return s
}
