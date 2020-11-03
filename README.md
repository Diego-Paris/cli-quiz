# CLI Quiz

Topics covered in this project: `command line arguments` `goroutines` `channels` `csv files` `flags` `string manipulation` `user input` `timed`

## Summary
A timed quiz given as a Command Line Interface (CLI), reads questions and answers from a CSV file. Uses concurrent processes to take in user input and keeps track of remaining time. The test is concluded once the user is finished or the time is up.

## Requirements

Go version 1.15.x or higher

For more information about installing Go, visit [golang.org/doc/install](https://golang.org/doc/install)

>Note: May work with older versions but is not guaranteed to function properly.

## Installation
```bash
git clone https://github.com/Diego-Paris/cli-quiz
```

## Flags
```
  -test string
        path to test file (default "problems.csv")
  -time int
        quiz duration in seconds (default 10)
  -help
        list all flags and descriptions
```

## Usage
User begins the quiz and can enter their answer using the Enter key.
> Note: Using time flag example
```
$ go run main.go -time 30
Quiz has started!
Total time: 30 seconds
Problem #1: 2+4 = 
...
```

### Answer sheet example
> Note: Project only reads csv files
```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

## Contributing
Pull requests are welcome.

## License
[MIT](https://github.com/Diego-Paris/cli-quiz/blob/master/LICENSE)