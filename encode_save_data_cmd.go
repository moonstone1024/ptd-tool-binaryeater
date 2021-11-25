package main

import (
	"os"

	"gitee.com/binaryeater/ptd-tool/internal"
)

type EncodeSaveDataCmd struct {
	File *os.File `arg:"" help:"Path to local save data file in JSON format." name:"file"`
}

func (e *EncodeSaveDataCmd) Run() error {
	bytes, err := internal.EncodeSaveData(e.File)
	if err != nil {
		return err
	}
	e.File.Close()
	os.Stdout.Write(bytes)
	return nil
}
