# test.log parser

Generates list of individual test to run from enterprise `test.log`:

```
$ cat test.log | go run main.go
testlog git:(master) ✗ cat test.log | go run main.go
Thought it would be a failed test, but formatted strangely: --- FAIL	aseuoth
Thought it would be a failed package, but formatted strangely: FAIL bobob
go test -tags 'ent prem' github.com/hashicorp/consul/agent -run TestRexecWriterAB
go test -tags 'ent prem' github.com/hashicorp/consul/agent -run TestRexecWriter
go test -tags 'ent prem' github.com/hashicorp/consul/agent -run TestCoordinate_Node
go test -tags 'ent prem' github.com/hashicorp/consul/watch -run TestAE_Run_QuitMORESPACES
go test -tags 'ent prem' github.com/hashicorp/consul/watch -run TestAE_Run_Quit
```

Errors are written to STDERR.
