`Mode` is a program for quickly aggregating and frequency sorting text from multiple sources and supports concurency.
- Tokenizes text with a lean for password cracking


### Getting Started

Usage information and other documentation can be found below:

- Usage documentation:
    - [Counting Frequency and Filtering Minimum Frequency & Length](https://github.com/JakeWnuk/mode/blob/main/docs/COUNT_MIN_FREQ_AND_LEN.md)
    - [Only Include or Remove Items](https://github.com/JakeWnuk/mode/blob/main/docs/INCLUDE_AND_REMOVE.md)

Mode supports input from `stdin`, `files`, and `URLs`.
- Pass multiple items as items after flags

Mode fits into a small tool ecosystem for password cracking and is designed for lightweight and easy usage with its companion tools:

- [maskcat](https://github.com/JakeWnuk/maskcat)
- [rulecat](https://github.com/JakeWnuk/rulecat)
- [mode](https://github.com/JakeWnuk/mode)

### Install from Source
```
go install github.com/jakewnuk/mode@v0.0.1
```
### Current Version 0.0.1:
```
Usage of Mode version (0.0.1):

mode [options] [URLS/FILES] [...]
Accepts standard input and/or additonal arguments.

Options:
  -a    Perform additional parsing of each item
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
