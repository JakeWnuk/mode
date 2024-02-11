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
Include items from a `file` or `URL`
```
$ cat keep.txt
string

$ mode -c -s -w keep.txt example.txt
1 string
```
Exclude items from a `file` or `URL`
```
$ cat remove.txt
string

$ mode -c -s -v remove.txt example.txt
2 this
1 is
1 A
1 test
```

### Include Items
Mode can be used to aggregate and sort input from `stdin`, `files`, and `URLs`
by frequency. This is the default mode and when the `-w` flag is provided items
that are in both `stdin` and the `-w` item are printed.

This option is often used to make join files quickly and multiple `-w`
and `-v` flags can be used together at once.

### Exclude Items
Mode can be used to aggregate and sort input from `stdin`, `files`, and `URLs`
by frequency. This is the default mode and when the `-v` flag is provided items
that are in `stdin` and not the `-v` item are printed.

This option is often used to make difference files quickly and multiple `-w`
and `-v` flags can be used together at once.
