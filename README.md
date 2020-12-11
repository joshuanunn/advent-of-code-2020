## :christmas_tree: Advent of Code 2020 in Go :christmas_tree:
These are my attempts at [Advent-of-Code](https://adventofcode.com) 2020 in Go. Benchmarks, tests and improvements will be added over time for each day.

## Benchmarks
Benchmarks have been included to test the functions for parts 1 and 2 for each day - input file parsing has not been included as part of the tests. Quoted benchmark times are from a laptop (i7-10510U). Benchmarks can be run for any day from the base directory:
```
$ go test ./day10_adapter_array -bench=.
```
| Day | Part | Time | Apparent Complexity (approx) |
| :---: | :---: | :---: | :---: |
| 3 | 1 | 2.2 μs ||
| 3 | 2 | 10 μs ||
| 4 | 1 | 966 μs ||
| 4 | 2 | 1.1 ms ||
| 5 | 1 | 952 ns ||
| 5 | 2 | 965 ns ||
| 6 | 1 | 961 ns ||
| 6 | 2 | 953 ns ||
| 7 | 1 | 39 ms ||
| 7 | 2 | 7.8 μs ||
| 8 | 1 | 17 μs ||
| 8 | 2 | 2.2 ms ||
| 9 | 1 | 150 μs ||
| 9 | 2 | 3.5 ms ||
| 10 | 1 | 3.4 μs | ```O(n*log(n)``` |
| 10 | 2 | 207 μs | ```O((n^2)*log(n))``` |

## Tests
Tests related to the specific input.txt can be run for each day from the base directory:
```
$ go test ./day10_adapter_array
```

## License
[MIT](LICENSE)  

Note that this repo contains example inputs for each puzzle day (input.txt), which have been generated on [Advent-of-Code](https://adventofcode.com) - these sample data files belong to [Eric Wastl](https://twitter.com/ericwastl).
