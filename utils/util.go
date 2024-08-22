package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)

func MathMD5(src io.Reader) string {
	md5Hash := md5.New()
	io.Copy(md5Hash, src)
	return fmt.Sprintf("%x", md5Hash.Sum(nil))
}
func GetHost(request *http.Request) string {
	scheme := "http"
	if request.TLS != nil {
		scheme = "https"
	}
	host := request.Host
	return scheme + "://" + host
}
