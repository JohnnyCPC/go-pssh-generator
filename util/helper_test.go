package util

import "testing"

func TestPsshBoxBinaryString(t *testing.T) {
	type args struct {
		ver      int
		sys      string
		kid      []byte
		psshdata string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Correct",
			args: args{
				ver:      0,
				sys:      "EDEF8BA979D64ACEA3C827DCD51D21ED",
				kid:      []byte{0xb5, 0x09, 0x16, 0x47, 0x7f, 0xb8, 0x74, 0xfc, 0xb8, 0x12, 0x20, 0xf7, 0x9b, 0xa6, 0xf9, 0x1b},
				psshdata: "08011210b50916477fb874fcb81220f79ba6f91b1a0d7769646576696e655f74657374220f546573745f436f6e74656e745f4944",
			},
			want: "000000547073736800000000EDEF8BA979D64ACEA3C827DCD51D21ED0000003408011210b50916477fb874fcb81220f79ba6f91b1a0d7769646576696e655f74657374220f546573745f436f6e74656e745f4944",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PsshBoxBinaryString(tt.args.ver, tt.args.sys, tt.args.kid, tt.args.psshdata); got != tt.want {
				t.Errorf("PsshBoxBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
