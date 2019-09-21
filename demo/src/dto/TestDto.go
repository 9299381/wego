package dto

import (
	"github.com/9299381/wego/validations"
	"strings"
)

// 验证函数写在 "valid" tag 的标签里
// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 各个函数的结果的 key 值为字段名.验证函数名

type TestDto struct {
	Name string `json:"name" valid:"Required;MinSize(1);MaxSize(5)"`
	Age  int    `json:"age" valid:"Required"`
	//Name   string `json:"name" valid:"Required;Match(/^wego.*/)"` // Name 不能为空并且以 wego 开头
	////有问题
	//Age    string    `json:"age" valid:"Range(1, 140)"` // 1 <= Age <= 140，超出此范围即为不合法
	//Email  string `json:"email" valid:"Email; MaxSize(100)"` // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	//Mobile string `json:"mobile" valid:"Mobile"` // Mobile 必须为正确的手机号
	//IP     string `json:"ip" valid:"IP"` // IP 必须为一个正确的 IPv4 地址
	Desc string `json:"desc" valid:"Required;Custom(CheckDesc)"`
}

func (it *TestDto) CheckDesc(v *validations.Validation) {
	if strings.Index(it.Desc, "desc") != -1 {
		_ = v.SetError("Desc", "名称里不能含有 desc")
	}

}

//最后执行
func (it *TestDto) Finish(v *validations.Validation) {
	if strings.Index(it.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		_ = v.SetError("Name", "名称里不能含有 admin")
	}
}
