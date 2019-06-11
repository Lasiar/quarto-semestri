package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/Lasiar/quarto-semestri/TIPS/graph"
	"github.com/Lasiar/quarto-semestri/TIPS/grid"
	"log"
	"os"

	"strconv"
	"strings"
	"time"
)

var (
	start, end, banned string
	def                bool
)

func init() {
	flag.StringVar(&start, "start", "71,71", "")
	flag.StringVar(&end, "end", "0,0", "")
	flag.StringVar(&banned, "banned", "", "")
	flag.BoolVar(&def, "default", true, "")
	flag.Parse()
}

func main() {

	s, e, b := parseInital()
	if err := validCord(s, e, b); err != nil {
		log.Println(err)
		return
	}

	g, startOnGrid, endOnGrid := grid.New(72, s, e, b)

	if err := g.Print(os.Stdout); err != nil {
		log.Fatalf("error write to file %v", err)
	}

	dt := time.Now()

	find := false

	i := 0
	g.TraverseWithStart(startOnGrid, func(node *graph.Node) bool {
		if find {
			return false
		}

		if *endOnGrid == *node {
			find = true
		}
		fmt.Printf("cord: %v,\titer: %v,\ttime duration: %v\n", node.Value.Cord, i, time.Since(dt))
		i++
		return true
	})
	if !find {
		log.Printf("end node not available with start node")
	}
	fmt.Println("summary time: ", time.Since(dt))
}

func getUserInitialParameters() (start, end [2]int, banned [][2]int) {
	start = getOnceCord("Enter start node\n x,y split coma (0,0) -> ")
	end = getOnceCord("Enter finish node -> ")
	banned = getSliceCord("Enter banned node separated semicolon ->")
	return start, end, banned
}

func getOnceCord(text string) (cord [2]int) {
	var err error
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	for sc.Scan() {
		cord, err = parseCord(sc.Text())
		if err != nil {
			log.Println(err)
			continue
		}
		break
	}
	return cord
}

func getSliceCord(text string) (cord [][2]int) {
	var err error
	cord = make([][2]int, 0)
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	for sc.Scan() {
		trimmedSpace := strings.TrimSpace(sc.Text())

		if trimmedSpace == "" {
			return nil
		}

		raw := strings.Split(trimmedSpace, ";")
		for _, r := range raw {
			var c [2]int
			c, err = parseCord(r)
			cord = append(cord, c)
		}
		if err != nil {
			log.Println(err)
			continue
		}
		break
	}
	return cord
}

func parseCord(raw string) ([2]int, error) {
	trimmedSpace := strings.TrimSpace(raw)
	rp := strings.Split(trimmedSpace, ",")
	ff, err := strconv.Atoi(rp[0])
	if err != nil {
		return [2]int{}, err
	}
	fs, err := strconv.Atoi(rp[1])
	if err != nil {
		return [2]int{}, err
	}
	return [2]int{ff, fs}, nil
}

func validCord(start, end [2]int, banned [][2]int) error {
	if start == end {
		return errors.New("error: start == end")
	}

	for _, b := range banned {
		if start == b {
			return fmt.Errorf("error: start: %v == banned: %v", start, b)
		}

		if end == b {
			return fmt.Errorf("error: end: %v == banned %v", end, b)
		}
	}
	return nil
}

func parseInital() (s, e [2]int, b [][2]int) {
	if def {
		cord, err := parseCord(start)
		if err != nil {
			log.Printf("Parse start: %v", err)
			os.Exit(1)
		}
		s = cord
		cord, err = parseCord(end)
		if err != nil {
			log.Printf("Parse end: %v", err)
		}
		e = cord
		bannedTrimmedSpace := strings.TrimSpace(banned)

		if bannedTrimmedSpace == "" {
			b = nil
			return s, e, b
		}

		raw := strings.Split(bannedTrimmedSpace, ";")
		for _, r := range raw {
			var c [2]int
			c, err = parseCord(r)
			b = append(b, c)
		}
		if err != nil {
			log.Println(err)
		}
		return s, e, b
	}
	return getUserInitialParameters()
}
