package parser

import (
	"strings"
)

func Parse(policy string) string {
	policy = strings.ToLower(policy)

	var result string

	if (strings.Index(policy, "content-security") == 0) {
		pos := strings.Index(policy, ":")
		key := policy[0:pos]
		policy = policy[pos + 1:len(policy)]
		result += key + ":\n\n"
	}

	directives := strings.Split(policy, ";")

	for _, directive := range directives {
		directive = strings.Trim(directive, " ")
		parts := strings.Split(directive, " ")
		var padding string

		for i, part := range parts {
			if (i > 0) {
				padding = "    "
			} else {
				padding = "  "
			}
			result += padding + part + "\n"
		}
		result += "\n"
	}

	return result
}