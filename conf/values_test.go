package conf

import "testing"

func Test_serviceMeshMapper_GetName(t *testing.T) {
	type fields struct {
		name *string
		host *string
		Str  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "01", fields: fields{Str: ""}, want: ""},
		{name: "02", fields: fields{Str: "test"}, want: "test"},
		{name: "03", fields: fields{Str: "test=1"}, want: "test"},
		{name: "04", fields: fields{Str: "="}, want: ""},
		{name: "05", fields: fields{Str: "=="}, want: ""},
		{name: "06", fields: fields{Str: "test=test=test"}, want: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceMeshMapper{
				name: tt.fields.name,
				host: tt.fields.host,
				Str:  tt.fields.Str,
			}
			if got := s.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serviceMeshMapper_GetHost(t *testing.T) {
	type fields struct {
		name *string
		host *string
		Str  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "01", fields: fields{Str: ""}, want: ""},
		{name: "02", fields: fields{Str: "test"}, want: ""},
		{name: "03", fields: fields{Str: "test=1"}, want: "1"},
		{name: "04", fields: fields{Str: "="}, want: ""},
		{name: "05", fields: fields{Str: "=="}, want: "="},
		{name: "06", fields: fields{Str: "test=test=test"}, want: "test=test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceMeshMapper{
				name: tt.fields.name,
				host: tt.fields.host,
				Str:  tt.fields.Str,
			}
			if got := s.GetHost(); got != tt.want {
				t.Errorf("GetHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
