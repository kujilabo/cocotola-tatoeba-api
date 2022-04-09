package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/pkg/profile"

	"github.com/kujilabo/cocotola-tatoeba-api/pkg/gateway"
)

func run(fileName string) (map[int]bool, error) {
	filePath := "../cocotola-data/datasource/tatoeba/" + fileName

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	iterator := gateway.NewTatoebaSentenceAddParameterReader(file)

	sentenceNumbers := make(map[int]bool)
	for {
		param, err := iterator.Next(context.Background())
		if errors.Is(err, io.EOF) {
			break
		}
		if param == nil {
			continue
		}
		sentenceNumbers[param.GetSentenceNumber()] = true
	}

	return sentenceNumbers, nil
}

func writeLinks(eng map[int]bool, jpn map[int]bool, fileName1, fileName2 string) error {

	filePath := "../cocotola-data/datasource/tatoeba/" + fileName1

	file1, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file1.Close()

	iterator := gateway.NewTatoebaLinkAddParameterReader(file1)

	filePath2 := "../cocotola-data/datasource/tatoeba/" + fileName2

	file2, err := os.Create(filePath2)
	if err != nil {
		return err
	}
	defer file2.Close()

	writer := csv.NewWriter(file2)
	writer.Comma = '\t'
	defer writer.Flush()

	for {
		param, err := iterator.Next(context.Background())
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		if param == nil {
			continue
		}
		if _, engOk := eng[param.GetFrom()]; engOk {
			if _, jpnOk := jpn[param.GetTo()]; jpnOk {
				record := []string{strconv.Itoa(param.GetFrom()), strconv.Itoa(param.GetTo())}
				if err := writer.Write(record); err != nil {
					return err
				}
			}
		}
		if _, jpnOk := jpn[param.GetFrom()]; jpnOk {
			if _, engOk := eng[param.GetTo()]; engOk {
				record := []string{strconv.Itoa(param.GetFrom()), strconv.Itoa(param.GetTo())}
				if err := writer.Write(record); err != nil {
					return err
				}
			}
		}
		param = nil
	}
	return nil
}

func main() {
	defer profile.Start(profile.MemProfile).Stop()

	eng, err := run("eng_sentences_detailed.tsv")
	if err != nil {
		panic(err)
	}
	fmt.Println("eng read")
	fmt.Printf("%d\n", len(eng))

	jpn, err := run("jpn_sentences_detailed.tsv")
	if err != nil {
		panic(err)
	}
	fmt.Println("jpn read")
	fmt.Printf("%d\n", len(jpn))

	if err := writeLinks(eng, jpn, "links.csv", "links2.csv"); err != nil {
		panic(err)
	}

}
