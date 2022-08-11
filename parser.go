package crondescriptor

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

// Parse the cron expression string
// return:A 7 part string array,one part for each component
//        of the cron expression (seconds, minutes, etc.)
func (p *Parser) Parse() {}

// NormalizeExpression Converts cron expression components into consistent,
// predictable formats.
// return: expression_parts: A 7 part string array, one part for each component
//         of the cron expression
func (p *Parser) NormalizeExpression() {}
