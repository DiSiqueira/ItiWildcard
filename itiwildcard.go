package itiwildcard

import (
	"net/http"
	"strings"
)

// WildcardMatcher Store the template to match with the request
type WildcardMatcher struct {
	template string
}

// New is the constructor to ItiWildcard
func New(template string) *WildcardMatcher {
	return &WildcardMatcher{
		template: template,
	}
}

// Match returns if the request can be handled by this Route.
func (t *WildcardMatcher) Match(req *http.Request) bool {

	if req.URL.Path == t.template {
		return true
	}

	req.URL.Path = strings.Trim(req.URL.Path, "/")
	t.template = strings.Trim(t.template, "/")

	pText := strings.Split(req.URL.Path, "/")
	pPattern := strings.Split(t.template, "/")

	if len(pText) != len(pPattern) {
		return false
	}

	for i := range pPattern {
		if pPattern[i] == "*" {
			continue
		}

		if pPattern[i] != pText[i] {
			return false
		}
	}

	return true
}
