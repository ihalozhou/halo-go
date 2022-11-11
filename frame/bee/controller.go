package bee

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/gookit/validate"
	"halo-admin-go/constant"
	"halo-admin-go/types"
)

// Controller 基准控制器
type Controller struct {
	web.Controller
}

// #### Response 响应处理

// DATA SUCCESS 操作成功
func (c *Controller) DATA(data interface{}, code int, message string) {
	c.Data["json"] = &types.Data{
		Data:    data,
		Code:    code,
		Message: message,
	}
	e := c.ServeJSON()
	if e != nil {
		println(e)
	}
	c.StopRun()
}

// INFO ERROR 操作失败
func (c *Controller) INFO(code int, message string) {
	c.Data["json"] = &types.Info{
		Code:    code,
		Message: message,
	}
	e := c.ServeJSON()
	if e != nil {
		println(e)
	}
	c.StopRun()
}

// PAGE 分页处理
func (c *Controller) PAGE(data interface{}, page int64, count int64, limit int64, code int, message string) {
	c.Data["json"] = &types.Page{
		Page:      page,
		Limit:     limit,
		Count:     count,
		PageCount: (count + limit - 1) / limit,
		Data:      data,
		Code:      code,
		Message:   message,
	}
	e := c.ServeJSON()
	if e != nil {
		println(e)
	}
	c.StopRun()
}

// ReqFormBind 请求参数绑定 form json params为指针结构体
func (c *Controller) ReqFormBind(params interface{}) {
	if err := c.BindForm(params); err != nil {
		c.INFO(constant.RequestBindFail, err.Error())
	}
}

// ReqJsonBind 请求参数绑定 form json params为指针结构体
func (c *Controller) ReqJsonBind(params interface{}) {
	if err := c.BindJSON(params); err != nil {
		c.INFO(constant.ParamsValidateFail, err.Error())
	}
}

// ParamsValid 请求参数验证
func (c *Controller) ParamsValid(params interface{}) {
	v := validate.Struct(params)
	if !v.Validate() {
		c.INFO(constant.ParamsValidateFail, v.Errors.One())
	}
}
