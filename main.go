package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type data struct {
	points int
	sheet  map[string]int
}

var (
	file     *string
	duration *int
)

func init() {
	file = flag.String("test", "problems.csv", "test filepath")
	duration = flag.Int("time", 10, "test duration in seconds")
}

func main() {

	flag.Parse()

	// Read the csv answer sheet
	records := readCSV(*file)

	// take each line of the csv and turn into a map with actual num
	sheet := make(map[string]int)
	for _, line := range records {
		num, _ := strconv.Atoi(line[1])
		sheet[line[0]] = num
	}

	quiz := data{
		points: 0,
		sheet:  sheet,
	}

	ch := make(chan int)

	go startQuiz(&quiz, ch)
	go timer(*duration, ch)

	<-ch

	fmt.Printf("You scored %v out of %v!\n", quiz.points, len(quiz.sheet))
}

func timer(delay int, ch chan<- int) {
	time.Sleep(time.Second * time.Duration(delay))
	ch <- 1
}

func startQuiz(quiz *data, ch chan<- int) {

	reader := bufio.NewReader(os.Stdin)

	i := 1
	for k, v := range quiz.sheet {

		fmt.Printf("Problem #%v: %v = ", i, k)
		text, _ := reader.ReadString('\n')

		answ, _ := strconv.Atoi(strings.Trim(text, "\n "))

		if answ == v {
			quiz.points++
		}

		i++
	}
	ch <- 1
}

func readCSV(file string) [][]string {
	// opens the file
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln("Unable to read input file:", file, err)
	}
	defer f.Close()

	// reads the the contents of the file as csv
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse file as CSV for:", file, err)
	}

	return records
}
