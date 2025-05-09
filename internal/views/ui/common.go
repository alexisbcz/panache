package ui

import "strings"

type CommonProps struct {
	Class string
}

// mergeClasses merges class names, removing duplicates and trimming spaces.
func mergeClasses(classes ...string) string {
	seen := make(map[string]bool)
	var result []string

	for _, group := range classes {
		if group == "" {
			continue
		}

		for _, class := range strings.Fields(group) {
			if !seen[class] {
				seen[class] = true
				result = append(result, class)
			}
		}
	}

	return strings.Join(result, " ")
}
