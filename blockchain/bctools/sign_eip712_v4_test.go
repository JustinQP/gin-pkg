package bctools

import (
	"testing"
)

func TestEip712V4Sign(t *testing.T) {
	type args struct {
		privKey         string
		domainSeparator string
		structHash      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1) Normal signature", args{
			"dd0927dbac09fb670433cd3adc302788763166609e5807c68fd37b83542097f2",
			"0x5bd49e0db6cf393076fe141d41e36a8eaf3013b2264eff33410abb11c6b7d2a3",
			"0x9c0bdcc036090dd65b5ab4241dbf17f2001d1f6933c3ad2a7875f8eae2091d17"},
			"75c4483354cab17e3e5eb64b5fd7119d319d0d244a38afd5a1febf48be67b8015f4c169beec517d0e595bf8d40085806ce4e15ca60ce812bfd1d02a732a3398600", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eip712V4Sign(tt.args.privKey, tt.args.domainSeparator, tt.args.structHash)
			if (err != nil) != tt.wantErr {
				t.Errorf("Eip712V4Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Eip712V4Sign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEip712V4SignVerify(t *testing.T) {
	type args struct {
		address         string
		domainSeparator string
		structHash      string
		sign            string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1) Normal signature", args{
			"0x17a99B62Eb6Db79D2b791eA895Dd61A404074C39",
			"0x5bd49e0db6cf393076fe141d41e36a8eaf3013b2264eff33410abb11c6b7d2a3",
			"0x9c0bdcc036090dd65b5ab4241dbf17f2001d1f6933c3ad2a7875f8eae2091d17",
			"75c4483354cab17e3e5eb64b5fd7119d319d0d244a38afd5a1febf48be67b8015f4c169beec517d0e595bf8d40085806ce4e15ca60ce812bfd1d02a732a3398600",
		},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eip712V4SignVerify(tt.args.address, tt.args.domainSeparator, tt.args.structHash, tt.args.sign); got != tt.want {
				t.Errorf("Eip712V4SignVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}
