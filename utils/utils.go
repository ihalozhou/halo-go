package utils

import (
	"HaloAdmin/halo/utils/encrypt/_md5"
	"encoding/json"
	_uuid "github.com/satori/go.uuid"
	"reflect"
	"time"
)

// TimeToFormatString 时间转格式化字符串 默认 Format "2006-01-02 15:04:05"
func TimeToFormatString(time time.Time, format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	return time.Format(format)
}

// InStringArray InArray
func InStringArray(value string, array []interface{}) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// StructToMap 结构体转map
func StructToMap[T any](obj T) (map[string]interface{}, error) {
	// 结构体转json
	j, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	// json转map
	var m map[string]interface{}
	err1 := json.Unmarshal(j, &m)
	if err1 != nil {
		return nil, err1
	}
	return m, nil
}

// StructCopy 结构体拷贝 参数均为 指针类型 没有考虑嵌套的问题
func StructCopy(from, to interface{}) {
	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to)
	// 必须是指针类型
	if fromValue.Kind() != reflect.Ptr || toValue.Kind() != reflect.Ptr {
		return
	}
	if fromValue.IsNil() || toValue.IsNil() {
		return
	}
	// 获取到来源数据
	fromElem := fromValue.Elem()
	// 需要的数据
	toElem := toValue.Elem()
	for i := 0; i < toElem.NumField(); i++ {
		toField := toElem.Type().Field(i)
		// 看看来源的结构体中是否有这个属性
		fromFieldName, ok := fromElem.Type().FieldByName(toField.Name)
		// 存在相同的属性名称并且类型一致
		// TODO 可以根据需要判断是否是空值
		if ok && fromFieldName.Type == toField.Type {
			toElem.Field(i).Set(fromElem.FieldByName(toField.Name))
		}
	}
}

// GenerateMD5 生成MD5密文
func GenerateMD5(encrypt string) (md5 string, err error) {
	encrypt, err = _md5.Encrypt(encrypt)
	if err != nil {
		return "", err
	}

	md5, err = _md5.Encrypt(encrypt)
	if err != nil {
		return "", err
	}

	return md5, nil
}

func GenerateUUID() (uuid string) {
	return _uuid.NewV4().String()
}

//
//func GeneratePass(password string, salt string) (encryptPwd string, s string) {
//	md5, err := GenerateMD5(password + salt)
//	if err != nil {
//		return "", s
//	}
//	return md5, s
//}

func TimeStrYMDHMS() string {
	return ""
}
