package internal

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type PTDResponse struct {
	PM string `json:"pm"`
}

func removePadding(inCiphertext []byte) []byte {
	for i := 0; i < len(inCiphertext); i++ {
		if inCiphertext[i] == 0 {
			return inCiphertext[:i]
		}
	}

	return inCiphertext
}

func tryDecodeResponse(inCiphertext []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", errors.Wrap(err, "aes.NewCipher failed")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	ciphertext := make([]byte, len(inCiphertext))
	mode.CryptBlocks(ciphertext, inCiphertext)
	ciphertext = removePadding(ciphertext)
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

func DecodeResponse(file *os.File, sharedSecurityKeyFile *os.File) (string, error) {
	inBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read input file")
	}

	var resp PTDResponse
	err = json.Unmarshal(inBytes, &resp)
	if err != nil {
		return "", errors.Wrap(err, "json.Unmarshal failed")
	}
	pm := strings.ReplaceAll(resp.PM, "=", "")
	inCiphertext, err := base64.RawStdEncoding.DecodeString(pm)
	if err != nil {
		return "", errors.Wrap(err, "failed to decode original pm base64")
	}
	var sharedSecurityKey []byte
	if sharedSecurityKeyFile != nil {
		sharedSecurityKey, err = ioutil.ReadAll(sharedSecurityKeyFile)
		if err != nil {
			return "", errors.Wrap(err, "failed to read sharedSecurityKey file")
		}
	}

	for i := 0; i < 4; i++ {
		var key []byte
		var pm string
		if sharedSecurityKey != nil {
			key = sharedSecurityKey
		} else {
			key = []byte(getApiKeyNext(i))
		}
		iv := []byte(getApiKey(i)[:16])
		pm, err = tryDecodeResponse(inCiphertext, key, iv)
		if err != nil {
			continue
		}
		return pm, nil
	}

	return "", err
}
