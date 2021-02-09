package main

import (
	"testing"

	"google.golang.org/grpc/codes"
)

func Test_gRPCStart(t *testing.T) {
	tests := []struct {
		name string
		want string
		code codes.Code
	}{
		{
			name: "test run server",
			want: "gRPC Server start",
			code: codes.OK,
		},
		{
			name: "test run server with error",
			want: "gRPCx Server start",
			code: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gRPCStart()
			if tt.code == codes.OK {
				if got != tt.want {
					t.Errorf("1: got = %v, want %v", got, tt.want)
				}
			} else {
				if got == tt.want {
					t.Errorf("2: Sholud be error occure as got != %v, want %v ", got, tt.want)
				}
			}
		})
	}
}
