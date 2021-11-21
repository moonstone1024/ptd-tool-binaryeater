package main

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"os"

	"github.com/mergermarket/go-pkcs7"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type EncodeSaveDataCmd struct {
	File *os.File `arg:"" help:"Path to local save data file in JSON format." name:"file"`
}

func (e *EncodeSaveDataCmd) Encrypt(inBytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(saveDataKey)
	if err != nil {
		return nil, err
	}

	bytes, err := pkcs7.Pad(inBytes, 16)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, saveDataIV)
	mode.CryptBlocks(bytes, bytes)

	return bytes, nil
}

func (e *EncodeSaveDataCmd) Run() error {
	saveDataJsonBytes, err := ioutil.ReadAll(e.File)
	if err != nil {
		return err
	}
	e.File.Close()

	var saveDataDoc bson.D
	err = bson.UnmarshalExtJSON(saveDataJsonBytes, false, &saveDataDoc)
	if err != nil {
		return errors.Wrap(err, "Failed to decode JSON")
	}

	inBytes, err := bson.Marshal(saveDataDoc)
	if err != nil {
		return errors.Wrap(err, "Failed to encode to BSON")
	}

	bytes, err := e.Encrypt(inBytes)
	if err != nil {
		return errors.Wrap(err, "Failed to encrypt BSON")
	}
	os.Stdout.Write(bytes)

	return nil
}
