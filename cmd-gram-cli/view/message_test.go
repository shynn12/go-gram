package view

import (
	"cmd-gram-cli/models"
	"testing"
)

func TestMessages(t *testing.T) {
	type args struct {
		msg *models.MessageDTO
		u   *models.User
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "1", args: args{msg: &models.MessageDTO{UserID: 0, Body: "hello"}, u: &models.User{Email: "@@@", ID: 0}}},
		{name: "2", args: args{msg: &models.MessageDTO{UserID: 1, Body: "hello"}, u: &models.User{Email: "@@@", ID: 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Messages(tt.args.msg, tt.args.u)
		})
	}
}
