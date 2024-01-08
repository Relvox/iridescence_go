package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/MakeNowJust/heredoc/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	header = heredoc.Doc(`
		package main_test
		
		import (
			"testing"
			"math"
			
			"github.com/stretchr/testify/assert"
			
			"github.com/relvox/iridescence_go/maths"
		)
	`)
)

var (
	uintTypes    = []string{"uint", "uint8", "uint16", "uint32", "uint64"}
	intTypes     = []string{"int", "int8", "int16", "int32", "int64"}
	floatTypes   = []string{"float32", "float64"}
	typeSubtypes = [3][]string{uintTypes, intTypes, floatTypes}
	typeNames    = [3]string{"Uint", "Int", "Float"}

	templates = map[string][3]string{
		"Sign": {
			sign_uint_tmpl,
			sign_int_tmpl,
			sign_float_tmpl,
		},
		"Abs": {
			abs_uint_tmpl,
			abs_int_tmpl,
			abs_float_tmpl,
		},
		"MinMax": {
			min_max_uint_tmpl,
			min_max_int_tmpl,
			min_max_float_tmpl,
		},
		"Sum": {
			sum_uint_tmpl,
			sum_int_tmpl,
			sum_float_tmpl,
		},
		"GeometricMean": {
			geom_mean_uint_tmpl,
			geom_mean_int_tmpl,
			geom_mean_float_tmpl,
		},
		"XenoSum": {
			xeno_sum_uint_tmpl,
			xeno_sum_int_tmpl,
			xeno_sum_float_tmpl,
		},
	}

	funcs = []string{"Sign", "Abs", "MinMax", "Sum", "GeometricMean", "XenoSum"}
)

var compiledTemplates = map[string][3]*template.Template{}

func init() {
	for fun, tmplTexts := range templates {
		compiledTemplates[fun] = [3]*template.Template{{}, {}, {}}
		for tmplIndex, tmplTxt := range tmplTexts {
			template, err := template.New(fmt.Sprintf("%s_%s", fun, typeNames[tmplIndex])).
				Funcs(template.FuncMap{
					"capitalize": cases.Title(language.AmericanEnglish).String,
				}).Parse(tmplTxt)
			if err != nil {
				panic(err)
			}
			*compiledTemplates[fun][tmplIndex] = *template
		}
	}
}

func main() {
	for _, fun := range funcs {
		file, err := os.Create(fmt.Sprintf("./zz_maths_%s_test.go", strings.ToLower(fun)))
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fmt.Fprintln(file, header)
		fmt.Fprintf(file, "func Test_%s(t *testing.T) {\n", fun)

		for typeIndex, typeTemplate := range compiledTemplates[fun] {
			subTypes := typeSubtypes[typeIndex]
			for _, numType := range subTypes {
				data := map[string]string{
					"numType": numType,
				}
				var buff bytes.Buffer
				if err := typeTemplate.Execute(&buff, data); err != nil {
					panic(err)
				}

				fmt.Fprintln(file, buff.String())
			}
			fmt.Fprintln(file, "")
		}

		fmt.Fprintln(file, "}")

	}
}
