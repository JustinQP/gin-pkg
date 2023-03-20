package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Text(code int) string {
	return ErrInfo[code]
}

// code, codeMsg, 自定义的数据
func RespErr(code int, data interface{}) (int, Response) {
	emsg := Text(code)

	var message string
	switch v := data.(type) {
	case string:
		message = v
	case error:
		message = v.Error()
		fmt.Printf("%+v\n", v)
		if emsg != "" {
			message = fmt.Sprintf("%s: %s", emsg, message)
		}
	default:
		message = fmt.Sprintln(v)
	}

	// 为空填充信息
	if message == "" {
		message = emsg
	}

	resp := Response{
		Status:  code,
		Message: message,
	}

	return http.StatusBadRequest, resp
}

// func RespErr(code int, message string) (int, Response) {
// 	if message == "" {
// 		message = ErrInfo[code]
// 	}

// 	resp := Response{
// 		Status:  code,
// 		Message: message,
// 	}

// 	return http.StatusBadRequest, resp
// }

func Resp400(code int) (int, Response) {
	message, isOK := ErrInfo[code]
	if !isOK {
		message = ""
	}

	resp := Response{
		Status:  code,
		Message: message,
	}

	return http.StatusBadRequest, resp
}

func Resp500(code int, message string) (int, Response) {
	if message == "" {
		message = ErrInfo[code]
	}

	resp := Response{
		Status:  code,
		Message: message,
	}

	return http.StatusInternalServerError, resp
}

func RespOK(data interface{}) (int, Response) {

	if data == nil {
		data = "operate successfully!"
	}

	resp := Response{
		Status: 200,
		Data:   data,
	}

	return http.StatusOK, resp
}

func RespPageData(cnt int64, data interface{}) (int, Response) {

	result := map[string]interface{}{
		"content": data,
		"count":   cnt,
	}

	if data == nil {
		data = "operate successfully!"
	}

	resp := Response{
		Status: 200,
		Data:   result,
	}

	return http.StatusOK, resp
}

func RespOKCode(code int, message string) (int, Response) {
	if message == "" {
		message = ErrInfo[code]
	}

	resp := Response{
		Status:  code,
		Message: message,
	}
	return http.StatusOK, resp
}

func JsonOkResp(data interface{}) string {

	if data == nil {
		data = "operate successfully!"
	}

	resp := Response{
		Status: 200,
		Data:   data,
	}

	bytesResp, err := json.Marshal(resp)
	if err != nil {
		resp = Response{Status: 500, Message: "Server Error"}
	}
	return string(bytesResp)
}

func JsonErrResp(code int, message string) string {
	if message == "" {
		message = ErrInfo[code]
	}

	if code == 0 {
		code = 400
	}
	resp := Response{
		Status:  code,
		Message: message,
	}

	bytesResp, err := json.Marshal(resp)
	if err != nil {
		resp = Response{Status: 500, Message: "Server Error"}
	}
	return string(bytesResp)
}
