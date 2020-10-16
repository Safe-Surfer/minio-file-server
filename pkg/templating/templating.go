package templating

import (
	"bytes"
	"github.com/minio/minio-go/v7"
	"html/template"
	"log"
)

type TemplatePath string

const (
	TemplateListing TemplatePath = "templates/listing.html"
	TemplateIndex   TemplatePath = "templates/index.html"
)

type TemplateListingObject struct {
	SiteTitle string
	Path      string
	Items     []minio.ObjectInfo
}

type TemplateIndexObject struct {
	SiteTitle string
	Path      string
}

func Template(templatePath TemplatePath, data interface{}) (error, string) {
	template, err := template.ParseFiles(string(templatePath))
	if err != nil {
		return err, ""
	}
	templateBuffer := new(bytes.Buffer)
	err = template.Execute(templateBuffer, data)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	templatedHTML := templateBuffer.String()
	return err, templatedHTML
}
