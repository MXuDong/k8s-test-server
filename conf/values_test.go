package conf

import (
	"reflect"
	"testing"
)

func TestInitMeshMapper(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *serviceMeshMapper
		wantErr bool
	}{
		{name: "1", args: args{str: "GET,POST|directly|test1|http://www.baidu.com"}, want: &serviceMeshMapper{Str: "GET,POST|directly|test1|http://www.baidu.com", methodList: []string{"GET", "POST"}, host: "http://www.baidu.com", name: "test1", mode: "directly"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitMeshMapper(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitMeshMapper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitMeshMapper() got = %v, want %v", got, tt.want)
			}
		})
	}
}
