
`go test -v -count=1 -timeout=30s -memprofile profile.mem.out.old -tags bench .`

`go tool pprof -http 127.0.0.1:8989 hw10_program_optimization.test profile.mem.out.old`

`go test -v -count=1 -timeout=30s -cpuprofile profile.cpu.out.old -tags bench .`

`go tool pprof -http 127.0.0.1:8989 hw10_program_optimization.test profile.cpu.out.old`

