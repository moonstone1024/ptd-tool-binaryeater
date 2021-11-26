package internal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type mdLoader struct {
	byteReader *bytes.Reader
}

func (m *mdLoader) load(config *MDLoaderConfig) ([]map[string]interface{}, error) {
	entryCount, err := m.loadEntryCount()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read entry count")
	}

	allEntries := make([]map[string]interface{}, 0)
	for i := 0; i < entryCount; i++ {
		entry := map[string]interface{}{}
		for _, field := range config.Fields {
			value, err := m.loadField(field)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("Failed to load a field. entry: %d/%d, name: %s", i, entryCount, field.Name))
			}
			if field.Include {
				entry[field.Name] = value
			}
		}

		if config.AddIndex {
			entry["index"] = i + 1
		}

		allEntries = append(allEntries, entry)
	}

	return allEntries, nil
}

func (m *mdLoader) loadEntryCount() (int, error) {
	header := make([]byte, 16)
	_, err := io.ReadAtLeast(m.byteReader, header, 16)
	entryCount := binary.LittleEndian.Uint32(header[4:])
	if err != nil {
		return 0, err
	}

	return int(entryCount), nil
}

func (m *mdLoader) decodeObfuscatedString(inCharacters []rune) string {
	shift := 11
	step := 3
	decodedChars := make([]rune, len(inCharacters))
	for i, char := range inCharacters {
		decodedChar := int(char) + shift
		if decodedChar > 65535 {
			decodedChar -= 65535
		}
		decodedChars[i] = rune(decodedChar)
		shift += step
	}

	return string(decodedChars)
}

func (m *mdLoader) decodeObfuscatedInt(inValue int) int {
	return inValue ^ 0x7fef
}

func (m *mdLoader) loadField(field MDField) (interface{}, error) {
	switch field.Type {
	case "string":
		count, err := binary.ReadUvarint(m.byteReader)
		if err != nil {
			pos, _ := m.byteReader.Seek(0, io.SeekCurrent)
			return nil, errors.Wrap(err, fmt.Sprintf("Failed to read string length at %d", pos))
		}
		b := make([]byte, count)
		_, err = io.ReadAtLeast(m.byteReader, b, int(count))
		if err != nil {
			pos, _ := m.byteReader.Seek(0, io.SeekCurrent)
			return nil, errors.Wrap(err, fmt.Sprintf("Failed to read string value at %d", pos))
		}
		obfuscatedString := string(b)
		return m.decodeObfuscatedString([]rune(obfuscatedString)), nil
	case "int":
		b := make([]byte, 4)
		_, err := io.ReadAtLeast(m.byteReader, b, 4)
		if err != nil {
			pos, _ := m.byteReader.Seek(0, io.SeekCurrent)
			return nil, errors.Wrap(err, fmt.Sprintf("Failed to read int value at %d", pos))
		}
		return m.decodeObfuscatedInt(int(binary.LittleEndian.Uint32(b))), nil
	}

	return nil, errors.Errorf("Unknown field type: %s", field.Type)
}
