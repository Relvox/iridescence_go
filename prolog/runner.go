package prolog

import (
	"fmt"

	"github.com/ichiban/prolog"

	"github.com/relvox/iridescence_go/utils"
)

type Runner struct {
	Interpreter *prolog.Interpreter
	DbText      string
}

func NewRunner(db string) *Runner {
	result := &Runner{
		Interpreter: prolog.New(nil, nil),
		DbText:      db,
	}

	if err := result.Interpreter.Exec(db); err != nil {
		panic(err)
	}
	return result
}

func (r *Runner) NewRunnerWith(prolog string) *Runner {
	return NewRunner(fmt.Sprintf("%s\n\n%s", r.DbText, prolog))
}

func (r *Runner) RebuildRunnerWith(prolog string) *Runner {
	r.DbText = fmt.Sprintf("%s\n\n%s", r.DbText, prolog)
	if err := r.Interpreter.Exec(r.DbText); err != nil {
		panic(err)
	}
	return r
}

func (r *Runner) AppendReplace(prolog string) *Runner {
	r.DbText = fmt.Sprintf("%s\n\n%s", r.DbText, prolog)
	if err := r.Interpreter.Exec(prolog); err != nil {
		panic(err)
	}
	return r
}

func (r *Runner) Query(query string) []map[string]any {
	var result []map[string]any
	sols, err := r.Interpreter.Query(query)
	if err != nil {
		panic(err)
	}
	defer sols.Close()

	for sols.Next() {
		s := map[string]any{}
		if err := sols.Scan(&s); err != nil {
			panic(err)
		}
		result = append(result, s)
	}

	if err := sols.Err(); err != nil {
		panic(err)
	}
	return result
}

func RunTypedQuery[TResult any](r *Runner, query string) []TResult {
	rawResults := r.Query(query)
	var results []TResult
	for _, rawResult := range rawResults {
		result, err := utils.MapToStruct[TResult](rawResult)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
	}
	return results
}
