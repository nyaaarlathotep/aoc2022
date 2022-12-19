package DataStruct

type Point3d struct {
	X int64
	Y int64
	Z int64
}

func (p *Point3d) Add(b Point3d) Point3d {
	return Point3d{
		X: p.X + b.X,
		Y: p.Y + b.Y,
		Z: p.Z + b.Z,
	}
}
