# Text Transformer

A powerful command-line text processing tool written in Go that applies a set of smart transformations to input text according to specific markup rules.

## Features

This tool processes text with the following transformations:

- `(hex)` → converts the previous word (hexadecimal) to decimal  
- `(bin)` → converts the previous word (binary) to decimal  
- `(up)` → converts the previous word to **UPPERCASE**  
- `(low)` → converts the previous word to **lowercase**  
- `(cap)` → converts the previous word to **Capitalized**  
- `(up, n)`, `(low, n)`, `(cap, n)` → applies the transformation to the **n previous words**  
- Smart **a/an** replacement based on the next word  
- Proper punctuation spacing (`. , ! ? : ;`) — no space before, one space after  
- Correct handling of ellipses `...` and combined punctuation like `?!`, `!?`  
- Proper placement of single quotes `'text'` and multi-word quoted phrases

## Usage

```bash
go run . <input-file> <output-file>
Examples
Bash$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.

$ go run . sample.txt result.txt

$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
Bash$ echo "Simply add 42 (hex) and 10 (bin) and you will see the result is 68." > sample.txt
$ go run . sample.txt result.txt
$ cat result.txt
Simply add 66 and 2 and you will see the result is 68.
Bash$ echo "There it was. A amazing rock!" > sample.txt
$ go run . sample.txt result.txt
$ cat result.txt
There it was. An amazing rock!
Bash$ echo "Punctuation tests are ... kinda boring ,what do you think ?" > sample.txt
$ go run . sample.txt result.txt
$ cat result.txt
Punctuation tests are... kinda boring, what do you think?
Bash$ echo "As Elton John said: ' I am the most well-known homosexual in the world '" > sample.txt
$ go run . sample.txt result.txt
$ cat result.txt
As Elton John said: 'I am the most well-known homosexual in the world'
```
## Requirements

Go 1.18+
Only standard library packages are used

## Project Structure
```bash
├── convert.go       # conversion function
├── fixarticle.go    # function of changing the article
├── fixpunct.go      # punctuation editing function 
├── fixquotes.go     # quote editing function
├── go.mod
├── main.go          # testing
├── README.md          
├── sample.txt.      # input
└── result.txt       # output
