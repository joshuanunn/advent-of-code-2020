## :christmas_tree: Advent of Code 2020 in Go :christmas_tree:
These are my attempts at [Advent-of-Code](https://adventofcode.com) 2020 in Go. Benchmarks, tests and improvements will be added over time for each day.

## Benchmarks
Benchmarks have been included to test the initial input file parsing to []string, and the part 1 and 2 function for each day. Quoted benchmark times are from a laptop (i7-10510U). Benchmarks can be run for any day from the base directory:
```
$ go test ./day10_adapter_array -bench=.
```
| Day | Read | Part 1 | Part 2 |
| :---: | :---: | :---: | :---: |
| 1 | 8 μs | 11 μs | 424 μs |
| 2 | 31 μs | 819 μs | 797 μs |
| 3 | 14 μs | 18 μs | 27 μs |
| 4 | 87 μs | 961 μs | 1.1 ms |
| 5 | 39 μs | 386 μs | 488 μs |
| 6 | 62 μs | 78 μs | 78 μs |
| 7 |  | 39 ms | 7.8 μs |
| 8 |  | 17 μs | 2.2 ms |
| 9 |  | 150 μs | 3.5 ms |
| 10 |  | 3.4 μs | 207 μs |
| 11 |  | 15 ms | 46 ms |
| 12 |  | 23 μs | 25 μs |
| 13 |  | 1.8 μs | 5.1 μs |

## Tests
Tests related to the specific input.txt can be run for each day from the base directory:
```
$ go test ./day10_adapter_array
```

## License
[MIT](LICENSE)  

Note that this repo contains example inputs for each puzzle day (input.txt), which have been generated on [Advent-of-Code](https://adventofcode.com) - these sample data files belong to [Eric Wastl](https://twitter.com/ericwastl).
