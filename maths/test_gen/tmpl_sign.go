package main

import (
	"github.com/MakeNowJust/heredoc/v2"
)

var (
	sign_uint_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var i {{ .numType }} = 0
			for ; i < 100; i++ {
				if i == 0 {
					if !assert.Equal(t, {{ .numType }}(0), maths.Sign(i), "i=%d", i) {
						t.FailNow()
					}
				} else {
					if !assert.Equal(t, {{ .numType }}(1), maths.Sign(i), "i=%d", i) {
						t.FailNow()
					}
				}
			}
			for i = math.Max{{ capitalize .numType }} - 100; i != 0; i++ {
				if !assert.Equal(t, {{ .numType }}(1), maths.Sign(i), "i=%d", i) {
					t.FailNow()
				}
			}
		})
	`)

	sign_int_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			var i {{ .numType }} = math.Min{{ capitalize .numType }}
			for ; i < math.Min{{ capitalize .numType }}+100; i++ {
				if !assert.Equal(t, {{ .numType }}(-1), maths.Sign(i), "i=%d", i) {
					t.FailNow()
				}
			}
			for i = -50; i < 50; i++ {
				if i == 0 {
					if !assert.Equal(t, {{ .numType }}(0), maths.Sign(i), "i=%d", i) {
						t.FailNow()
					}
				} else if i < 0 {
					if !assert.Equal(t, {{ .numType }}(-1), maths.Sign(i), "i=%d", i) {
						t.FailNow()
					}
				} else {
					if !assert.Equal(t, {{ .numType }}(1), maths.Sign(i), "i=%d", i) {
						t.FailNow()
					}
				}
			}
			for i = math.Max{{ capitalize .numType }} - 100; i > 0; i++ {
				if !assert.Equal(t, {{ .numType }}(1), maths.Sign(i), "i=%d", i) {
					t.FailNow()
				}
			}
		})
	`)

	sign_float_tmpl = heredoc.Doc(`
		t.Run("{{ .numType }}", func(t *testing.T) {
			for sign := {{ .numType }}(-1); sign <= 1; sign += 2 {
				for f, j := {{ .numType }}(0), 0; j < 100; j++ {
					if j == 0 {
						if !assert.Equal(t, {{ .numType }}(0), maths.Sign(f), "f=%f", f) {
							t.FailNow()
						}
					} else {
						if !assert.Equal(t, sign, maths.Sign(f), "f=%f", f) {
							t.FailNow()
						}
					}
					f = f*2 + math.SmallestNonzero{{ capitalize .numType }}*sign
				}
		
				for f, j := (math.Max{{ capitalize .numType }}-100)*sign, 0; j < 100; j++ {
					if !assert.Equal(t, sign, maths.Sign(f), "f=%f", f) {
						t.FailNow()
					}
					f += 1.0
				}
			}
			for f, j := {{ .numType }}(-50), 0; j < 100; j++ {
				if f == 0 {
					if !assert.Equal(t, {{ .numType }}(0), maths.Sign(f), "f=%f", f) {
						t.FailNow()
					}
				} else if f < 0 {
					if !assert.Equal(t, {{ .numType }}(-1), maths.Sign(f), "f=%f", f) {
						t.FailNow()
					}
				} else {
					if !assert.Equal(t, {{ .numType }}(1), maths.Sign(f), "f=%f", f) {
						t.FailNow()
					}
				}
				f += 1.0
			}
		})
	`)
)
