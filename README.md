# test.log parser

## Usage

Generates list of individual test to run from `test.log`:

```
$ cat test.log | go run main.go
Thought it would be a failed test, but formatted strangely: --- FAIL	aseuoth
Thought it would be a failed package, but formatted strangely: FAIL bobob
go test github.com/hashicorp/consul/agent -run TestRexecWriterAB
go test github.com/hashicorp/consul/agent -run TestRexecWriter
go test github.com/hashicorp/consul/agent -run TestCoordinate_Node
go test github.com/hashicorp/consul/watch -run TestAE_Run_QuitMORESPACES
go test github.com/hashicorp/consul/watch -run TestAE_Run_Quit
```

or when you have it in your `PATH` and want to run the test right away:

```
$ cat test.log | testlog | bash
```

## Features

Errors are written to STDERR and it supports a flag for `-ent`.
