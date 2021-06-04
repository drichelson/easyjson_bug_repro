package example

import (
	_ "github.com/drichelson/easyjson_bug_repro/bad_actor"
	_ "github.com/mailru/easyjson/gen"
)

// easyjson:json
type A struct{}
