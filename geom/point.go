package geom

type Point2 [2]int
type Point3 [3]int
type Point4 [4]int

func (v Point2) X() int         { return v[0] }
func (v Point2) Y() int         { return v[1] }
func (v Point2) XY() (int, int) { return v[0], v[1] }
func (v Point2) Point2() Point2 { return v }
func (v Point2) Point3() Point3 { return Point3{v[0], v[1]} }
func (v Point2) Point4() Point4 { return Point4{v[0], v[1]} }
func (v Point2) Offset(o Point2) Point2 {
	var result Point2
	for i, x := range v {
		result[i] = x + o[i]
	}
	return result
}

func (v Point3) X() int         { return v[0] }
func (v Point3) Y() int         { return v[1] }
func (v Point3) Z() int         { return v[2] }
func (v Point3) Point2() Point2 { return Point2{v[0], v[1]} }
func (v Point3) Point3() Point3 { return v }
func (v Point3) Point4() Point4 { return Point4{v[0], v[1], v[2]} }
func (v Point3) Offset(o Point3) Point3 {
	var result Point3
	for i, x := range v {
		result[i] = x + o[i]
	}
	return result
}

func (v Point4) X() int         { return v[0] }
func (v Point4) Y() int         { return v[1] }
func (v Point4) Z() int         { return v[2] }
func (v Point4) W() int         { return v[3] }
func (v Point4) Point2() Point2 { return Point2{v[0], v[1]} }
func (v Point4) Point3() Point3 { return Point3{v[0], v[1], v[2]} }
func (v Point4) Point4() Point4 { return v }
func (v Point4) Offset(o Point4) Point4 {
	var result Point4
	for i, x := range v {
		result[i] = x + o[i]
	}
	return result
}

var (
	N = Point2{0, 1}
	S = Point2{0, -1}
	E = Point2{1, 0}
	W = Point2{-1, 0}
)
