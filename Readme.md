# Benchmark with Formatter strings
Project aiming to benchmark the ways to format strings in golang

## How to run

```bash
$: go test -bench=. -count 3
```

## Small strings

TestCase used [`BenchmarkFormatSmallStrings`](./benchmark_test.go)
```bash
goos: linux
goarch: amd64
pkg: benchmarks
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkFormatSmallStrings/concat_strings_with_plus(+)_operator-12             1000000000               0.0000025 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_plus(+)_operator-12             1000000000               0.0000026 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_plus(+)_operator-12             1000000000               0.0000022 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_fmt()_func-12                   1000000000               0.0000029 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_fmt()_func-12                   1000000000               0.0000032 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_fmt()_func-12                   1000000000               0.0000029 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_text/template_package-12        1000000000               0.0000301 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_text/template_package-12        1000000000               0.0000296 ns/op
BenchmarkFormatSmallStrings/concat_strings_with_text/template_package-12        1000000000               0.0000227 ns/op
PASS
ok      benchmarks      0.015s
```

## Want to contribute?
Open your PR and add new benchmarks in this repository