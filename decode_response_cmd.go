package main

import (
	"os"

	"gitee.com/binaryeater/ptd-tool/internal"
)

type DecodeResponseCmd struct {
	File              *os.File `arg:"" help:"Path to encrypted json response file." name:"file"`
	SharedSecurityKey *os.File `help:"Path to shared security key file for decoding responses after login." name:"shared-security-key"`
}

func (d *DecodeResponseCmd) Run() error {
	r, err := internal.DecodeResponse(d.File, d.SharedSecurityKey)
	d.File.Close()
	if d.SharedSecurityKey != nil {
		d.SharedSecurityKey.Close()
	}
	if err != nil {
		return err
	}
	os.Stdout.WriteString(r)
	return nil
}
