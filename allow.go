package http405

import (
	"slices"
	"strings"
)

// allow return the value of the "Allow" HTTP-header from a list of methods.
func allow(methods ...string) string {
	var a []string
	for _, method := range methods {
		method = strings.ToUpper(method)

		// Just in case the user put a common (",") in a method.
		// We assume this means they put multiple methods in there.
		pieces := strings.Split(method, ",")
		for _, piece := range pieces {
			piece = strings.TrimSpace(piece)
			if "" == piece {
				continue
			}
			a = append(a, piece)
		}
	}

	slices.Sort(a)

	return strings.Join(a, ", ")
}
