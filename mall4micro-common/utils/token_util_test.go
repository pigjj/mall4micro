package utils

import (
	"testing"
)

func TestTokenUtil_Generate(t *testing.T) {
	type fields struct {
		Username string
		Email    string
		Mobile   string
		Status   int
	}
	f := fields{
		Username: "admin",
		Email:    "amazing2j@qq.com",
		Mobile:   "12312341234",
		Status:   1,
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "c1", fields: f, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			t := &TokenUtil{
				Username: tt.fields.Username,
				Email:    tt.fields.Email,
				Mobile:   tt.fields.Mobile,
				Status:   tt.fields.Status,
			}
			got, err := t.Generate()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Logf("%+v", got)
		})
	}
}

func TestTokenUtil_Parse(t1 *testing.T) {
	type fields struct {
		Username string
		Email    string
		Mobile   string
		Status   int
	}
	f := fields{}
	type args struct {
		tokenStr string
	}
	c1 := args{
		tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFtYXppbmcyakBxcS5jb20iLCJleHAiOjE2NDA5NzgwMjEsImlhdCI6MTY0MDM3MzIyMSwiaXNzIjoibWFsbDRtaWNybyIsIm1vYmlsZSI6IjEyMzEyMzQxMjM0IiwibmJmIjoxNDQ0NDc4NDAwLCJzdGF0dXMiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.x6pLT_IUXC3J8zc-6yAAPvYAGPranL3g5koaftnEUhk",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "c1", fields: f, args: c1},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &TokenUtil{}
			t.Parse(tt.args.tokenStr)
			t1.Logf("%+v", t)
		})
	}
}
