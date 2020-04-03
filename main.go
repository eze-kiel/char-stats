package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"unicode"

	"github.com/namsral/flag"
	"github.com/wcharczuk/go-chart"

	log "github.com/sirupsen/logrus"
)

func main() {
	m := make(map[string]int)
	total := 0
	var file, graphname, layout string

	flag.StringVar(&file, "file", "", "file which contain the data")
	flag.StringVar(&graphname, "output", "graph.png", "name of the output file")
	flag.StringVar(&layout, "layout", "alpha", "values layout on the graph : alpha, asc")
	flag.Parse()

	fmt.Printf("[*] reading file...\n")
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error while opening the file: %v", err)
	}

	fmt.Printf("[*] processing...\n")
	// Collect characters between 'a' and 'z'
	for x := 0; x < len(content); x++ {
		carac := rune(content[x])
		carac = unicode.ToLower(carac)
		if 97 <= int(carac) && int(carac) <= 122 {
			m[string(carac)]++
			total++
		}
	}

	fmt.Printf("[*] charaters counted: %d\n", total)

	cvalues := []chart.Value{}

	keys, values := sortKeysAndValues(m)

	switch layout {
	case "alpha":
		for _, k := range keys {
			cvalues = append(cvalues, chart.Value{Value: float64(m[k] * 100 / total), Label: k})
		}

	case "asc":
		for _, v := range values {
			key := mapkey(m, v)
			cvalues = append(cvalues, chart.Value{Value: float64(v * 100 / total), Label: key})
		}
	}

	graphCreation(cvalues, graphname)

}

func mapkey(m map[string]int, value int) (key string) {
	for k, v := range m {
		if v == value {
			key = k
			delete(m, key) // Remove the entry to avoid duplicated labels
			return key
		}
	}
	return
}

func sortKeysAndValues(m map[string]int) ([]string, []int) {
	values := []int{}
	for v := range m {
		values = append(values, m[v])
	}
	sort.Ints(values)

	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys, values
}

func graphCreation(values []chart.Value, graphname string) {
	graph := chart.BarChart{
		Title: "Occurences (in %)",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		Height:   512,
		BarWidth: 30,
		Bars:     values,
	}

	f, _ := os.Create(graphname)
	defer f.Close()
	graph.Render(chart.PNG, f)

	fmt.Printf("[*] graph generated : %s\n", graphname)
}
