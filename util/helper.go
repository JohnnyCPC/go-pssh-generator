package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func CreateIntHexByteString(tag int) string {
	res := fmt.Sprintf("%08X", []int{tag})
	return res[1 : len(res)-1]

}
func CreateByteString(tag []byte) string {
	return fmt.Sprintf("%s", tag)
}

func CreateRandomK() string {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		panic("CreateRandomK Fail")
	}
	return fmt.Sprintf("%x", key)
}

func DecodeHex(input []byte) ([]byte, error) {
	db := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(db, input)
	if err != nil {
		return nil, err
	}
	return db, nil
}
