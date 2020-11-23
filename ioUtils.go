package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * TODO: document and unit test
 */
func initFile(filePath string) *os.File {
	stdOut, err := os.Create(filePath)
	checkError(err)
	defer stdOut.Close()
	return stdOut
}

/*
 * TODO: document and unit test
 */
func initReader() *bufio.Reader {
	return bufio.NewReaderSize(os.Stdin, 1024*1024)
}

func initWriter() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

/*
 * TODO: document and unit test
 */
func readLine(reader *bufio.Reader) string {
	readLine, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	safeString := string(readLine)
	safeString = strings.TrimRight(safeString, "\r\n")
	safeString = strings.TrimSpace(safeString)
	return safeString
}

/*
 * TODO: document and unit test
 */
func readLines(reader *bufio.Reader, splitter string) []string {
	return strings.Split(readLine(reader), splitter)
}

/*
 * TODO: document and unit test
 */
func readLineAsString(reader *bufio.Reader) string {
	fmt.Println("Please enter a string:")
	return readLine(reader)
}

/*
 * TODO: document and unit test
 */
func readLineAsInt(reader *bufio.Reader) int {
	fmt.Println("Please enter an integer:")
	readLine, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	return int(readLine)
}

/*
 * TODO: document and unit test
 */
func readLineAsInts(reader *bufio.Reader, splitter string) []int {
	fmt.Println("Please enter an integer array deliniated by spaces:")
	readStrings := readLines(reader, splitter)
	var parsedInts []int
	for i := 0; i < len(readStrings); i++ {
		readString, err := strconv.ParseInt(readStrings[i], 10, 64)
		checkError(err)
		parsedInt := int(readString)
		parsedInts = append(parsedInts, parsedInt)
	}
	return parsedInts
}

/*
 * TODO: document and unit test
 */
func readLineAsStrings(reader *bufio.Reader, splitter string) []string {
	fmt.Println("Please enter a string array deliniated by spaces:")
	return readLines(reader, splitter)
}

/*
 * TODO: document and unit test
 */
func writeFile(stdOut *os.File, result string) {
	writer := bufio.NewWriterSize(stdOut, 1024*1024)
	fmt.Fprintf(writer, result)
	writer.Flush()
}
