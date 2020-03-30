package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"unicode"

	"github.com/namsral/flag"
	"github.com/wcharczuk/go-chart"
)

func main() {
	m := make(map[string]int)
	total := 0
	var file, graphname string
	flag.StringVar(&file, "f", "", "file which contain the data")
	flag.StringVar(&graphname, "o", "output", "name of the output file")
	flag.Parse()

	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error while opening the file: %v", err)
	}

	for x := 0; x < len(content); x++ {
		carac := rune(content[x])
		carac = unicode.ToLower(carac)
		if 97 <= int(carac) && int(carac) <= 122 {
			m[string(carac)]++
			total++
		}
	}

	fmt.Println(m)
	fmt.Println("[*] charaters counted: ", total)

	cvalues := []chart.Value{}
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		cvalues = append(cvalues, chart.Value{Value: float64(m[k] * 100 / total), Label: k})
	}

	graph := chart.BarChart{
		Title: "Occurences",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height:   512,
		BarWidth: 60,
		Bars:     cvalues,
	}

	f, _ := os.Create(graphname + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)

	fmt.Printf("[*] graph generated : %s.png\n", graphname)
}
