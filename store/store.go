package store

type Store struct {
	Ip          string
	Port        uint32
	DataDir     string
	TotalVolume uint64
	TotalNeedle uint64
	FreeSpace   uint64
}

func (s Store) Run() {

}
