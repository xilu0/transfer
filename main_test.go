package main

import (
	"testing"
)

// func TestGetFile(t *testing.T) {
// 	file := GetFile("https://golang.google.cn/dl/go1.15.windows-amd64.msi")
// 	fmt.Println(file)
// }

func TestInspec(t *testing.T) {
	Inspect("heishui/kube-apiserver:v1.19.0")
}

func TestInstallPackage(t *testing.T) {
	InstallPackage("jq")
}

func TestCheckJq(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "first checkjq"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckJq()
		})
	}
}
