package navprot

import (
	"reflect"
	"testing"
)

func TestNDTP_Parse(t *testing.T) {
	tests := []struct {
		name        string
		message     []byte
		wantRestBuf []byte
		wantNDTP    *NDTP
		wantErr     bool
	}{
		{"navigation", packetNav(), []byte(nil), ndtpNav(), false},
		{"external_device", packetExt(), []byte{1}, ndtpExt(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ndtp := new(NDTP)
			restBuf, err := ndtp.Parse(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("NDTP.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(restBuf, tt.wantRestBuf) {
				t.Errorf("NDTP.Parse() = %v, want %v", restBuf, tt.wantRestBuf)
			}
			if !reflect.DeepEqual(ndtp, tt.wantNDTP) {
				t.Error("got:      ", ndtp, "\nexpected: ", tt.wantNDTP)
			}
		})
	}
}

func packetNav() []byte {
	return []byte{0, 80, 86, 161, 44, 216, 192, 140, 96, 196, 138, 54, 8, 0, 69, 0, 0, 129, 102, 160, 64, 0, 125, 6,
		18, 51, 10, 68, 41, 150, 10, 176, 70, 26, 236, 153, 35, 56, 151, 147, 73, 96, 98, 94, 76, 40, 80,
		24, 1, 2, 190, 27, 0, 0, 126, 126, 74, 0, 2, 0, 107, 210, 2, 0, 0, 0, 0, 0, 0, 1, 0, 101, 0, 1, 0, 171,
		20, 0, 0, 0, 0, 36, 141, 198, 90, 87, 110, 119, 22, 201, 186, 64, 33, 224, 203, 0, 0, 0, 0, 83, 1, 0,
		0, 220, 0, 4, 0, 2, 0, 22, 0, 67, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 167, 97, 0, 0, 31, 6, 0, 0, 8,
		0, 2, 0, 0, 0, 0, 0}
}

func ndtpNav() *NDTP {
	data := NavData{1522961700, 37.6925783, 55.7890249, 339, 0, false, 1, 1, true}
	nph := NphData{1, 101, true, 5291, &data}
	npl := NplData{0x02, make([]byte, 4), 0x00}
	packExpected := []byte{126, 126, 74, 0, 2, 0, 107, 210, 2, 0, 0, 0, 0, 0, 0, 1, 0, 101, 0, 1, 0, 171,
		20, 0, 0, 0, 0, 36, 141, 198, 90, 87, 110, 119, 22, 201, 186, 64, 33, 224, 203, 0, 0, 0, 0, 83, 1, 0,
		0, 220, 0, 4, 0, 2, 0, 22, 0, 67, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 167, 97, 0, 0, 31, 6, 0, 0, 8,
		0, 2, 0, 0, 0, 0, 0}
	return &NDTP{&npl, &nph, packExpected}
}

func packetExt() []byte {
	return []byte{126, 126, 90, 1, 2, 0, 33, 134, 2, 0, 4, 0, 0, 144, 7, 5, 0, 100, 0, 0, 0, 1, 0, 0, 0, 18, 0, 0, 128, 0, 0, 0, 0, 1, 0, 0, 0, 60, 78, 65, 86, 83, 67, 82, 32, 118, 101, 114, 61, 49, 46, 48, 62, 60, 73, 68, 62, 49, 56, 60, 47, 73, 68, 62, 60, 70, 82, 79, 77, 62, 83, 69, 82, 86, 69, 82, 60, 47, 70, 82, 79, 77, 62, 60, 84, 79, 62, 85, 83, 69, 82, 60, 47, 84, 79, 62, 60, 84, 89, 80, 69, 62, 81, 85, 69, 82, 89, 60, 47, 84, 89, 80, 69, 62, 60, 77, 83, 71, 32, 116, 105, 109, 101, 61, 54, 48, 32, 98, 101, 101, 112, 61, 49, 32, 116, 121, 112, 101, 61, 98, 97, 99, 107, 103, 114, 111, 117, 110, 100, 62, 60, 98, 114, 47, 62, 60, 98, 114, 47, 62, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 194, 251, 32, 236, 229, 237, 255, 32, 241, 235, 251, 248, 232, 242, 229, 63, 60, 98, 114, 47, 62, 60, 98, 114, 47, 62, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 60, 98, 116, 110, 49, 62, 196, 224, 60, 47, 98, 116, 110, 49, 62, 60, 98, 114, 47, 62, 60, 98, 114, 47, 62, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 60, 98, 116, 110, 50, 62, 205, 229, 242, 60, 47, 98, 116, 110, 50, 62, 60, 98, 114, 47, 62, 60, 47, 77, 83, 71, 62, 60, 47, 78, 65, 86, 83, 67, 82, 62, 1}
}

func ndtpExt() *NDTP {
	data := ExtDevice{18, 32768, 0}
	nph := NphData{NphSrvExternalDevice, nphSedDeviceTitleData, false, 1, &data}
	npl := NplData{0x02, []byte{0, 4, 0, 0}, 1936}
	packExpected := []byte{126, 126, 90, 1, 2, 0, 33, 134, 2, 0, 4, 0, 0, 144, 7, 5, 0, 100, 0, 0, 0, 1, 0, 0, 0, 18, 0, 0, 128, 0, 0, 0, 0, 1, 0, 0, 0, 60, 78, 65, 86, 83, 67, 82, 32, 118, 101, 114, 61, 49, 46, 48, 62, 60, 73, 68, 62, 49, 56, 60, 47, 73, 68, 62, 60, 70, 82, 79, 77, 62, 83, 69, 82, 86, 69, 82, 60, 47, 70, 82, 79, 77, 62, 60, 84, 79, 62, 85, 83, 69, 82, 60, 47, 84, 79, 62, 60, 84, 89, 80, 69, 62, 81, 85, 69, 82, 89, 60, 47, 84, 89, 80, 69, 62, 60, 77, 83, 71, 32, 116, 105, 109, 101, 61, 54, 48, 32, 98, 101, 101, 112, 61, 49, 32, 116, 121, 112, 101, 61, 98, 97, 99, 107, 103, 114, 111, 117, 110, 100, 62, 60, 98, 114, 47, 62, 60, 98, 114, 47, 62, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 194, 251, 32, 236, 229, 237, 255, 32, 241, 235, 251, 248, 232, 242, 229, 63, 60, 98, 114, 47, 62, 60, 98, 114, 47, 62, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 60, 98, 116, 110, 49, 62, 196, 224, 60, 47, 98, 116, 110, 49, 62, 60, 98, 114, 47, 62, 60, 98, 114, 47, 62, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 38, 110, 98, 115, 112, 59, 60, 98, 116, 110, 50, 62, 205, 229, 242, 60, 47, 98, 116, 110, 50, 62, 60, 98, 114, 47, 62, 60, 47, 77, 83, 71, 62, 60, 47, 78, 65, 86, 83, 67, 82, 62}
	return &NDTP{&npl, &nph, packExpected}
}
