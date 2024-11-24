package data

type AntFarm struct {
	Ants     int      //number of ants in the file
	Start    int      //the start room
	End      int      // the end room
	Turns    int      // number of turns
	Filedata []string //the whole filedata
	Rooms    []string
	Tunnels  []string
	Paths    []string
}

func NewFarm() *AntFarm {
	return &AntFarm{}
}

type Ant struct {
	Index       int
	Target      int //target path
	TargetIndex int // the room of the target path
	Iteration   int // the specific turn
}

func NewAnt() *Ant {
	return &Ant{}
}
