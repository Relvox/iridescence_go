package main

import "github.com/MakeNowJust/heredoc/v2"

var (
	sum_uint_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 10, 20, 30
			if !assert.Equal(t, {{ .numType }}(60), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	sum_int_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = -10, 0, 10
			if !assert.Equal(t, {{ .numType }}(0), maths.Sum(a, b, c), "Sum of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	sum_float_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = -1.5, 0.0, 1.5
			if !assert.Equal(t, {{ .numType }}(0.0), maths.Sum(a, b, c), "Sum of %f, %f, %f", a, b, c) {
				t.FailNow()
			}
		})
	`)

	geom_mean_uint_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 1, 2, 4
			if !assert.Equal(t, {{ .numType }}(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	geom_mean_int_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 1, 2, 4
			if !assert.Equal(t, {{ .numType }}(2), maths.GeometricMean(a, b, c), "Geometric mean of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	geom_mean_float_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 1.0, 2.0, 4.0
			if !assert.Equal(t, {{ .numType }}(2.0), maths.GeometricMean(a, b, c), "Geometric mean of %f, %f, %f", a, b, c) {
				t.FailNow()
			}
		})
	`)

	xeno_sum_uint_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 4, 4, 4
			if !assert.Equal(t, {{ .numType }}(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	xeno_sum_int_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 4, 4, 4
			if !assert.Equal(t, {{ .numType }}(7), maths.XenoSum(a, b, c), "Xeno sum of %d, %d, %d", a, b, c) {
				t.FailNow()
			}
		})
	`)

	xeno_sum_float_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var a, b, c {{ .numType }} = 4.0, 4.0, 4.0
			if !assert.Equal(t, {{ .numType }}(7.0), maths.XenoSum(a, b, c), "Xeno sum of %f, %f, %f", a, b, c) {
				t.FailNow()
			}
		})
	`)
)
