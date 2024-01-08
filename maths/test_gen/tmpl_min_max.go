package main

import (
	"github.com/MakeNowJust/heredoc/v2"
)

var (
	min_max_uint_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 10, 20, 30
			if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
			if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	min_max_int_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = -10, 0, 10
			if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
			if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	min_max_float_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = -1.5, 0.0, 1.5
			if !assert.Equal(t, a, maths.Min(a, b, c), "Min of %f, %f, %f", a, b, c) {
				t.FailNow()
			}
			if !assert.Equal(t, c, maths.Max(a, b, c), "Max of %f, %f, %f", a, b, c) {
				t.FailNow()
			}
			// Testing with extreme values
			var max, min {{ .numType }} = math.Max{{ capitalize .numType }}, -math.Max{{ capitalize .numType }}
			if !assert.Equal(t, min, maths.Min(min, 0, max), "Min of %f, %f, %f", min, 0.0, max) {
				t.FailNow()
			}
			if !assert.Equal(t, max, maths.Max(min, 0, max), "Max of %f, %f, %f", min, 0.0, max) {
				t.FailNow()
			}
			
			max, min = math.SmallestNonzero{{ capitalize .numType }}, -math.SmallestNonzero{{ capitalize .numType }}
			if !assert.Equal(t, min, maths.Min(min, 0, max), "Min of %f, %f, %f", min, 0.0, max) {
				t.FailNow()
			}
			if !assert.Equal(t, max, maths.Max(min, 0, max), "Max of %f, %f, %f", min, 0.0, max) {
				t.FailNow()
			}
		})
	`)
)
