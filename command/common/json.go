package common

import (
	"encoding/json"
	"github.com/jc-lab/go-wimlib/model"
	"os"
)

func WriteJson(payload *model.Payload) {
	raw, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(raw)
	os.Stdout.WriteString("\n")
}

func WriteFinishSuccess(data interface{}) {
	WriteJson(&model.Payload{
		Type: model.PayloadFinish,
		Finish: &model.FinishData{
			Success: true,
			Data:    data,
		},
	})
}

func WriteFinishError(err error) {
	WriteJson(&model.Payload{
		Type: model.PayloadFinish,
		Finish: &model.FinishData{
			Success:      false,
			ErrorMessage: err.Error(),
		},
	})
}
