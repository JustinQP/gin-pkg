package bctools

import "testing"

func TestVerifyAddress(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{"(11)", args{"0x2624f6118DbecC733277d05D34B63556A9D19Ae5"}, "0x2624f6118DbecC733277d05D34B63556A9D19Ae5", true},
		{"(12)", args{"0x2624f6118Dbecc733277d05d34B63556A9D19ae5"}, "0x2624f6118DbecC733277d05D34B63556A9D19Ae5", true},
		{"(21)", args{""}, "0x0000000000000000000000000000000000000000", false},
		{"(22)", args{"aaa"}, "0x0000000000000000000000000000000000000aaa", false},
		{"(23)", args{"0x2624f6118DbecC733277d05D34B"}, "", false},
		{"(24)", args{"0x2624f6118DbecC733277d05D34B63556A9D19Ae5aaa111CCC"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := VerifyAddress(tt.args.addr)
			if tt.want != "" && got != tt.want {
				t.Errorf("VerifyAddress() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("VerifyAddress() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
