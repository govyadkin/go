package main

import (
	"flag"
	"hw1.1/hw1.1/unique"
	"hw1.1/hw1.1/workWithFiles"
	"log"
)

var flagCount = flag.Bool("c", false, "count the number of occurrences of a string in the input. Print this number before the line, separated by a space.")
var flagDuplicate = flag.Bool("d", false, "print only those lines that are repeated in the input.")
var flagUnique = flag.Bool("u", false, "print only those lines that are not repeated in the input.")
var flagField = flag.Int("f", 0, "ignore the first num_fields fields in the line. A field in a string is a non-empty character set separated by a space.")
var flagChar = flag.Int("s", 0, "ignore the first num_chars characters in the string. When used in conjunction with the -f option, the first characters after the num_fields fields are counted (excluding the space separator after the last field).")
var flagRegister = flag.Bool("i", false, "not case sensitive.")

func main() {
	flag.Parse()

	fields := false
	if *flagField != 0 {
		fields = true
	}
	chars := false
	if *flagChar != 0 {
		chars = true
	}
	options := unique.Options{
		Count:     *flagCount,
		Duplicate: *flagDuplicate,
		Unique:    *flagUnique,
		Fields:    fields,
		NumFields: *flagField,
		Chars:     chars,
		NumChars:  *flagChar,
		Register:  *flagRegister,
	}

	inputFileName := ""
	outputFileName := ""
	args := flag.Args()
	if len(args) == 2 {
		inputFileName = args[0]
		outputFileName = args[1]
	} else if len(args) == 1 {
		inputFileName = args[0]
	}

	input, err := workWithFiles.Read(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	output, err := unique.FindUnique(input, options)
	if err != nil {
		log.Fatal(err)
	}

	err = workWithFiles.Write(output, outputFileName)
	if err != nil {
		log.Fatal(err)
	}
}
