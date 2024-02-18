package maths

import "fmt"

type AggFunc[TNum Number] func(ts ...TNum) TNum

type Aggregator string

const (
	AggMax     Aggregator = "Max"
	AggMin     Aggregator = "Min"
	AggSum     Aggregator = "Sum"
	AggAvg     Aggregator = "Avg"
	AggProd    Aggregator = "Prod"
	AggGeoMean Aggregator = "Geom"
	AggHarMean Aggregator = "Harm"
	AggXenoSum Aggregator = "Xeno"
)

func GetAgg[TNum Number](agg Aggregator) AggFunc[TNum] {
	switch agg {
	case AggMax:
		return Max[TNum]
	case AggMin:
		return Min[TNum]
	case AggSum:
		return Sum[TNum]
	case AggAvg:
		return func(ts ...TNum) TNum {
			if len(ts) == 0 {
				return 0
			}
			return Sum[TNum](ts...) / TNum(len(ts))
		}
	case AggProd:
		return Prod[TNum]
	case AggGeoMean:
		return GeometricMean[TNum]
	case AggHarMean:
		return HarmonicMean[TNum]
	case AggXenoSum:
		return XenoSum[TNum]
	default:
		panic(fmt.Errorf("unknown aggregator: %s", agg))
	}

}
