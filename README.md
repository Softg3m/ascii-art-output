# ASCII Art Output (Go)

## Description

This project is an extension of the ASCII-art program written in Go.
It converts input text into ASCII art using banner files and supports writing the output into a file using a command-line flag.

The program reads characters from a banner file (e.g. `standard.txt`, `shadow.txt`, `thinkertoy.txt`) and prints or saves the corresponding ASCII representation.


## Features

* Convert text to ASCII art
* Support multiple banner styles:

  * standard
  * shadow
  * thinkertoy
* Output result to terminal
* Save output to a file using `--output` flag
* Handle new lines (`\n`) in input


##  Usage

### 1. Run normally (print to terminal)

```bash
go run . "hello"
```

```bash
go run . "hello" standard
```


### 2. Save output to file

```bash
go run . --output=banner.txt "hello" standard
```

Then view the file:

```bash
cat banner.txt
```


## Flag Format

The output flag must follow this exact format:

```bash
--output=<fileName.txt>
```

### Invalid Examples

```bash
--output banner.txt
-output=banner.txt
--output=
```

If the format is incorrect, the program displays:

```
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```



##  Project Structure

```
├── main.go
├── main_test.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
└── README.md
```

---

## How It Works

* Each character is represented by 8 lines in the banner file
* Characters are indexed using ASCII values
* The program calculates the correct position in the banner file and builds the output line by line
* If `--output` is provided, the result is written to a file instead of printed



## Example

```bash
go run . --output=output.txt "Hi" standard
```

Output (in `output.txt`):

```

 _    _   _  
| |  | | (_) 
| |__| |  _  
|  __  | | | 
| |  | | | | 
|_|  |_| |_| 
```

##  Error Handling

The program handles:

* Missing arguments
* Invalid flag format
* Invalid banner files
* File writing errors


## Packages

Only standard Go packages are used:

* fmt
* os
* strings

