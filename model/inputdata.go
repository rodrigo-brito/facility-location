package model

type InputData struct {
	InstallationCost []float64
	Distance [][]float64
	Flow [][]float64
}

func InputDataFromFile(file string) (*InputData, error) {
	return &InputData{

	}, nil
}