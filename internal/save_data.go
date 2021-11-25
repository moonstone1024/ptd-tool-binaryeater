package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"io/ioutil"
	"os"

	"github.com/mergermarket/go-pkcs7"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func decryptSaveData(inBytes []byte) ([]byte, error) {
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

func DecodeSaveData(file *os.File) (string, error) {
	inCiphertext, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read input file")
	}

	saveDataBson, err := decryptSaveData(inCiphertext)
	if err != nil {
		return "", errors.Wrap(err, "Failed to decrypt input file")
	}
	var saveDataDoc bson.D
	err = bson.Unmarshal(saveDataBson, &saveDataDoc)
	if err != nil {
		return "", errors.Wrap(err, "Failed to decode BSON")
	}
	saveDataJson, err := bson.MarshalExtJSON(saveDataDoc, false, false)
	if err != nil {
		return "", errors.Wrap(err, "Failed to encode to JSON")
	}

	return string(saveDataJson), nil
}

func encryptSaveData(inBytes []byte) ([]byte, error) {
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

func EncodeSaveData(file *os.File) ([]byte, error) {
	saveDataJsonBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var saveDataDoc bson.D
	err = bson.UnmarshalExtJSON(saveDataJsonBytes, false, &saveDataDoc)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decode JSON")
	}

	inBytes, err := bson.Marshal(saveDataDoc)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encode to BSON")
	}

	bytes, err := encryptSaveData(inBytes)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to encrypt BSON")
	}

	return bytes, nil
}
