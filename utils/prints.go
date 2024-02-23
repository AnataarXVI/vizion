package utils

import "fmt"

const (

	// Colors
	RED    string = "\033[31m"
	GREEN  string = "\033[32m"
	YELLOW string = "\033[33m"
	BLUE   string = "\033[34m"
	PURPLE string = "\033[35m"
	CYAN   string = "\033[36m"
	WHITE  string = "\033[37m"

	// Styles
	BOLD      string = "\033[1m"
	UNDERLINE string = "\033[4m"
	STRIKE    string = "\033[9m"
	ITALIC    string = "\033[3m"

	RESET string = "\033[0m"
)

func Display_Layer(layerName string) {
	fmt.Printf("▼ [%s %s %s]\n", RED, layerName, RESET)
}

func Display_Fields(fieldName string, fieldValue any, fieldEnum any) {
	if fieldEnum != nil {
		fmt.Printf("\t%s%-6s%s  = %s %v %s %s( 0x%x )%s \n", PURPLE, fieldName, RESET, YELLOW, fieldEnum, RESET, GREEN, fieldValue, RESET)
	} else {
		fmt.Printf("\t%s%-6s%s  = %s %v %s\n", PURPLE, fieldName, RESET, YELLOW, fieldValue, RESET)
	}
}
