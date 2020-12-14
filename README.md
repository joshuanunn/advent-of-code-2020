## :christmas_tree: Advent of Code 2020 in Go :christmas_tree:
These are my attempts at [Advent-of-Code](https://adventofcode.com) 2020 in Go. Benchmarks, tests and improvements will be added over time for each day.

## Benchmarks
Benchmarks have been included to test the functions for parts 1 and 2 for each day - input file parsing has not been included as part of the tests. Quoted benchmark times are from a laptop (i7-10510U). Benchmarks can be run for any day from the base directory:
```
$ go test ./day10_adapter_array -bench=.
```
| Day | Part 1 | Part 2 |
| :---: | :---: | :---: |
| 3 | 2.2 μs | 10 μs |
| 4 | 966 μs | 1.1 ms |
| 5 | 952 ns | 965 ns |
| 6 | 961 ns | 953 ns |
| 7 | 39 ms | 7.8 μs |
| 8 | 17 μs | 2.2 ms |
| 9 | 150 μs | 3.5 ms |
| 10 | 3.4 μs | 207 μs |
| 11 | 15 ms | 46 ms |
| 12 | 23 μs | 25 μs |
| 13 | 1.8 μs | 5.1 μs |

## Tests
Tests related to the specific input.txt can be run for each day from the base directory:
```
$ go test ./day10_adapter_array
```

## License
[MIT](LICENSE)  

Note that this repo contains example inputs for each puzzle day (input.txt), which have been generated on [Advent-of-Code](https://adventofcode.com) - these sample data files belong to [Eric Wastl](https://twitter.com/ericwastl).
