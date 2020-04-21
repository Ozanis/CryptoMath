package hardware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	. "text/template"
)

const rootPath = "core/number/skel/"

type Skel struct {
	Name, UseTemplate, PackageName, OutFilePath, Type, Zero string
}

func ReadFile(filePath string) []byte {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
	}
	return file
}

func TemplateFromJson(jsonData []byte) Skel {
	var skel Skel
	err := json.Unmarshal(jsonData, &skel)
	if err != nil {
		log.Println(err)
	}
	return skel
}

func ParseTemplate(skel Skel) *Template {
	return Must(New(skel.PackageName).Parse(string(ReadFile(skel.PackageName))))
}

func CodeGen(template *Template) error {
	return template.Execute(os.Stdout, template)
}

func BuildKernel() {
	templateFile := rootPath + UseGpu()
	if _, err := os.Stat(templateFile); os.IsExist(err) {
		log.Println("Kernel has been built already")
		return
	}
	err := CodeGen(ParseTemplate(TemplateFromJson(ReadFile(templateFile))))
	if err != nil {
		panic("unable to build kernel")
	}
}
