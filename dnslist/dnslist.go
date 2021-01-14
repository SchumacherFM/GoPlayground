package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	statTLD := map[string]int{}
	statDomain := map[string]int{}
	for sc.Scan() {
		txt := sc.Text()
		if strings.HasPrefix(txt, "#") {
			continue
		}
		data := strings.Split(txt, " ")
		if len(data) < 2 {
			println("NOT PARSED: ", txt)
			continue
		}
		domainSplit := strings.Split(data[1], ".")

		statTLD[domainSplit[len(domainSplit)-1]]++

		ds := len(domainSplit) - 2
		if ds < 0 {
			ds = 0
		}
		statDomain[strings.Join(domainSplit[ds:], ".")]++
	}
	reduce(statTLD)
	reduce(statDomain)

	sortAndPrint(statTLD)
	sortAndPrint(statDomain)
}

func reduce(stat map[string]int) {
	for txt, count := range stat {
		if count < 5 {
			delete(stat, txt)
		}
	}
}

func sortAndPrint(stat map[string]int) {
	temp := make([]string, 0, len(stat))
	for txt, count := range stat {
		temp = append(temp, fmt.Sprintf("%06d\t%s", count, txt))
	}
	sort.Strings(temp)
	for _, t := range temp {
		println(t)
	}
}
