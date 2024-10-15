package apply

import (
	"flag"
	"github.com/jc-lab/go-wimlib/binding"
	"github.com/jc-lab/go-wimlib/command/common"
	"github.com/jc-lab/go-wimlib/model"
	"github.com/pkg/errors"
	"log"
)

type SingleFlagDefine struct {
	Name        string
	Desc        string
	OpenFlag    int
	ExtractFlag int
}
type BoolValue struct {
	Define *SingleFlagDefine
	Value  bool
}

var singleFlags = []*SingleFlagDefine{
	{
		Name:     "check",
		Desc:     "",
		OpenFlag: model.WIMLIB_OPEN_FLAG_CHECK_INTEGRITY,
	},
	{
		Name:        "recover-data",
		Desc:        "",
		ExtractFlag: model.WIMLIB_EXTRACT_FLAG_RECOVER_DATA,
	},
}

func Main(appFlag *common.AppFlag, args []string) (interface{}, error) {
	flagSet := flag.NewFlagSet(appFlag.AppName, flag.ExitOnError)
	appFlag.InitFlags(flagSet)

	singleFlagValues := make(map[string]*BoolValue)
	for _, define := range singleFlags {
		v := &BoolValue{
			Define: define,
		}
		flagSet.BoolVar(&v.Value, define.Name, false, define.Desc)
		singleFlagValues[define.Name] = v
	}

	flagSet.Parse(args)

	var openFlags int
	var extractFlags int
	for _, v := range singleFlagValues {
		if v.Value {
			openFlags |= v.Define.OpenFlag
			extractFlags |= v.Define.ExtractFlag
		}
	}

	var imageNumOrName string
	var imageNum int = 1
	var target string
	wimFile := flagSet.Arg(0)

	wim, err := binding.OpenWimWithProgress(wimFile, openFlags, func(msgType model.ProgressMsgType, msg *model.ProgressMsg, err error) model.ProgressStatus {
		info, err := msg.GetInfo()
		if appFlag.Json {
			payload := &model.Payload{
				Type:         model.PayloadProgress,
				ProgressType: &msgType,
				ProgressInfo: info,
			}
			common.WriteJson(payload)
		} else {
			log.Printf("PROGESS: %v, %+v", msgType, info)
		}
		_ = err
		return model.WIMLIB_PROGRESS_STATUS_CONTINUE
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to open wim")
	}
	defer wim.Close()

	info, err := binding.GetWimInfoEx(wim.WIMStruct)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get wim info")
	}

	switch {
	case flagSet.NArg() == 3:
		imageNumOrName = flagSet.Arg(1)
		target = flagSet.Arg(2)

		if imageNum, err = binding.ResolveImage(wim.WIMStruct, imageNumOrName); err != nil {
			return nil, errors.Wrap(err, "failed to resolve image")
		}

	case flagSet.NArg() == 2:
		target = flagSet.Arg(1)

		if info.ImageCount != 1 {
			return nil, errors.Errorf("contains %d images; Please select one (or all)", info.ImageCount)
		}
	default:
		return nil, errors.New("invalid arguments")
	}

	extractFlags |= checkTarget(target)

	err = binding.ExtractImage(wim.WIMStruct, imageNum, target, extractFlags)
	if err != nil {
		return nil, errors.Wrap(err, "failed to extract image")
	}

	return nil, nil
}
