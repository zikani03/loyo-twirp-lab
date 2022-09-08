//go:build tools
// +build tools

package tools

import (
	_ "github.com/twitchtv/twirp/protoc-gen-twirp"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	//	_ "github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc"
)
