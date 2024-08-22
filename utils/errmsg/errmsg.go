package errmsg

const (
	SUCCESS = 200

	ERROR_IMAGE_EXIST = 409
	ERROR_FILE_TYPE   = 419
)

var codeMsg = map[int]string{
	SUCCESS:           "OK",
	ERROR_IMAGE_EXIST: "图片已存在",
	ERROR_FILE_TYPE:   "图片类型错误",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
