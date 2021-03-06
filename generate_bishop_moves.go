package maurice

func (p Position) bishopMoves() []Move {
	var friendly = p.attackers()
	var bb = p.PieceBitboards[Bishop+p.color]
	var occ = p.occupied()
	var moves = make([]Move, 0)

	for bb > 0 {
		square := bb & -bb
		squareNum := square.FirstSquare()
		bb &= bb - 1

		blocker := occ & bishopMagic[squareNum].mask
		index := (blocker * bishopMagic[squareNum].magic) >> 55
		moveBb := bishopMagicMoves[squareNum][index]

		legalMovesBb := moveBb & (^friendly)

		newMoves := p.movesFromBitboard(legalMovesBb, func(_ Bitboard) Bitboard {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func (p Position) bishopAttacks(color int) Bitboard {
	var bb Bitboard = p.PieceBitboards[Bishop+color]
	var occ Bitboard = p.occupied()
	var attackBB Bitboard = 0

	for bb > 0 {
		square := bb & -bb
		squareNum := square.FirstSquare()
		bb &= bb - 1

		blocker := occ & bishopMagic[squareNum].mask
		index := (blocker * bishopMagic[squareNum].magic) >> 55
		moveBb := bishopMagicMoves[squareNum][index]

		attackBB |= moveBb
	}

	return attackBB
}
