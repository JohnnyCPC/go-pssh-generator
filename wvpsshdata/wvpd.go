package wvpsshdata

import (
	"encoding/hex"

	wvproto "github.com/JohnnyCPC/go-pssh-generator/proto"
)

// Generate widevine pssh data with provider, algorithm, content id
func Generatewidevinepsshdatawithcid(provider *string, cid []byte) *wvproto.WidevinePsshData {
	return &wvproto.WidevinePsshData{
		Provider:  provider,
		Algorithm: wvproto.WidevinePsshData_AESCTR.Enum(),
		ContentId: cid,
	}
}

// Generate widevine pssh data with provider, algorithm, key id
func Generatewidevinepsshdatawithkey(provider *string, kid *string) *wvproto.WidevinePsshData {

	kidhex, _ := hex.DecodeString(*kid)

	pd := &wvproto.WidevinePsshData{
		Provider:  provider,
		Algorithm: wvproto.WidevinePsshData_AESCTR.Enum(),
		KeyId:     [][]byte{kidhex},
	}

	return pd
}

// Generate widevine pssh data with provider, algorithm, content id, key id
func Generatewidevinepsshdata(provider *string, cid []byte, kid *string) *wvproto.WidevinePsshData {

	kidhex, _ := hex.DecodeString(*kid)

	// psdint := uint32(0x63656E63)
	// Protect Scheme: cenc (default)
	// 0x63 65 6E 63 (cenc), 0x63 65 6E 73 (cens), 0x63 62 63 31 (cbc1), 0x63 62 63 73 (cbcs)
	pd := &wvproto.WidevinePsshData{
		Provider:  provider,
		Algorithm: wvproto.WidevinePsshData_AESCTR.Enum(),
		ContentId: cid,
		KeyId:     [][]byte{kidhex},
		//ProtectionScheme: &psdint,
	}

	return pd
}
