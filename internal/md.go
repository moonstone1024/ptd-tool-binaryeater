package internal

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"os"

	"github.com/mergermarket/go-pkcs7"
	"github.com/pkg/errors"
)

func decrypt(inBytes []byte) ([]byte, error) {
	block, err := aes.NewCipher(mdKey)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, mdIV)
	bytes := make([]byte, len(inBytes))
	mode.CryptBlocks(bytes, inBytes)

	return pkcs7.Unpad(bytes, 16)
}

func LoadMDFromFile(file *os.File, config *MDLoaderConfig) ([]map[string]interface{}, error) {
	inMDBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read input file")
	}
	compressedMDBytes, err := decrypt(inMDBytes)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to decrypt input file")
	}

	buf := bytes.NewBuffer(compressedMDBytes)
	z, err := gzip.NewReader(buf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create gzip reader to decompress response data")
	}
	z.Close()
	mdBytes, err := ioutil.ReadAll(z)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decompress response data")
	}

	l := mdLoader{
		byteReader: bytes.NewReader(mdBytes),
	}

	return l.load(config)
}
