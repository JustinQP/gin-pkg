package bctools

import (
	"testing"
)

// private key : 289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032
// public key  : 037db227d7094ce215c3a0f57e1bcc732551fe351f94249471934567e0f5dc1bf7
// address     : 0x970E8128AB834E8EAC17Ab8E3812F010678CF791

func TestEthSign(t *testing.T) {
	type args struct {
		pKey string
		msg  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"1) Normal signature", args{
			"289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032",
			"ethereum"},
			"48d7ef5a1b42ce54f74c34c5041d14148148add145d51067def46f21d5cb7c6f7e41ed5801832168242c7a3c144fc89ad0e994737db76770605d3487993a7c5400", false},
		{"2) Normal signature", args{
			"289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032",
			"justin"},
			"1186d7ae490473e2bc978946153cc06931cd4e4c031979fc8ea103ab8501c5e52f83169324909c88b1d4af1518bba30a02f838d7efcec41c00abba524c26b90900", false},
		{"3 )The private key is wrong", args{"289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032a", "justin"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EthSign(tt.args.pKey, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("EthSign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EthSign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthVerifySignature(t *testing.T) {
	type args struct {
		pubKey string
		msg    string
		signer string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"(1) Signature is fine", args{
			"037db227d7094ce215c3a0f57e1bcc732551fe351f94249471934567e0f5dc1bf7",
			"ethereum",
			"a58359ac5733e426f063d0cbc3123b556fe3abf29c953968cbfcbd8976a8ffff205c37b48415525227d23b94d5bbbd36805807f7f9e95edd66b5084f617fb33000",
		}, true},

		{"(2) The signature is abnormal, the signature content is incorrect", args{
			"037db227d7094ce215c3a0f57e1bcc732551fe351f94249471934567e0f5dc1bf7",
			"ethereum0001",
			"a58359ac5733e426f063d0cbc3123b556fe3abf29c953968cbfcbd8976a8ffff205c37b48415525227d23b94d5bbbd36805807f7f9e95edd66b5084f617fb33000",
		}, false},

		{"(3) The signature is abnormal, the signature result is incorrect", args{
			"037db227d7094ce215c3a0f57e1bcc732551fe351f94249471934567e0f5dc1bf7",
			"ethereum",
			"6dea785bc93454f0066ee0555b6a3b100c65a9ea9748174f1d01d214ce51b61d7f5813924d5dec65a2450bdb97bc7e27cdafb22c6f2052bda951191b11e22e4301",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EthVerifySignature(tt.args.pubKey, tt.args.msg, tt.args.signer); got != tt.want {
				t.Errorf("EthVerifySignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEthVerifySignAddress(t *testing.T) {
	type args struct {
		address   string
		msg       string
		signature string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"(11) Signature is fine, normal address format", args{
			"0x970E8128AB834E8EAC17Ab8E3812F010678CF791",
			"ethereum",
			"48d7ef5a1b42ce54f74c34c5041d14148148add145d51067def46f21d5cb7c6f7e41ed5801832168242c7a3c144fc89ad0e994737db76770605d3487993a7c5400",
		}, true},

		{"(12) Signature is fine, all lowercase mode", args{
			"0x970e8128ab834e8eac17ab8e3812f010678cf791",
			"ethereum",
			"48d7ef5a1b42ce54f74c34c5041d14148148add145d51067def46f21d5cb7c6f7e41ed5801832168242c7a3c144fc89ad0e994737db76770605d3487993a7c5400",
		}, true},

		{"(13) Signature is fine, signature with 0x", args{
			"0x970e8128ab834e8eac17ab8e3812f010678cf791",
			"ethereum",
			"0x48d7ef5a1b42ce54f74c34c5041d14148148add145d51067def46f21d5cb7c6f7e41ed5801832168242c7a3c144fc89ad0e994737db76770605d3487993a7c5400",
		}, true},

		{"(21) abnormal, the content is wrong", args{
			"0x970E8128AB834E8EAC17Ab8E3812F010678CF791",
			"ethereum0001",
			"a58359ac5733e426f063d0cbc3123b556fe3abf29c953968cbfcbd8976a8ffff205c37b48415525227d23b94d5bbbd36805807f7f9e95edd66b5084f617fb33000",
		}, false},

		{"(22) abnormal, wrong signature result", args{
			"0x970E8128AB834E8EAC17Ab8E3812F010678CF791",
			"ethereum",
			"6dea785bc93454f0066ee0555b6a3b100c65a9ea9748174f1d01d214ce51b61d7f5813924d5dec65a2450bdb97bc7e27cdafb22c6f2052bda951191b11e22e4301",
		}, false},

		{"(23) abnormal, incorrect signature content", args{
			"0x970E8128AB834E8EAC17Ab8E3812F010678CF791",
			"ethereum",
			"086ca1268fc2f05c3a08fea99db64f684bc566af03659fb85e8bf97c84cf40071d8625db555f0b8449eefafb41a38f70dbf0b7dba03100ec85cde76d0fcdf6c501",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EthVerifySignAddress(tt.args.address, tt.args.msg, tt.args.signature); got != tt.want {
				t.Errorf("EthVerifySignAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
