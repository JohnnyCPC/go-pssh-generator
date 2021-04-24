package psshbox

import (
	"encoding/hex"

	"github.com/JohnnyCPC/go-pssh-generator/util"
)

//Generate Pssh Box with binary string format
func PsshBoxBinaryString(ver int, sys string, psshdata string) string {

	psshbyte := []byte("pssh")
	psshhex := make([]byte, hex.EncodedLen(len(psshbyte)))
	hex.Encode(psshhex, psshbyte)

	psshhexstr := util.CreateByteString(psshhex)
	verstr := util.CreateIntHexByteString(ver)
	psshlen := util.CreateIntHexByteString(len(psshdata) / 2)

	ret := psshhexstr + verstr + sys + psshlen + psshdata
	ret = util.CreateIntHexByteString(len(ret)/2+4) + ret

	return ret
}
