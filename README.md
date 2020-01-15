<p align="center">
  <b>
    <span style="font-size:larger;">conncheck-go</span>
  </b>
  <br />
   <a href="https://travis-ci.org/detailyang/conncheck-go"><img src="https://travis-ci.org/detailyang/conncheck-go.svg?branch=master" /></a>
   <a href="https://ci.appveyor.com/project/detailyang/conncheck-go"><img src="https://ci.appveyor.com/api/projects/status/ux7lf3h9wf8bx8ep?svg=true" /></a>
   <br />
   <b>conncheck-go checks whether the connection was closed or not</b>
</p>

```bash
go test -v -benchmem -run="^$" github.com/detailyang/conncheck-go -bench "^Benchmark"
goos: darwin
goarch: amd64
pkg: github.com/detailyang/conncheck-go
BenchmarkCheck-8   	  782926	      1534 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	github.com/detailyang/conncheck-go	2.166s
```
