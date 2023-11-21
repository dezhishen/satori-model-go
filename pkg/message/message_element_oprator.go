package message

type MessageElementButton struct {
	*noAliasMessageElement
	//	id	string?	发	按钮的 ID
	//
	// type	string?	发	按钮的类型
	// href	string?	发	按钮的链接
	// text	string?	发	待输入文本
	// theme	string?	发	按钮的样式
	Id    string
	Type  string
	Href  string
	Text  string
	Theme string
}

func (e *MessageElementButton) Tag() string {
	return "quote"
}

func (e *MessageElementButton) Stringify() string {
	result := "<" + e.Tag()
	if e.Id != "" {
		result += ` id="` + e.Id + `"`
	}
	if e.Type != "" {
		result += ` type="` + e.Type + `"`
	}
	if e.Href != "" {
		result += ` href="` + e.Href + `"`
	}
	if e.Text != "" {
		result += ` text="` + e.Text + `"`
	}
	if e.Theme != "" {
		result += ` theme="` + e.Theme + `"`
	}
	return result + " />"
}
