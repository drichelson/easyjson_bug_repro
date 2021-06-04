// +build ignore

// TEMPORARY AUTOGENERATED FILE: easyjson bootstapping code to launch
// the actual generator.

package main

import (
  "fmt"
  "os"

  "github.com/mailru/easyjson/gen"

  pkg "github.com/drichelson/easyjson_bug_repro/example"
)

func main() {
  g := gen.NewGenerator("example_easyjson.go")
  g.SetPkg("example", "github.com/drichelson/easyjson_bug_repro/example")
  g.Add(pkg.EasyJSON_exporter_A(nil))
  if err := g.Run(os.Stdout); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}