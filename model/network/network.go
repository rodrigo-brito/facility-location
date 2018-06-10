package network

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rodrigo-brito/hub-spoke-go/util/log"
)

// InputData store all network data
type Data struct {
	Size             int
	ScaleFactor      float64
	InstallationCost []float64
	Distance         [][]float64
	Flow             [][]float64
	FlowOrigin       []float64
	FlowDestiny      []float64
}

// parseLine validate a line of values and parse to a float array
func parseLine(line string, divisionFactor float64) (bool, []float64, error) {
	values := strings.Split(strings.TrimSpace(line), " ")
	numbers := make([]float64, 0)

	for _, value := range values {
		if len(value) == 0 {
			continue
		}

		number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			log.Error(err)
			return false, nil, err
		}

		numbers = append(numbers, number/divisionFactor)
	}

	if len(numbers) > 0 {
		return true, numbers, nil
	}
	return false, numbers, nil
}

// nextLine validate and return the next valid line
func nextLine(scanner *bufio.Scanner, divisionFactor float64) (bool, []float64, error) {
	if ok := scanner.Scan(); !ok {
		return false, nil, fmt.Errorf("unexpected end of file")
	}

	line := scanner.Text()
	if ok, values, err := parseLine(line, divisionFactor); ok {
		return true, values, nil
	} else if err != nil {
		return false, nil, err
	}

	// return the next valid line
	return nextLine(scanner, divisionFactor)
}

// FromFile read a input file and generate the network data
func FromFile(fileName string) (*Data, error) {
	data := new(Data)

	file, err := os.Open(fileName)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)

	// First line: network size
	if ok, line, err := nextLine(sc, 1); ok {
		data.Size = int(line[0])
	} else if err != nil {
		return nil, err
	}

	// Second line: scale factor
	if ok, line, err := nextLine(sc, 1); ok {
		data.ScaleFactor = line[0]
	} else if err != nil {
		return nil, err
	}

	ex := float64(1)
	if data.Size >= 170 {
		ex = 5
	} else if data.Size >= 70 {
		ex = 2
	}

	// Hub installation cost
	data.InstallationCost = make([]float64, data.Size)
	for i := 0; i < data.Size; i++ {
		if ok, line, err := nextLine(sc, 10000); ok {
			data.InstallationCost[i] = ex * line[0]
		} else if err != nil {
			return nil, err
		}
	}

	// Flow between nodes
	data.Flow = make([][]float64, data.Size, data.Size)
	for i := 0; i < data.Size; i++ {
		if ok, line, err := nextLine(sc, 100); ok {
			if len(line) != data.Size {
				return nil, fmt.Errorf("flow matrix should have dimension %dx%d", data.Size, data.Size)
			}
			data.Flow[i] = line
		} else if err != nil {
			return nil, err
		}
	}

	data.FlowDestiny = make([]float64, data.Size)
	data.FlowOrigin = make([]float64, data.Size)
	for i := 0; i < data.Size; i++ {
		for j := 0; j < data.Size; j++ {
			data.FlowOrigin[i] += data.Flow[i][j]
			data.FlowDestiny[i] += data.Flow[j][i]
		}
	}

	// Distance between nodes
	data.Distance = make([][]float64, data.Size, data.Size)
	for i := 0; i < data.Size; i++ {
		if ok, line, err := nextLine(sc, 100); ok {
			if len(line) != data.Size {
				return nil, fmt.Errorf("distance matrix should have size %d", data.Size)
			}
			data.Distance[i] = line
		} else if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (data *Data) Print() {
	fmt.Println("Instalation\n-----")
	for _, c := range data.InstallationCost {
		fmt.Printf("%.4f\n", c)
	}

	fmt.Println("\nFlow\n-----")
	for _, line := range data.Flow {
		fmt.Println(line)
	}

	fmt.Println("\nDistance\n-----")
	for _, line := range data.Distance {
		fmt.Println(line)
	}

	fmt.Println("\nOrigin Flow\n-----")
	for _, c := range data.FlowOrigin {
		fmt.Printf("%.4f\n", c)
	}

	fmt.Println("\nDestiny Flow\n-----")
	for _, c := range data.FlowDestiny {
		fmt.Printf("%.4f\n", c)
	}
}
