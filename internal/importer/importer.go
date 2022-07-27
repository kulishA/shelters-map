package importer

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"os"
	"sheltes-map/internal/domain"
	"sheltes-map/internal/repository"
)

type IImporter interface {
}

type Importer struct {
	repo repository.IShelterRepository
}

func NewImporter(repo repository.IShelterRepository) *Importer {
	return &Importer{repo: repo}
}

func (i *Importer) Import() error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		filePath := fmt.Sprintf("%s%s", dirPath, file.Name())
		if err := i.processFile(filePath); err != nil {
			return err
		}
	}
	return nil
}

func (i *Importer) processFile(filePath string) error {
	shelterFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer shelterFile.Close()

	var shelters []*domain.Shelter

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = ';'
		return r
	})

	if err := gocsv.UnmarshalFile(shelterFile, &shelters); err != nil {
		return err
	}

	for _, shelter := range shelters {
		if _, err := i.repo.Save(shelter); err != nil {
			return err
		}
	}

	return nil
}
