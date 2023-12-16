package geom

type Vector2 [2]float64
type Vector3 [3]float64
type Vector4 [4]float64

func (v Vector2) X() float64       { return v[0] }
func (v Vector2) Y() float64       { return v[1] }
func (v Vector2) Vector2() Vector2 { return v }
func (v Vector2) Vector3() Vector3 { return Vector3{v[0], v[1]} }
func (v Vector2) Vector4() Vector4 { return Vector4{v[0], v[1]} }
func (v Vector2) Offset(o Vector2) Vector2 {
	var result Vector2
	for i, x := range v {
		result[i] = x + o[i]
	}
	return result
}

func (v Vector3) X() float64       { return v[0] }
func (v Vector3) Y() float64       { return v[1] }
func (v Vector3) Z() float64       { return v[2] }
func (v Vector3) Vector2() Vector2 { return Vector2{v[0], v[1]} }
func (v Vector3) Vector3() Vector3 { return v }
func (v Vector3) Vector4() Vector4 { return Vector4{v[0], v[1], v[2]} }
func (v Vector3) Offset(o Vector3) Vector3 {
	var result Vector3
	for i, x := range v {
		result[i] = x + o[i]
	}
	return result
}

func (v Vector4) X() float64       { return v[0] }
func (v Vector4) Y() float64       { return v[1] }
func (v Vector4) Z() float64       { return v[2] }
func (v Vector4) W() float64       { return v[3] }
func (v Vector4) Vector2() Vector2 { return Vector2{v[0], v[1]} }
func (v Vector4) Vector3() Vector3 { return Vector3{v[0], v[1], v[2]} }
func (v Vector4) Vector4() Vector4 { return v }
func (v Vector4) Offset(o Vector4) Vector4 {
	var result Vector4
	for i, x := range v {
		result[i] = x + o[i]
	}
	return result
}
