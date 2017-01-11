package timeconversion

import (
	"fmt"
	"log"
	"time"
)

var layout = "03:05:04PM"

func ConvertTime(input string) string {
	t, err := time.Parse(layout, input)
	if err != nil {
		log.Printf("Error parsing time: %v", err)
	}
	return t.Format("15:05:04")
}

func main() {
	var inputTime string
	fmt.Scanf("%v", &inputTime)
	outputTime := ConvertTime(inputTime)
	fmt.Println(outputTime)
}
