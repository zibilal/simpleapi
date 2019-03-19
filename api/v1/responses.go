package v1

import "github.com/zibilal/simpleapi/api"

type GeneralSuccessCodeResponse struct {
	Code    int         `json:"code"`
	Message interface{}  `json:"message"`
}

func(r *GeneralSuccessCodeResponse) ResponseWithMessage(msg interface{}) interface {} {
	r.Code = api.GeneralSuccessCode
	r.Message = msg
	return r
}

func ()
