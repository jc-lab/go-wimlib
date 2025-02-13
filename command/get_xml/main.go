package get_xml

import (
	"flag"
	"fmt"
	"github.com/jc-lab/go-wimlib/binding"
	"github.com/jc-lab/go-wimlib/command/common"
	"github.com/jc-lab/go-wimlib/model"
	"github.com/jc-lab/go-wimlib/util"
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

	xmlRaw, err := binding.GetXMLData(wim.WIMStruct)
	if err != nil {
		return nil, errors.Wrap(err, "get xml failed")
	}

	xmlStr := util.DetectAndConvertToString(xmlRaw)

	if appFlag.Json {
		return &model.WimGetXmlData{
			Header: wimInfo,
			Xml:    xmlStr,
		}, nil
	} else {
		fmt.Println("WIM XML:")
		fmt.Println("----------------")
		fmt.Println(xmlStr)
	}

	return nil, nil
}
