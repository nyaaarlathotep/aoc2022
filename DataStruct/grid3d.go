package DataStruct

type Grid3d[T any] struct {
	xLen      int64
	yLen      int64
	zLen      int64
	state     [][][]T
	movements []Point
}

func NewGrid3d[T any](xLen, yLen, zLen int64) *Grid3d[T] {
	state := make([][][]T, zLen)
	for z := int64(0); z < zLen; z++ {
		state[z] = make([][]T, yLen)
		for y := int64(0); y < yLen; y++ {
			state[z][y] = make([]T, xLen)
		}
	}

	return &Grid3d[T]{
		xLen:  xLen,
		yLen:  yLen,
		zLen:  zLen,
		state: state,
	}
}

func (g *Grid3d[T]) XLen() int64 {
	return g.xLen
}

func (g *Grid3d[T]) YLen() int64 {
	return g.yLen
}

func (g *Grid3d[T]) ZLen() int64 {
	return g.zLen
}

func (g *Grid3d[T]) isValid(x, y, z int64) bool {
	switch {
	case x < 0, x >= g.xLen, y < 0, y >= g.yLen, z < 0, z >= g.zLen:
		return false
	default:
		return true
	}
}
func (g *Grid3d[T]) SetState(x, y, z int64, state T) {
	if g.isValid(x, y, z) {
		g.state[z][y][x] = state
	}
}

func (g *Grid3d[T]) GetState(z, y, x int64) T {
	return g.state[z][y][x]
}

//
//func (g *Grid3d[T]) StateMap() map[Point]T {
//	ret := make(map[Point]T)
//
//	for y, l := range g.state {
//		for x, s := range l {
//			ret[Point{int64(x), int64(y)}] = s
//		}
//	}
//
//	return ret
//}

func (g *Grid3d[T]) StateMapWhere(f func(T) bool) map[Point3d]T {
	ret := make(map[Point3d]T)
	for z, p := range g.state {
		for y, l := range p {
			for x, s := range l {
				if f(s) {
					ret[Point3d{int64(x), int64(y), int64(z)}] = s
				}
			}
		}
	}

	return ret
}
