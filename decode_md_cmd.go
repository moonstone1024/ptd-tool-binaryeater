package main

import (
	"encoding/json"
	"os"

	"gitee.com/binaryeater/ptd-tool/internal"
)

type DecodeMDCmd struct {
	File   *os.File `arg:"" help:"Path to encrypted MD file." name:"file"`
	Config *os.File `help:"Path to MD loader config file." name:"config" required:""`
}

func (d *DecodeMDCmd) Run() error {
	config, err := internal.LoadMDLoaderConfigFromFile(d.Config)
	d.Config.Close()
	if err != nil {
		return err
	}
	allEntries, err := internal.LoadMDFromFile(d.File, config)
	d.File.Close()
	if err != nil {
		return err
	}
	entriesJSON, err := json.Marshal(allEntries)
	if err != nil {
		return err
	}
	os.Stdout.WriteString(string(entriesJSON))
	return nil
}
