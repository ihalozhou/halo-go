package bee

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"strings"
)

type Service struct{}

func (s Service) ModelFieldUnique(params interface{}, table string, model interface{}) error {
	t := reflect.TypeOf(params)
	o := orm.NewOrm()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if strings.Contains(field.Tag.Get("haloValid"), "unique") {
			qs := o.QueryTable(table)
			qs = qs.Filter(strings.ToLower(field.Name),
				// Handler crashed with error reflect: call of reflect.Value.FieldByName on interface Value
				reflect.ValueOf(&params).Elem().FieldByName(strings.Title(field.Name)).String())
			err := qs.One(&model)
			if err == orm.ErrNoRows {
				// 没有找到记录
				fmt.Printf("Not row found")
			} else if reflect.ValueOf(&model).Elem().FieldByName(field.Name).String() ==
				reflect.ValueOf(&params).Elem().FieldByName(field.Name).String() {
				return errors.New(field.Name + "已存在")
			}
		}
	}
	return nil
}

func (s Service) Query(table interface{}) (query orm.QuerySeter) {
	query = orm.NewOrm().QueryTable(table)
	return query
}

func (s Service) Order(query orm.QuerySeter, field string, mode string) orm.QuerySeter {
	switch mode {
	case "-":
		query = query.OrderBy(mode + field)
	default:
		query = query.OrderBy(field)
	}
	return query
}
