package generator

import (
	"github.com/speps/go-hashids/v2"
)

func CreateNewHashId(salt string) (*hashids.HashID, error) {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 10
	hd.Alphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	return h, nil
}
