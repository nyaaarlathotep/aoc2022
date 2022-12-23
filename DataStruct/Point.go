package DataStruct

type Point struct {
	X int64
	Y int64
}

func (p *Point) Add(b Point) Point {
	return Point{
		X: p.X + b.X,
		Y: p.Y + b.Y,
	}
}

func (p *Point) Clone() Point {
	return Point{
		X: p.X,
		Y: p.Y,
	}
}
