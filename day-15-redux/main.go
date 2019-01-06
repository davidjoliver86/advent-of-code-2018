package main

const (
	Wall   = '#'
	Goblin = 'G'
	Elf    = 'E'
	Space  = '.'
)

func main() {
	game := NewGame("test_board.txt")
	// fmt.Println(game)
	game.Turn()
	// elf := game.creatures[4]
	// goblin := game.creatures[0]
	// fmt.Println(elf.Reachable(Node{goblin.x, goblin.y}, game.Board()))
}
