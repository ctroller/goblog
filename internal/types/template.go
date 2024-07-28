package types

import "html/template"

type JSScript struct {
	Link     string
	Content  template.JS
	Async    bool
	Footer   bool
	UniqueId string
}

type CSSFile struct {
	Link string
}
