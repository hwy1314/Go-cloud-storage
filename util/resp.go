package util

import (
	"encoding/json"
	"fmt"
	"log"
)

// RespMsg: http��Ӧ���ݵ�ͨ�ýṹ
type RespMsg struct {
	Code	int				`json:"code"`
	Msg		string			`json:"msg"`
	Data	interface{}	`json:"data"`
}

// NewRespMsg: ����response����
func NewRespMsg(code int, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// JSONBytes: ����תjson��ʽ�Ķ���������
func (resp *RespMsg) JSONBytes() []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return r
}

// JSONString: ����תjson��ʽ��string
func (resp *RespMsg) JSONString() string {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return string(r)
}

// GenSimpleRespStream: ֻ����code��message����Ӧ��([]byte)
func GenSimpleRespStream(code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
}

// GenSimpleRespString: ֻ����code��message����Ӧ��(string)
func GenSimpleRespString(code int, msg string) string {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg)
}