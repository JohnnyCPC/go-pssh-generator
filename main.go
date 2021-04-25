package main

import (
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/JohnnyCPC/go-pssh-generator/psshbox"
	"github.com/JohnnyCPC/go-pssh-generator/util"
	"github.com/JohnnyCPC/go-pssh-generator/wvpsshdata"

	"google.golang.org/protobuf/proto"
)

const (
	WIDEVINESYSTEMID = "EDEF8BA979D64ACEA3C827DCD51D21ED"
)

func main() {

	Version := 0
	contentid := "Test_Content_ID"
	provider := "widevine_test"
	//keyid := "12341234123412341234123412341234"

	cidFlag := flag.String("contentid", contentid, "Content ID")
	providerFlag := flag.String("provider", provider, "Provider name")
	//kidFlag := flag.String("keyid", keyid, "Key ID")
	//hexBoolFlag := flag.Bool("hex", false, "The output can be in hex")
	//wvBoolFlag := flag.Bool("widevine", false, "widevine system id")

	flag.Parse()

	cid := []byte(*cidFlag)

	//wv := wvpsshdata.Generatewidevinepsshdata(providerFlag, cid, kidFlag)
	wv := wvpsshdata.Generatewidevinepsshdatawithcid(providerFlag, cid)
	//wv := wvpsshdata.Generatewidevinepsshdatawithkey(providerFlag, kidFlag)

	out, _ := proto.Marshal(wv)
	outstr := fmt.Sprintf("%x", out)
	outbase64str := base64.StdEncoding.EncodeToString(out)

	fmt.Println("PSSH Data: ", outstr)
	fmt.Println("PSSH Data(base64): ", outbase64str)

	result := psshbox.PsshBoxBinaryString(Version, WIDEVINESYSTEMID, outstr)
	reshex, _ := util.DecodeHex([]byte(result))
	resbase64 := base64.StdEncoding.EncodeToString(reshex)

	fmt.Printf("PSSH BOX: %v \n", result)
	fmt.Println("PSSH BOX(base64): ", string(resbase64))
	//fmt.Println("KeyID: ", *kidFlag)

}
