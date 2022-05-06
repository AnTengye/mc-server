package utils

import (
	"fmt"
	"mc-server/config"
	"testing"
)

func TestEncodeRsa(t *testing.T) {
	type args struct {
		data string
		pKey []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRes []byte
		wantErr bool
	}{
		{
			name: "true",
			args: args{
				data: "abvc",
				pKey: []byte(config.PublicKey),
			},
			wantRes: nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := EncryptOAEP(tt.args.data, tt.args.pKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeRsa() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(string(gotRes))
		})
	}
}

func TestDecryptOAEP(t *testing.T) {
	type args struct {
		ciphertext string
		privateKey []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "true",
			args: args{
				ciphertext: "bV0Zq/Y5R6WBo9PQN+SkFAA+zoNyNfAp0j8C9zXjEenWBxX7YzWaXU+HZiqd0TUp2Cb5AmrvENbceHqJ7xSeHUhWNC7l81y1LkPahnMuiOsiycTnM684ac6GjYhtnOweceMo+SfPWOvFi9RXn26WaQWiUEOUa8RWhA7GmmrCc/akpLMSiVAmMyxs8Z9EYU0MiBZBuumGuFNEv/fFYICqbzhJk0iSmiQfhXNWag6EWnuCDsJgYCdQ3qCt3Wf/VLDchAfmcbz1ijU3WZgYEdMxWpIpch+LOfw+WGS58cb/6B1O10k7PMtXweH/8bA196gc20+d90ftYGsdn783d7GbYxcRZTbbr2SEDaEFR2wQa7cCaHSMUK4tCrrPPhep074/tIRxglfxGBU9EfEYzceHKQf2CanEKHnClMA0g3Dx3XStntmXxj4XlxOLPo4NkpRx+WkkQNnFRiX92JpDsPDslDrr1EFYox1qBp6YOjo+UH//XVZcXYsV0lnOG5pt1v1oYKJH0RSwhZjSdioJxp1r2oALqouNyhcSGSSw0uTsAYPq/AehGlEly59CATG4Z9qQdgN6AXYYm9aSUpuUn/opc0f1nu3cPQ0PWpeuJ6jJMIDFgbdEvcFgy3n9wGvop4r5/0T4S0sY9TFIcYAPHNuaQ0XN4D+XVs8zlNKg9nmBUXQ=",
				privateKey: []byte(config.PrivateKey),
			},
			want:    "abvc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptOAEP(tt.args.ciphertext, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptOAEP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptOAEP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
