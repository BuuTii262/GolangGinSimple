package helper

import "MyGO.com/m/model"

type Response struct {
	ErrorCode    uint64      `json:"err_code"`
	ErrorMessage string      `json:"err_msg"`
	Data         interface{} `json:"data"`
}

type ResponseErr struct {
	ErrorCode    uint64 `json:"err_code"`
	ErrorMessage string `json:"err_msg"`
}

type ResponseListData struct {
	List  []model.User `json:"list"`
	Total int64        `json:"total"`
}

type EmptyObj struct{}

func ResponseData(err_code uint64, err string, data interface{}) Response {
	response := Response{
		ErrorCode:    err_code,
		ErrorMessage: err,
		Data:         data,
	}
	return response
}

func ResponseErrorData(err_code uint64, err string) ResponseErr {
	response := ResponseErr{
		ErrorCode:    err_code,
		ErrorMessage: err,
	}
	return response
}
