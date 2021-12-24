package utils

import (
	"testing"
)

func TestPasswordUtil_Generate(t *testing.T) {
	type fields struct {
		Password string
		SaltStr  string
	}
	f := fields{}
	type args struct {
		password string
	}
	c1 := args{
		password: "admin",
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "c1", fields: f, args: c1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PasswordUtil{}
			p.Generate(tt.args.password)
			t.Logf("%+v", p)
		})
	}
}
