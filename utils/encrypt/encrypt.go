package encrypt

import (
	"HaloAdmin/halo/utils/encrypt/_md5"
)

func Md5(password string) (string, error) {
	// 第一次MD5加密
	password, err := _md5.Encrypt(password)
	if err != nil {
		return "", err
	}
	// 第二次MD5加密
	password2, err := _md5.Encrypt(password)
	if err != nil {
		return "", err
	}
	return password2, nil
}
