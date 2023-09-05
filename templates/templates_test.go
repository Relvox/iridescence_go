package templates

import (
	"bytes"
	"testing"
	"text/template"
)

func TestApplyTemplateFunc(t *testing.T) {
	const fragmentTemplateStr = `
	{{- define "fragment1" }}Hello, {{ .Name }}!{{ end }}
	{{- define "fragment2" }}Goodbye, {{ .Name }}!{{ end }}
	`

	tmpl := template.New("main")
	tmpl = tmpl.Funcs(template.FuncMap{
		"applyTemplate": ApplyTemplateFuncGen(tmpl),
	})
	tmpl, err := tmpl.Parse(fragmentTemplateStr + `{{- applyTemplate .TemplateName .Data }}`)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	tests := []struct {
		name         string
		templateName string
		data         interface{}
		expected     string
		expectError  bool
	}{
		{
			name:         "Valid Template 1",
			templateName: "fragment1",
			data:         map[string]string{"Name": "John"},
			expected:     "Hello, John!",
			expectError:  false,
		},
		{
			name:         "Valid Template 2",
			templateName: "fragment2",
			data:         map[string]string{"Name": "John"},
			expected:     "Goodbye, John!",
			expectError:  false,
		},
		{
			name:         "Invalid Template Name",
			templateName: "fragment3",
			data:         map[string]string{"Name": "John"},
			expected:     "",
			expectError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tmpl.Execute(&buf, map[string]interface{}{
				"TemplateName": tt.templateName,
				"Data":         tt.data,
			})
			if (err != nil) != tt.expectError {
				t.Fatalf("Expected error: %v, got: %v", tt.expectError, err)
			}

			actual := buf.String()
			if actual != tt.expected {
				t.Fatalf("Unexpected output: got %v, want %v", actual, tt.expected)
			}
		})
	}
}
