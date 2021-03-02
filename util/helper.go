package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func PsshBoxBinaryString(ver int, sys string, kid []byte, psshdata string) string {

	psshbyte := []byte("pssh")
	psshhex := make([]byte, hex.EncodedLen(len(psshbyte)))
	hex.Encode(psshhex, psshbyte)

	psshhexstr := CreateByteString(psshhex)
	verstr := CreateIntHexByteString(ver)
	psshlen := CreateIntHexByteString(len(psshdata) / 2)

	ret := psshhexstr + verstr + sys + psshlen + psshdata
	ret = CreateIntHexByteString(len(ret)/2+4) + ret

	//fmt.Printf("%v \n", sys)
	return ret
}

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
		panic("createRandomK Fail")
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
