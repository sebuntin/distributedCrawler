package engine

// Request 定义解析器的输出类
type Request struct {
	CurURL string
	RefURL string
	Args   []string
	Parser Parser
}

type Parser interface {
	Parse([]byte, ...string) ParserResult
	Serialize() (string, []string)
}

type ParseFunConversion interface {
	Serialize(Request) (string, []string)
	Deserialize(string) (ParseFunc, error)
}

type ParseFunc func([]byte, ...string) ParserResult

type FuncParser struct {
	name  string
	args  []string
	parse ParseFunc
}

func (f *FuncParser) Parse(content []byte, url ...string) ParserResult {
	return f.parse(content, url...)
}

func (f *FuncParser) Serialize() (string, []string) {
	return f.name, f.args
}

func NewFuncParser(name string, args []string, p ParseFunc) *FuncParser {
	return &FuncParser{
		name:  name,
		parse: p,
		args:  args,
	}
}

// ParserResult 定义解析器的输出结果
type ParserResult struct {
	Requests []Request
	Items    []interface{}
}

type Item struct {
	Url     string
	Id      string
	PayLoad interface{}
}

type NilParser struct {
}

func (n NilParser) Parse(_ []byte, _ string) ParserResult {
	return ParserResult{}
}

func (n NilParser) Serialize() (string, []string) {
	return "NilParser", []string{}
}

//func NilParser([]byte) ParserResult {
//	return ParserResult{}
//}
