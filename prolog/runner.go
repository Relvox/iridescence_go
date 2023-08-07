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

func NewRunner(db string) (*Runner, error) {
	result := &Runner{
		Interpreter: prolog.New(nil, nil),
		DbText:      db,
	}

	err := result.Interpreter.Exec(db)
	if err != nil {
		return nil, fmt.Errorf("exec initial db: %w", err)
	}
	return result, nil
}

func (r *Runner) NewRunnerWith(prolog string) (*Runner, error) {
	return NewRunner(fmt.Sprintf("%s\n\n%s", r.DbText, prolog))
}

func (r *Runner) RebuildRunnerWith(prolog string) (*Runner, error) {
	r.DbText = fmt.Sprintf("%s\n\n%s", r.DbText, prolog)
	err := r.Interpreter.Exec(r.DbText)
	if err != nil {
		return nil, fmt.Errorf("exec rebuild with: %w", err)
	}
	return r, nil
}

func (r *Runner) AppendReplace(prolog string) (*Runner, error) {
	r.DbText = fmt.Sprintf("%s\n\n%s", r.DbText, prolog)
	err := r.Interpreter.Exec(prolog)
	if err != nil {
		return nil, fmt.Errorf("exec append replace: %w", err)
	}
	return r, nil
}

func (r *Runner) Query(query string) ([]map[string]any, error) {
	var result []map[string]any
	sols, err := r.Interpreter.Query(query)
	if err != nil {
		return result, fmt.Errorf("query %s: %w", query, err)
	}
	defer sols.Close()

	for sols.Next() {
		s := map[string]any{}
		if err := sols.Scan(&s); err != nil {
			return result, fmt.Errorf("scan result: %w", err)
		}
		result = append(result, s)
	}

	if err := sols.Err(); err != nil {
		return result, err
	}
	return result, nil
}

func RunTypedQuery[TResult any](r *Runner, query string) ([]TResult, error) {
	rawResults, err := r.Query(query)
	if err != nil {
		var z TResult
		return nil, fmt.Errorf("typed [%T] query %s: %w", z, query, err)
	}
	var results []TResult
	for _, rawResult := range rawResults {
		result, err := utils.MapToStruct[TResult](rawResult)
		if err != nil {
			return results, fmt.Errorf("formatting result: %w", err)
		}
		results = append(results, result)
	}
	return results, nil
}
