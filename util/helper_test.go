package util

import (
	"reflect"
	"testing"
)

func TestDecodeHex(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeHex(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
