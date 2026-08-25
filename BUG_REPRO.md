# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
?   	example.com/casescript/cmd/casescript	[no test files]
--- FAIL: TestCaseListStableForSameDate (0.02s)
    case_list_regression_test.go:33: unexpected stable order: []domain.LegalCase{domain.LegalCase{ID:"case-a", Number:"2024-002", Title:"第二案", Summary:"同日资料二", PublishDate:"2024-06-01", Status:"published", CreatedAt:"2024-05-01"}, domain.LegalCase{ID:"case-z", Number:"2024-001", Title:"第一案", Summary:"同日资料一", PublishDate:"2024-06-01", Status:"published", CreatedAt:"2024-05-01"}}
FAIL
FAIL	example.com/casescript	0.017s
ok  	example.com/casescript/cases	0.018s
ok  	example.com/casescript/cli	0.009s
ok  	example.com/casescript/content	0.011s
ok  	example.com/casescript/domain	0.001s
ok  	example.com/casescript/publish	0.024s
ok  	example.com/casescript/store	0.019s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/casescript): exit `0`
- Frontend build (web): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/casescript): exit `0`
- Frontend build (web): exit `0`
