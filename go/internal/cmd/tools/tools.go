// Command tools exists solely as a way to import tool dependencies so that Go
// modules won't automatically remove them from go.mod. See following
// discussion: https://github.com/golang/go/issues/25922#issuecomment-402918061
package main

import (
	_ "github.com/pseudomuto/protoc-gen-doc"
)

func main() {
	panic("Why do I exist? Run 'go doc gitlab.com/gitlab-org/gitaly-proto/go/internal/cmd/tools'")
}
