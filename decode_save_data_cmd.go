package main

import (
	"os"

	"gitee.com/binaryeater/ptd-tool/internal"
)

type DecodeSaveDataCmd struct {
	File *os.File `arg:"" help:"Path to encrypted local save data file." name:"file"`
}

func (d *DecodeSaveDataCmd) Run() error {
	saveDataJson, err := internal.DecodeSaveData(d.File)
	d.File.Close()
	if err != nil {
		return err
	}
	os.Stdout.WriteString(saveDataJson)

	return nil
}
