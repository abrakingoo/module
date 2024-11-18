package data

type AntFarm struct {
    Ants int
    Start int
    End int
    Turns int
    Farm []string
    Rooms []string
    Tunnels []string
    Paths []string
}

func NewFarm() *AntFarm {
    return &AntFarm{}
}

type Ant struct {
	Index int
	Target int
	TargetIndex int
	Iteration int
}

func NewAnt() *Ant {
	return &Ant{}
}

