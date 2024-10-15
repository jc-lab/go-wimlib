package info

import (
	"flag"
	"fmt"
	"github.com/jc-lab/go-wimlib/binding"
	"github.com/jc-lab/go-wimlib/command/common"
	"github.com/jc-lab/go-wimlib/model"
	"github.com/pkg/errors"
	"log"
)

type SingleFlagDefine struct {
	Name     string
	Desc     string
	OpenFlag int
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
	for _, v := range singleFlagValues {
		if v.Value {
			openFlags |= v.Define.OpenFlag
		}
	}

	var imageNumOrName string
	var imageNum int = 1
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

	wimInfo, err := binding.GetWimInfoEx(wim.WIMStruct)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get wim wimInfo")
	}

	switch {
	case flagSet.NArg() == 0:
		// nothing

	case flagSet.NArg() == 1:
		imageNumOrName = flagSet.Arg(1)

		if imageNum, err = binding.ResolveImage(wim.WIMStruct, imageNumOrName); err != nil {
			return nil, errors.Wrap(err, "failed to resolve image")
		}

	default:
		return nil, errors.New("invalid arguments")
	}

	if appFlag.Json {
		return &model.WimInfoData{
			Header: wimInfo,
		}, nil
	} else {
		fmt.Println("WIM Information:")
		fmt.Println("----------------")
		fmt.Printf("Path:                   %s\n", wimFile)
		fmt.Printf("GUID:                   %s\n", wimInfo.Guid)
		fmt.Printf("Version:                %d\n", wimInfo.WimVersion)
		fmt.Printf("Image Count:            %d\n", wimInfo.ImageCount)
		fmt.Printf("Compression:            %s\n", wimInfo.CompressionType.String())
		fmt.Printf("Chunk Size:             %d bytes\n", wimInfo.ChunkSize)
		fmt.Printf("Part Number:            %d/%d\n", wimInfo.PartNumber, wimInfo.TotalParts)
		fmt.Printf("Boot Index:             %d\n", wimInfo.BootIndex)
		fmt.Printf("Size:                   %d bytes\n", wimInfo.TotalBytes)
		fmt.Printf("Integrity Info:         %v\n", wimInfo.HasIntegrityTable)
		fmt.Printf("Relative path junction: %v\n", wimInfo.HasRpfix)
		fmt.Printf("Pipable:                %v\n", wimInfo.Pipable)
	}

	_ = imageNum

	return nil, nil
}
