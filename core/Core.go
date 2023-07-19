package core

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	dir         = "result"
	ext         = ".html"
	defFileName = "template-%v"
	filePerm    = 0755
)

type Elements []map[string]interface{}

func Generate(templatePath, dataPath, fileIdentifier string, size int8) {
	log.Print("Html generation init...")
	elements := chargeData(dataPath)
	template := chargeTemplate(templatePath)
	generateFile(int(size), fileIdentifier, elements, template)
	log.Print("Html generation done-")
}

func generateFile(size int, fileIdentifier string, elements Elements, template *template.Template) {
	setupDirectory()
	if size == 0 {
		size = len(elements)
	}
	for i := 0; i < size; i++ {
		fileName := fmt.Sprintf(defFileName, i)
		if !strings.Contains(fileIdentifier, "template") {
			fileName = fmt.Sprintf("%v", elements[i][fileIdentifier])
		}
		builderData := &strings.Builder{}
		template.Execute(builderData, &elements[i])
		createAndWriteFile(fileName, builderData.String())
	}
}

func createAndWriteFile(fileName, data string) {
	fileNameFinal := fileName + ext
	err := os.WriteFile(currentDir()+dir+"/"+fileNameFinal, []byte(data), os.ModePerm)
	ifErrorStop(err)
	log.Print("File generated:" + fileNameFinal)
}

func setupDirectory() {
	workDir := filepath.Join(currentDir(), dir)
	if err := os.RemoveAll(workDir); err != nil {
		ifErrorStop(err)
	}
	if err := os.MkdirAll(workDir, os.ModePerm); err != nil {
		ifErrorStop(err)
	}
	log.Print("Using directory:" + workDir)
}

func chargeTemplate(templatePath string) *template.Template {
	temp, err := template.ParseFiles(currentDir() + templatePath)
	ifErrorStop(err)
	return temp
}

func currentDir() string {
	curDir, err := os.Getwd()
	ifErrorStop(err)
	return curDir + "/"
}

func chargeData(dataPath string) Elements {
	dataPath = fmt.Sprintf("%s%s", currentDir(), dataPath)
	switch {
	case strings.Contains(dataPath, ".json"):
		return dataFromJson(dataPath)
	case strings.Contains(dataPath, ".csv"):
		return dataFromCsv(dataPath)
	default:
		panic("Error not supported file")
	}
}

func dataFromCsv(path string) Elements {
	fileContent, err := os.Open(path)
	ifErrorStop(err)
	data, err := csv.NewReader(fileContent).ReadAll()
	ifErrorStop(err)

	var headers []string
	var elements Elements
	for i, col := range data {
		if i == 0 {
			for _, cell := range col {
				headers = append(headers, cell)
			}
			continue
		}
		line := map[string]interface{}{}
		for j, cell := range col {
			line[headers[j]] = cell
		}
		elements = append(elements, line)
	}
	return elements
}

func dataFromJson(path string) Elements {
	byteContent, err := os.ReadFile(path)
	ifErrorStop(err)
	elements := Elements{}
	ifErrorStop(json.Unmarshal(byteContent, &elements))
	return elements
}

func ifErrorStop(err error) {
	if err != nil {
		log.Panic("ERROR:" + err.Error())
	}
}
