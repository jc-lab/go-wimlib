//go:build !linux
// +build !linux

package apply

import "github.com/jc-lab/go-wimlib/model"

func checkTarget(target string) model.WimlibExtractFlag {
	return 0
}
