<h1 align="center">
Mode
 </h1> 

 `Mode` is a program for taking `stdin` and aggregating by frequency to `stdout`.

- Display words with and without item count
- Parse `stdin` to n-grams by spaces and add to count
- Parse `stdin` by a custom dictionary and add to count
- All flags can be used together for more granular parsing

## Getting Started

- [Install](#Install)
- [Count Frequency](#Count-Frequency)
- [Parse Text](#Parse-Text)
- [Custom Dictionary](#Custom-Dictionary)
- [Speed Test](#Speed-Test)

### Install
```
go install -v github.com/jakewnuk/mode@latest
```
```
Usage of mode (version 1.0.0):
  -c    Display the frequency count of each item
        Example: mode -c
  -f string
        Parse items from a dictionary file and add to count. The file should contain one item per line.
        Example: mode -f dict.txt
  -s    Split items into n-grams by spaces and add to count.
        Example: mode -s
```
```
$ cat test.txt
hello
hello
Test1234
hello World!
TestWorld
Testing This Hello
```

## Count Frequency
- Display words with and without item count
```
$ cat test.txt | mode -c
2 hello
1 Test1234
1 hello World!
1 TestWorld
1 Testing This Hello
```

## Parse Text
- Parse `stdin` to tokens by spaces and add to count
```
$ cat test.txt | mode -s -c
5 hello
2 TestWorld
2 Test1234
1 Testing This Hello
1 Hello
1 hello World!
1 Testing
1 This
1 World!
```

## Custom Dictionary
- Parse `stdin` by a custom dictionary and add to count
- Attempts to match largest substring first then stops search after match
```
$ cat dict.txt
World
Test

$ cat test.txt | mode -c -f dict.txt
3 Test
2 hello
1 hello World!
1 TestWorld
1 Testing This Hello
1 Test1234
1 World
```
- Input is sorted by length then checked
- The largest token is returned first from the first position
- See example where `Testing` is added to `dict.txt`:
```
$ cat dict.txt
World
Test
Testing

$ cat test.txt | mode -c -f dict.txt
2 hello
2 Test
1 Test1234
1 World
1 hello World!
1 TestWorld
1 Testing
1 Testing This Hello
```
- Checking the first position means equal size text will return based on
  position
```
$ cat dict.txt
World
hello

$ cat test.txt | mode -f dict.txt
5 hello
1 Testing This Hello
1 Test1234
1 hello World!
1 World
1 TestWorld
```

## Speed Test
- Comparing speeds with `sort | uniq -c | sort -rn`
- Compare Command: `LC_ALL=C sort -T ./ $1 | uniq -c | LC_ALL=C sort -T ./ -rn`

#### Generation
```
$ base64 /dev/urandom | head -c $((250 * 1024 * 1024)) > 250M.txt
$ base64 /dev/urandom | head -c $((1024 * 1024 * 1024)) > 1G.txt
$ base64 /dev/urandom | head -c $((5 * 1024 * 1024 * 1024)) > 5G.txt
```

#### Comparing Speeds
- 250MB
    - Mode: `4.681`
    - SortUniq: `3.148`
```
$ cat 250M.txt | (time (LC_ALL=C sort -T ./ | uniq -c | LC_ALL=C sort -T ./ -rn)) | grep -o '[0-9.]* total'
( LC_ALL=C sort -T ./ | uniq -c | LC_ALL=C sort -T ./ -rn; )  2.61s user 1.00s system 114% cpu 3.148 total
$ cat 250M.txt| (time (./mode -c)) | grep -o '[0-9.]* total'
( ./mode -c; )  3.41s user 1.43s system 103% cpu 4.681 total
```
- 1GB
    - Mode: `22.698`
    - SortUniq: `13.188`
```
$ cat 1G.txt | (time (LC_ALL=C sort -T ./ | uniq -c | LC_ALL=C sort -T ./ -rn)) | grep -o '[0-9.]* total'
( LC_ALL=C sort -T ./ | uniq -c | LC_ALL=C sort -T ./ -rn; )  10.87s user 4.33s system 115% cpu 13.188 total
$ cat 1G.txt| (time (./mode -c)) | grep -o '[0-9.]* total'
( ./mode -c; )  18.44s user 6.08s system 108% cpu 22.698 total
```
- 5GB
    - Mode: `1:57.33`
    - SortUniq: `1:21.07`
```
$ cat 5G.txt | (time (LC_ALL=C sort -T ./ | uniq -c | LC_ALL=C sort -T ./ -rn)) | grep -o '[0-9.]* total'
( LC_ALL=C sort -T ./ | uniq -c | LC_ALL=C sort -T ./ -rn; )  60.66s user 28.76s system 110% cpu 1:21.07 total
$ cat 5G.txt| (time (./mode -c)) | grep -o '[0-9.]* total'
( ./mode -c; )  101.28s user 30.63s system 112% cpu 1:57.33 total
```
