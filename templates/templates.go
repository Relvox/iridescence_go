package templates

import (
	"bytes"
	"html/template"
)

func ApplyTemplateFuncGen(t *template.Template) func(string, any) (string, error) {
	return func(templateName string, data any) (string, error) {
		var buf bytes.Buffer
		err := t.ExecuteTemplate(&buf, templateName, data)
		if err != nil {
			return "", err
		}
		return buf.String(), nil
	}
}
