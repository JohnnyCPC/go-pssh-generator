package wvpsshdata

import (
	"fmt"
	"reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGeneratewidevinepsshdatawithcid(t *testing.T) {

	provider := "widevine_test"
	contentid := "test_id"
	cid := []byte(contentid)

	type args struct {
		provider *string
		cid      []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Correct",
			args: args{
				provider: &provider,
				cid:      cid,
			},
			want: "08011a0d7769646576696e655f746573742207746573745f6964",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, _ := proto.Marshal(Generatewidevinepsshdatawithcid(tt.args.provider, tt.args.cid))
			got := fmt.Sprintf("%x", res)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generatewidevinepsshdatawithcid() = %v, want %v", got, tt.want)
			}
		})
	}
}
