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
| 7 | 34 μs | 41 ms | 1.9 ms |
| 8 | 20 μs | 17 μs | 2.3 ms |
| 9 | 27 μs | 152 μs | 3.7 ms |
| 10 | 5.4 μs | 4.7 μs | 214 μs |
| 11 | 8.3 μs | 16 ms | 46 ms |
| 12 | 22 μs | 237 μs | 255 μs |
| 13 | 2.6 μs | 1.9 μs | 5.1 μs |
| 14 | 22 μs | 382 μs | 15 ms |
| 15 | - | 11 μs | 754 ms |
| 16 | 15 μs | 1.5 ms | 1.9 ms |
| 17 | 2.8 μs | 142 μs | 9.1 s |
| 18 | 18 μs | 1.6 ms | 1.7 ms |
| 19 | | | |
| 20 | | | |
| 21 | 9.9 μs | 357 μs | 296 μs |
| 22 | 4.1 μs | 47 μs | 712 ms |
| 23 | | | |
| 24 | 16 μs | 327 μs | 1.1 s |

## Tests
Tests related to the specific input.txt can be run for each day from the base directory:
```
$ go test ./day10_adapter_array
```

## License
[MIT](LICENSE)  

Note that this repo contains example inputs for each puzzle day (input.txt), which have been generated on [Advent-of-Code](https://adventofcode.com) - these sample data files belong to [Eric Wastl](https://twitter.com/ericwastl).
