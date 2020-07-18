package apiserver

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name       string
		configPath string
		want       *Config
		wantErr    bool
	}{
		{
			"successfully",
			"../../configs/apiserver.yml",
			&Config{
				Server:   server{"0.0.0.0", "5000"},
				Database: database{"127.0.0.1", "5432", "", ""},
				Static:   static{"/", "./static"},
			},
			false,
		},
		{
			"fail",
			"",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewConfig(tt.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateConfigPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"no args",
			args{},
			true,
		},
		{
			"empty config path",
			args{path: ""},
			true,
		},
		{
			"wrong config path",
			args{path: "C:\\config.yml"},
			true,
		},
		{
			"successfully getting config path",
			args{path: "../../configs/apiserver.yml"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateConfigPath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ValidateConfigPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
