package main 
import (
	"fmt"
	"os"
	"flag"
	"encoding/csv"
)
func main(){
	csvFilename := flag.String("csv" , "problems.csv" , "a csv file in the format of 'question , answer'")
	flag.Parse()
	file , err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file:%s\nz",*csvFilename))
	}
	
	// io reader , google it 
	r := csv.NewReader(file)
	lines , err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file")}
		//fmt.Println(lines)
		problems := parseLines(lines)
			fmt.Println(problems)
		}

func parseLines(lines [][]string )[]problem{
	ret := make([]problem , len(lines)) //make function make dynamic size array 
	for i ,line := range lines{
		ret[i] = problem{
			question : line[0],
			answer : line[1],
		}
	}
	return ret
}

type problem struct {
	question string 
	answer string 
}
func exit (msg string){
	fmt.Println(msg)
	os.Exit(1)

}