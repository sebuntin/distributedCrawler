package engine

type BookParseFuncConversion interface {
	Serialize(string) BookParseFunc
	Deserialize(BookParseFunc) string
}

type BookParseFunc func([]byte, string, string) ParserResult

type BookParser struct {
	FunName  string
	BookName string
	parse    BookParseFunc
}

func (b *BookParser) Parse(content []byte, url string) ParserResult {
	return b.parse(content, url, b.BookName)
}

func (b *BookParser) Serialize() (string, []interface{}) {
	return b.FunName, []interface{}{b.BookName}
}

func NewBookParser(funcName, bookName string, parseFunc BookParseFunc) *BookParser {
	return &BookParser{
		FunName:  funcName,
		BookName: bookName,
		parse:    parseFunc,
	}
}
