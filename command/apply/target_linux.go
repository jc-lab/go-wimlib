//go:build linux
// +build linux

package apply

import (
	"github.com/jc-lab/go-wimlib/model"
	"gopkg.in/ro-ag/posix.v1"
	"log"
	"os"
	"syscall"
)

func checkTarget(target string) model.WimlibExtractFlag {
	stat, err := os.Stat(target)
	if err != nil {
		log.Fatalf("could not stat target %s: %s", target, err)
	}
	detail := stat.Sys().(*syscall.Stat_t)
	mode := posix.ModeT(detail.Mode)
	if mode.S_ISBLK() || mode.S_ISREG() {
		return model.WIMLIB_EXTRACT_FLAG_NTFS
	}
	return 0
}
