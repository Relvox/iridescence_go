package test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/relvox/iridescence_go/maths"
// 	"github.com/relvox/iridescence_go/maths/integer/intervals"
// )

// /* Cases

// Values:
// Singleton Cases
// ---|-----|-----|-----|-----|-----|-----|-----|---
// ...0.....1.....2.....3.....4.....5.....6.....7...
// ...a.....b.....c.....d.....e.....f.....g.....h...
// ---|-----|-----|-----|-----|-----|-----|-----|---
// ...Reference...X=================Y...Interval....[ 4 Reference configs ]
// ---|-----|-----|-----|-----|-----|-----|-----|---
// ...|.....S0....S1....S2....|.....S3....S4........[ 5 Singleton cases ]
// ---|-----|-----|-----|-----|-----|-----|-----|---
// ...a.....b.....c.....d.....e.....f.....g.....h...
// Interval Cases
// ...a.....b.....c.....d.....e.....f.....g.....h...
// ...Reference...X=================Y...Interval....
// ---|-----|-----|-----|-----|-----|-----|-----|---
// ...A_====A_....|.....|.....|.....|.....|.....|...[ 1 Case ]
// ...|.....B_====BX....|.....|.....|.....|.....|...[ 2 Cases ]
// ...|.....C_==========CX....|.....|.....|.....|...[ 2 Cases ]
// ...|.....D_======================DX....|.....|...[ 2 Cases ]
// ...|.....E_============================E_....|...[ 1 Case ]
// ...|.....|.....|.....|.....|.....|.....|.....|...
// ...a.....b.....c.....d.....e.....f.....g.....h...
// ...Reference...X=================Y...Interval....
// ---|-----|-----|-----|-----|-----|-----|-----|---
// ...|.....|.....FX====FY....|.....|.....|.....|...[ 4 Cases ]
// ...|.....|.....GX================GY....|.....|...[ 4 Cases ]
// ...|.....|.....HX======================H_....|...[ 2 Cases ]
// ...|.....|.....|.....|.....|.....|.....|.....|...
// ...a.....b.....c.....d.....e.....f.....g.....h...
// ...Reference...X=================Y...Interval....
// ---|-----|-----|-----|-----|-----|-----|-----|---
// ...|.....|.....|.....IX====IY....|.....|.....|...[ 4 Cases ]
// ...|.....|.....|.....JX==========JY....|.....|...[ 4 Cases ]
// ...|.....|.....|.....KX================K_....|...[ 2 Cases ]
// ...|.....|.....|.....|.....|.....LX====L_....|...[ 2 Cases ]
// ...|.....|.....|.....|.....|.....|.....M_====M_..[ 1 Case ]
// ...|.....|.....|.....|.....|.....|.....|.....|...
// */

// type intervalTestCase[T Number] struct {
// 	reference       intervals.Interval[T]
// 	target          intervals.Interval[T]
// 	expectedOverlap intervals.Interval[T]
// }

// // func generateSingletonTestCases[T Number](reference intervals.Interval[T], values [5]T) []intervalTestCase[T] {
// // 	var res []intervalTestCase[T]
// // }

// func generatePairTestCases[T Number](values [8]T) []intervalTestCase[T] {
// 	for i := 1; i < len(values); i++ {
// 		if values[i-1] >= values[i] {
// 			panic(fmt.Errorf("values must be monotonic and rising: %v", values))
// 		}
// 	}
// 	// references := [4]intervals.Interval[T]{
// 	// 	intervals.NewInterval(values[2], true, values[5], true),
// 	// 	intervals.NewInterval(values[2], true, values[5], false),
// 	// 	intervals.NewInterval(values[2], false, values[5], true),
// 	// 	intervals.NewInterval(values[2], false, values[5], false),
// 	// }
// 	// allCases := []intervalTestCase[T]{}
// 	// for _, ref := range references {
// 	// allCases = append(allCases, generateSingletonTestCases(ref, [5]T{values[1], values[2], values[3], values[5], values[6]})...)
// 	// }
// 	return nil
// }
// func Test(t *testing.T) {

// }
