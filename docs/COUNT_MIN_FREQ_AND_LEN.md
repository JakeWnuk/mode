### Quick Start
Files used
```
$ cat example.txt
this
this
is
A
test string
```
Counting frequency
```
$ mode -c example.txt
2 this
1 is
1 A
1 test string
```
Minimum frequency
```
$ mode -m 2 example.txt
2 this
```
Removing items under length
```
$ mode -x 2 -c example.txt
2 this
1 test string
```

### Counting Frequency
Mode can be used to aggregate and sort input from `stdin`, `files`, and `URLs`
by frequency. This is the default mode and when the `-c` flag is provided the
corresponding item counts are displayed.

### Minimum Frequency
Mode can be used to aggregate and sort input from `stdin`, `files`, and `URLs`
by frequency. This is the default mode and when the `-m` flag is provided the
minimum frequency to display can be changed (default all).

### Remove Items Under Length
Mode can be used to aggregate and sort input from `stdin`, `files`, and `URLs`
by frequency. This is the default mode and when the `-x` flag is provided items
with length above the `N` value provided will be displayed.

