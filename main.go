package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type data struct {
	points int
	sheet  map[string]string
}

var (
	file     *string
	duration *int
)

func init() {
	file = flag.String("test", "problems.csv", "path to test file")
	duration = flag.Int("time", 10, "test duration in seconds")
}

func main() {

	flag.Parse()

	records := readCSV(*file)

	sheet := make(map[string]string)
	for _, line := range records {
		sheet[line[0]] = line[1]
	}

	quiz := data{
		points: 0,
		sheet:  sheet,
	}

	ch := make(chan int)

	fmt.Println("Quiz has started!")
	fmt.Printf("Total time: %v seconds\n", *duration)
	go startQuiz(&quiz, ch)
	go timer(*duration, ch)

	<-ch

	fmt.Printf("\nYou scored %v out of %v!\n", quiz.points, len(quiz.sheet))
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

		// Allows to read from windows terminals
		//text = strings.Replace(text, "\r\n", "", -1)
		text = strings.Trim(text, "\r\n ")
		
		//fmt.Println("\""+ text + "\"")
		if text == v {
			//fmt.Println("point added")
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
