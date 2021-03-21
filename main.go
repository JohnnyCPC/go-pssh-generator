package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"

	wvproto "github.com/JohnnyCPC/go-pssh-generator/proto"
	"github.com/JohnnyCPC/go-pssh-generator/util"

	"github.com/golang/protobuf/proto"
)

const (
	WIDEVINESYSTEMID = "EDEF8BA979D64ACEA3C827DCD51D21ED"
)

func main() {

	Version := 0
	contentid := "Test_Content_ID"
	provider := "widevine_test"

	cidFlag := flag.String("contentid", contentid, "Content ID")
	providerFlag := flag.String("provider", provider, "Provider name")
	kidFlag := flag.String("keyid", "", "Key ID")
	//hexBoolFlag := flag.Bool("hex", false, "The output can be in hex")
	//wvBoolFlag := flag.Bool("widevine", false, "widevine system id")

	flag.Parse()

	cid := []byte(*cidFlag)

	// Protect Scheme: cenc (default)
	// 0x63 65 6E 63 (cenc), 0x63 65 6E 73 (cens), 0x63 62 63 31 (cbc1), 0x63 62 63 73 (cbcs)
	wv := &wvproto.WidevinePsshData{
		Provider:  providerFlag,
		Algorithm: wvproto.WidevinePsshData_AESCTR.Enum(),
		//KeyId:     [][]byte{kid},
		ContentId: cid,
		//ProtectionScheme: &psdint,
	}
	if *kidFlag != "" {
		kid, _ := hex.DecodeString(*kidFlag)
		wv.KeyId = [][]byte{kid}
	}

	out, _ := proto.Marshal(wv)
	outstr := fmt.Sprintf("%x", out)
	outbase64str := base64.StdEncoding.EncodeToString(out)

	fmt.Println("PSSH Data: ", outstr)
	fmt.Println("PSSH Data(base64): ", outbase64str)

	result := util.PsshBoxBinaryString(Version, WIDEVINESYSTEMID, outstr)
	reshex, _ := util.DecodeHex([]byte(result))
	resbase64 := base64.StdEncoding.EncodeToString(reshex)

	fmt.Printf("PSSH BOX: %v \n", result)
	fmt.Println("PSSH BOX(base64): ", string(resbase64))
	//fmt.Println("KeyID: ", *kidFlag)

}
