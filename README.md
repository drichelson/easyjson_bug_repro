# easyjson_bug_repro
repo for reproducing easyjson behavior when importing a package whose init code emits to stdout 


To reproduce:
1.  `easyjson -leave_temps -all example/example.go`
2. Observe error: `Bootstrap failed: 1:29: expected ';', found from (and 5 more errors)`
3. Observe the output from [here](bad_actor/side_effect.go:13) in the first line of [example/example_easyjson.go.tmp](example/example_easyjson.go.tmp) file 


I checked in the temp files so it's easier to see the issue without checking out and running this code.