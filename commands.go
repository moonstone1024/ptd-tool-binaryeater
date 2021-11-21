package main

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

var key0 = [...]string{
	"mfQMG92mJ6R",
	"D2L1L1aLJSBjLlLQ3fCdS8at",
	"52hHwSInpbOVRq",
	"M88yny3dvB5ZTVHpQD",
}
var key1 = [...]string{
	"MW1LkPfxIrDB7PNTdupvw",
	"jAWLuydO",
	"sIQPwGnXVkOhMqks0p",
	"W7rkDhakfa1PI0",
}

type DecodeCmd struct {
	File              *os.File `arg:"" help:"Path to encrypted json response file." name:"file"`
	SharedSecurityKey *os.File `help:"Path to shared security key file for decoding responses after login." name:"shared-security-key"`
}

type PTDResponse struct {
	PM string `json:"pm"`
}

func GetApiKey(idx int) string {
	return key0[idx] + key1[idx]
}

func GetApiKeyNext(idx int) string {
	idx = (idx + 1) % 4
	return key0[idx] + key1[idx]
}

func RemovePadding(inCiphertext []byte) []byte {
	for i := 0; i < len(inCiphertext); i++ {
		if inCiphertext[i] == 0 {
			return inCiphertext[:i]
		}
	}

	return inCiphertext
}

func TryDecode(inCiphertext []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.Wrap(err, "aes.NewCipher failed")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	ciphertext := make([]byte, len(inCiphertext))
	mode.CryptBlocks(ciphertext, inCiphertext)
	ciphertext = RemovePadding(ciphertext)
	compressedBytes := make([]byte, base64.StdEncoding.DecodedLen(len(ciphertext)))
	n, err := base64.StdEncoding.Decode(compressedBytes, ciphertext)
	if err != nil {
		return "", errors.Wrap(err, "failed to base64 decode compressed response data")
	}

	bytesReader := bytes.NewBuffer(compressedBytes[:n])
	z, err := gzip.NewReader(bytesReader)
	if err != nil {
		return "", errors.Wrap(err, "failed to create gzip reader to decompress response data")
	}
	pmBytes, err := ioutil.ReadAll(z)
	if err != nil {
		return "", errors.Wrap(err, "failed to decompress response data")
	}

	return string(pmBytes), nil
}

func (d *DecodeCmd) Run() error {
	inBytes, err := ioutil.ReadAll(d.File)
	if err != nil {
		return errors.Wrap(err, "Failed to read input file")
	}
	d.File.Close()

	var resp PTDResponse
	err = json.Unmarshal(inBytes, &resp)
	if err != nil {
		return errors.Wrap(err, "json.Unmarshal failed")
	}
	inCiphertext, err := base64.StdEncoding.DecodeString(resp.PM)
	if err != nil {
		return errors.Wrap(err, "failed to decode original pm base64")
	}
	var sharedSecurityKey []byte
	if d.SharedSecurityKey != nil {
		sharedSecurityKey, err = ioutil.ReadAll(d.SharedSecurityKey)
		if err != nil {
			return errors.Wrap(err, "failed to read sharedSecurityKey file")
		}
		d.SharedSecurityKey.Close()
	}

	for i := 0; i < 4; i++ {
		var key []byte
		var pm string
		if sharedSecurityKey != nil {
			key = sharedSecurityKey
		} else {
			key = []byte(GetApiKeyNext(i))
		}
		iv := []byte(GetApiKey(i)[:16])
		pm, err = TryDecode(inCiphertext, key, iv)
		if err != nil {
			continue
		}
		os.Stdout.WriteString(pm)
		return nil
	}

	return err
}
