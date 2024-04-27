package utils

import (
	"crypto/md5"
	"encoding/hex"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func FmtResponse(data interface{}, err error) Response {
	if err != nil {
		return Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		}
	}
	return Response{
		Success: true,
		Message: "Success",
		Data:    data,
	}
}

func Md5String(content string) string {
	hash := md5.Sum([]byte(content))
	return hex.EncodeToString(hash[:])
}
