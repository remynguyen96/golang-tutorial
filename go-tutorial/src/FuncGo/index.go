package FuncGo
import (
"errors"
"fmt"
"strings"
)

func MainFunc() {
	plants := []PowerPlant{
		{hydro, 300, active},
		{wind, 100, unavailable},
		{solar, 200, inactiver},
		{wind, 400, unavailable},
		{hydro, 500, inactiver},
		{solar, 600, active},
	}
	grid := PowerGrid{300, plants}

	//plantCapacities := []float64{10, 20, 30, 40, 50}
	//activePlants := []int{0, 1}
	//gridLoad := 75.

	if option, err := requestOption(); err == nil {
		fmt.Println("")
		switch option {
		case "1":
			grid.generatePlantReport()
			//generatePlantCapacityReport()
		case "2":
			grid.generateGridReport()
			//generatePowerGridReport(activePlants, plantCapacities, gridLoad)
		}
	} else {
		fmt.Println(err.Error())
	}
}

func requestOption() (option string, err error) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Gird Report")
	fmt.Print("Please choose an option: ")

	fmt.Scanln(&option)

	if option != "1" && option != "2" {
		err = errors.New("Invalid option selected")
	}

	return
}

func generatePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}

	fmt.Printf("%-20s%.0f\n", "Capacity", capacity)
	fmt.Printf("%-20s%.0f\n", "Load", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization", gridLoad/capacity*100)
}

type PlantType string
type PlantStatus string
const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

const (
	active      PlantStatus = "Active"
	inactiver   PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type:", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity:", p.capacity)
		fmt.Printf("%-20s%s\n", "Status:", p.status)
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}

	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity:", capacity)
	fmt.Printf("%-20s%.0f\n", "Load:", pg.load)
	fmt.Printf("%-20s%.2f%%\n", "Utilization:", pg.load/capacity*100)
}

func filterCharacter(arrCharacter ...string) (object map[string][]string) {
	object = make(map[string][]string)
	for _, val := range arrCharacter {
		if strings.ToUpper(val) == val {
			object["uppercase"] = append(object["uppercase"], val)
		} else if strings.ToLower(val) == val {
			object["lowercase"] = append(object["lowercase"], val)
		} else {
			object["propercase"] = append(object["propercase"], val)
		}
	}
	return
	//valTest := []string{"Hello", "World", "good", "MORNING"}
	//fmt.Println(filterCharacter(valTest...))
}
