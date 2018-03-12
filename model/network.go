package model

type Network struct {
	Size int

	Hubs []int32
	HubsBin map[int32]bool
	Allocation [][]int

	Data *InputData
}