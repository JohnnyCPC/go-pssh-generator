package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	wvproto "github.com/JohnnyCPC/go-pssh-generator/proto"
	"github.com/JohnnyCPC/go-pssh-generator/util"

	"github.com/golang/protobuf/proto"
)

func main() {

	Version := 0
	WIDEVINESYSTEMID := "EDEF8BA979D64ACEA3C827DCD51D21ED"

	provider := "widevine_test"

	keyid := util.CreateRandomK()
	contentid := "Test_Content_ID"
	kid, _ := hex.DecodeString(keyid)
	cid := []byte(contentid)

	// Protect Scheme: cenc (default)
	// 0x63 65 6E 63 (cenc), 0x63 65 6E 73 (cens), 0x63 62 63 31 (cbc1), 0x63 62 63 73 (cbcs)
	wv := &wvproto.WidevinePsshData{
		Provider:  &provider,
		Algorithm: wvproto.WidevinePsshData_AESCTR.Enum(),
		KeyId:     [][]byte{kid},
		ContentId: cid,
		//ProtectionScheme: &psdint,
	}

	out, _ := proto.Marshal(wv)
	outstr := fmt.Sprintf("%x", out)
	outbase64str := base64.StdEncoding.EncodeToString(out)

	//fmt.Printf("PSSH Data: %x \n", out)
	fmt.Println("PSSH Data(base64): ", outbase64str)

	result := util.PsshBoxBinaryString(Version, WIDEVINESYSTEMID, kid, outstr)
	reshex, _ := util.DecodeHex([]byte(result))
	resbase64 := base64.StdEncoding.EncodeToString(reshex)

	fmt.Printf("PSSH BOX: %v \n", result)
	fmt.Println("PSSH BOX(base64): ", string(resbase64))
	fmt.Println("KeyID: ", keyid)

}
