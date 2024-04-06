package message

import (
	"golang.org/x/net/html"
)

type MessageElmentBr struct {
	*noAliasMessageElement
}

func (e *MessageElmentBr) Tag() string {
	return "br"
}

func (e *MessageElmentBr) Stringify() string {
	return "<br/>"
}

func (e *MessageElmentBr) Parse(n *html.Node) (MessageElement, error) {
	return &MessageElmentBr{}, nil
}

type MessageElmentP struct {
	*noAliasMessageElement
	*ChildrenMessageElement
}

func (e *MessageElmentP) Tag() string {
	return "p"
}

func (e *MessageElmentP) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElmentP) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElmentP{
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementMessage struct {
	Id      string
	Forward bool
	*noAliasMessageElement
	*ChildrenMessageElement
}

func (e *MessageElementMessage) Tag() string {
	return "message"
}

func (e *MessageElementMessage) Stringify() string {
	result := ""
	if e.Id != "" {
		result += ` id="` + e.Id + `"`
	}
	if e.Forward {
		result += ` forward`
	}
	childrenStr := e.stringifyChildren()
	if childrenStr == "" {
		return `<` + e.Tag() + result + ` />`
	}
	return `<` + e.Tag() + result + `>` + childrenStr + `</` + e.Tag() + `>`
}

func (e *MessageElementMessage) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementMessage{
		Forward: false,
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}
	if id, ok := attrMap["id"]; ok {
		result.Id = id
	}
	if forwardAttr, ok := attrMap["forward"]; ok {
		result.Forward = forwardAttr == "" || forwardAttr == "true" || forwardAttr == "1"
	}
	return result, nil
}

func init() {
	RegsiterParserElement(&MessageElmentBr{})
	RegsiterParserElement(&MessageElmentP{})
	RegsiterParserElement(&MessageElementMessage{})
}
