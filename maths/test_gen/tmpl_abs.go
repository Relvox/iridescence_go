package main

import "github.com/MakeNowJust/heredoc/v2"

var (
	abs_uint_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var i {{ .numType }} = 0
			for ; i < 100; i++ {
				if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
					t.FailNow()
				}
			}
			for i = math.Max{{ capitalize .numType }} - 100; i != 0; i++ {
				if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
					t.FailNow()
				}
			}
		})
	`)

	abs_int_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var i {{ .numType }} = math.Min{{ capitalize .numType }}
			for ; i < math.Min{{ capitalize .numType }}+100; i++ {
				if !assert.Equal(t, {{ .numType }}(-i), maths.Abs(i), "i=%d", i) {
					t.FailNow()
				}
			}
			for i = -50; i < 50; i++ {
				expected := {{ .numType }}(i)
				if i < 0 {
					expected = -expected
				}
				if !assert.Equal(t, expected, maths.Abs(i), "i=%d", i) {
					t.FailNow()
				}
			}
			for i = math.Max{{ capitalize .numType }} - 100; i > 0; i++ {
				if !assert.Equal(t, i, maths.Abs(i), "i=%d", i) {
					t.FailNow()
				}
			}
		})
	`)

	abs_float_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			for sign := {{ .numType }}(-1); sign <= 1; sign += 2 {
				for f, j := {{ .numType }}(0), 0; j < 100; j++ {
					expected := f
					if f < 0 {
						expected = -f
					}
					if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
						t.FailNow()
					}
					f = f*2 + math.SmallestNonzero{{ capitalize .numType }}*sign
				}

				for f, j := (math.Max{{ capitalize .numType }}-100)*sign, 0; j < 100; j++ {
					expected := f
					if f < 0 {
						expected = -f
					}
					if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
						t.FailNow()
					}
					f += 1.0
				}
			}
			for f, j := {{ .numType }}(-50), 0; j < 100; j++ {
				expected := f
				if f < 0 {
					expected = -f
				}
				if !assert.Equal(t, expected, maths.Abs(f), "f=%f", f) {
					t.FailNow()
				}
				f += 1.0
			}
		})
	`)
)
