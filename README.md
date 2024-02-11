`Mode` is a program for quickly aggregating and frequency sorting text from multiple sources and supports concurency.
- Tokenizes text with a lean for password cracking


### Getting Started

Usage information and other documentation can be found below:

- Usage documentation:
    - [Counting Frequency and Filtering Minimum Frequency & Length](https://github.com/JakeWnuk/mode/blob/main/docs/COUNT_MIN_FREQ_AND_LEN.md)
    - [Only Include or Remove Items](https://github.com/JakeWnuk/mode/blob/main/docs/INCLUDE_AND_REMOVE.md)

Mode supports input from `stdin`, `files`, and `URLs`.
- Pass multiple items as items after flags

### Install from Source
```
go install github.com/JakeWnuk/mode@latest
```
### Current Version 0.0.0:
```
Usage of Mode version (0.0.0):

mode [options] [URLS/FILES]
Accepts standard input and/or additonal arguments.

Options:
  -c    Show the frequency count of each item
  -m int
        Minimum frequency to include in output. Value should be an integer.
  -v value
        Only include items not in a file.
  -w value
        Only include items in a file.
  -x int
        Exclude items less than or equal to a length from output. Length should be an integer.
```
