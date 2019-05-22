package main

import (
	"fmt"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Foo string
	Bar string
}

func main() {
	config, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", config)

}

func run() (*Config, error) {
	l := lua.NewState()
	defer l.Close()
	name := os.Args[1]
	if err := l.DoFile(fmt.Sprintf("%s.lua", name)); err != nil {
		return nil, err
	}
	var config Config
	if err := gluamapper.Map(l.GetGlobal(strings.ToLower(reflect.TypeOf(config).Name())).(*lua.LTable), &config); err != nil {
		return nil, err
	}
	return &config, nil
}
