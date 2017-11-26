package maurice

import (
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

func perft() {
	b, err := ioutil.ReadFile( "perft.epd")

	if err != nil {
		log.Fatal(err)
	}

	positions := strings.Split(string(b), "\n")
	positive := true

	for p := 0; p < len(positions) && positive; p ++ {
		position := positions[p]

		parts := strings.Split(position, ";")
		fen := parts[0]

		fmt.Println(p, "Checking board", fen)

		board := PositionFromFEN(fen)

		for i := 1; i < len(parts) && positive; i ++ {
			part := parts[i]
			expectedNumMoves , _ := strconv.Atoi(strings.Split(part, " ")[1])

			start := time.Now()
			actualNumMoves := board.Perft(i)
			end := time.Now()

			if expectedNumMoves != actualNumMoves {
				positive = false
			}

			fmt.Println(expectedNumMoves == actualNumMoves, i, expectedNumMoves, actualNumMoves, end.Sub(start))
		}
		fmt.Println("----------------")
	}

}

func (p Position) Divide (d int) error {
	sum := 0
	for _, m := range p.legalMoves() {
		nb := p.Move(m)
		num := nb.Perft(d - 1)
		sum += num
		fmt.Println(m.toString(), num)
	}
	fmt.Println("Moves:", sum)

	return nil
}

func (p Position)Perft(d int) int {
	if d == 0 {
		return 1
	}
	nodes := 0
	px := &p

	for _, m := range p.pseudoMoves() {
		nb := px.Move(m)
		if !nb.isInCheck() {
			nodes += nb.Perft(d - 1)
		}
	}

	return nodes
}