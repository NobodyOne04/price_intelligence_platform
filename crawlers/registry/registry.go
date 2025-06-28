package registry

import (
	"crawlers/amazon"
	"crawlers/common"
)

var parsers = map[string]common.Parser{
	"amazon": amazon.NewParser(),
}

func Get(name string) (common.Parser, bool) {
	parser, ok := parsers[name]
	return parser, ok
}

func All() []common.Parser {
	out := make([]common.Parser, 0, len(parsers))
	for _, parser := range parsers {
		out = append(out, parser)
	}
	return out
}

