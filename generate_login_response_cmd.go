package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gitee.com/binaryeater/ptd-tool/internal/generator"
	"github.com/pkg/errors"
)

type GenerateLoginResponseCmd struct {
	DecodedMDDir string   `name:"decoded-md-dir" help:"Path to Decoded MD JSON files" required:""`
	Config       *os.File `name:"config" help:"Path to config file" required:""`
}

func (g *GenerateLoginResponseCmd) Run() error {
	cfgBytes, err := ioutil.ReadAll(g.Config)
	g.Config.Close()
	if err != nil {
		return errors.Wrap(err, "Failed to read config file")
	}
	cfg := generator.LoginResponseGeneratorConfig{}
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		return errors.Wrap(err, "Failed to decode config JSON")
	}

	r := generator.GenearateLogin(g.DecodedMDDir, cfg)
	bytes, err := json.Marshal(r)
	if err != nil {
		return err
	}
	os.Stdout.Write(bytes)
	return nil
}
