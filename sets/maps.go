package sets

func KeySet[TK comparable, TV any, TM ~map[TK]TV](src TM) Set[TK] {
	res := NewSet[TK]()
	for k := range src {
		res.Add(k)
	}
	return res
}
