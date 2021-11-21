package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type DecodeSaveDataCmd struct {
	File *os.File `arg:"" help:"Path to encrypted local save data file." name:"file"`
}

func (d *DecodeSaveDataCmd) Decrypt(inBytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(saveDataKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, saveDataIV)
	bytes := make([]byte, len(inBytes))
	mode.CryptBlocks(bytes, inBytes)
	len := binary.LittleEndian.Uint32(bytes)

	return bytes[:len], nil
}

func (d *DecodeSaveDataCmd) Run() error {
	inCiphertext, err := ioutil.ReadAll(d.File)
	if err != nil {
		return errors.Wrap(err, "Failed to read input file")
	}
	d.File.Close()

	saveDataBson, err := d.Decrypt(inCiphertext)
	if err != nil {
		return errors.Wrap(err, "Failed to decrypt input file")
	}
	var saveDataDoc bson.D
	err = bson.Unmarshal(saveDataBson, &saveDataDoc)
	if err != nil {
		return errors.Wrap(err, "Failed to decode BSON")
	}
	saveDataJson, err := bson.MarshalExtJSON(saveDataDoc, false, false)
	if err != nil {
		return errors.Wrap(err, "Failed to encode to JSON")
	}
	os.Stdout.Write(saveDataJson)

	return nil
}
