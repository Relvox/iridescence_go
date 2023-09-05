package templates

import (
	"bytes"
	"html/template"
)

func ApplyTemplateFuncGen(t *template.Template) func(string, interface{}) (template.HTML, error) {
	return func(templateName string, data interface{}) (template.HTML, error) {
		var buf bytes.Buffer
		err := t.ExecuteTemplate(&buf, templateName, data)
		if err != nil {
			return "", err
		}
		return template.HTML(buf.String()), nil
	}
}
