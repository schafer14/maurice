package main


import (
	"github.com/schafer14/maurice"
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
	"time"
)

var (
	position = "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1"
)

func main() {


	p := maurice.PositionFromFEN(position)

	reader := bufio.NewReader(os.Stdin)

	REPL:
		for {
			text, _ := reader.ReadString('\n');
			args := strings.Split(text, " ")

			switch cmd := strings.TrimSpace(args[0]); cmd {
			case "q":
				break REPL
			case "m":
				fmt.Println(p.HumanFriendlyMoves())
			case "d":
				p.Draw()
			case "r":
				p = maurice.PositionFromFEN(position)
			case "s":
				fmt.Println(p.ToFen())
			case "divide":
				i, _ := strconv.Atoi(strings.TrimSpace(args[1]))
				p.Divide(i)
			case "perft":
				i, _ := strconv.Atoi(strings.TrimSpace(args[1]))
				start := time.Now()
				n := p.Perft(i)
				finish := time.Since(start)
				fmt.Printf("  Depth: %d\n", i)
				fmt.Printf("  Nodes: %d\n", n)
				fmt.Println("Elapsed:", finish)
				fmt.Printf("Nodes/s: %dK\n", int(float64(n) / finish.Seconds() / 1000))
			default:
				fmt.Println("Moving")
				p = p.HumanFriendlyMove(text)
			}
		}
}

func ms(duration int64) string {
	mm := duration / 1000 / 60
	ss := duration / 1000 % 60
	xx := duration - mm * 1000 * 60 - ss * 1000

	return fmt.Sprintf(`%02d:%02d.%03d`, mm, ss, xx)
}
