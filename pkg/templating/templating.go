package templating

import (
	"bytes"
	"github.com/minio/minio-go/v7"
	"html/template"
	"log"
)

// TemplatePath ...
// valid paths for templates
type TemplatePath string

// template types
const (
	TemplateListing TemplatePath = "templates/listing.html"
	TemplateIndex   TemplatePath = "templates/index.html"
)

// TemplateListingObject ...
// fields for the list template
type TemplateListingObject struct {
	SiteTitle string
	Path      string
	Items     []minio.ObjectInfo
}

// TemplateIndexObject ...
// fields for the index template
type TemplateIndexObject struct {
	SiteTitle string
	Path      string
}

// Template ...
// given a template path and data, return the rendered template
func Template(templatePath TemplatePath, data interface{}) (string, error) {
	template, err := template.ParseFiles(string(templatePath))
	if err != nil {
		return "", err
	}
	templateBuffer := new(bytes.Buffer)
	err = template.Execute(templateBuffer, data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	templatedHTML := templateBuffer.String()
	return templatedHTML, err
}
