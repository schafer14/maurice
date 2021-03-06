package maurice

func (p Position) queenMoves() []Move {
	var friendly = p.attackers()
	var bb = p.PieceBitboards[Queen+p.color]
	var occ = p.occupied()
	var moves = make([]Move, 0)

	for bb > 0 {
		square := bb & -bb
		squareNum := square.FirstSquare()
		bb &= bb - 1

		blockerB := occ & (bishopMagic[squareNum].mask)
		indexB := (blockerB * bishopMagic[squareNum].magic) >> 55
		moveBb := bishopMagicMoves[squareNum][indexB]

		blockerR := occ & (rookMagic[squareNum].mask)
		indexR := (blockerR * rookMagic[squareNum].magic) >> 52
		moveBb |= rookMagicMoves[squareNum][indexR]

		legalMovesBb := moveBb & (^friendly)

		newMoves := p.movesFromBitboard(legalMovesBb, func(_ Bitboard) Bitboard {
			return square
		})

		moves = append(moves, newMoves...)
	}

	return moves
}

func (p Position) queenAttacks(color int) Bitboard {
	bb := p.PieceBitboards[Queen+color]
	occ := p.occupied()
	attackBB := Bitboard(0)

	for bb > 0 {
		square := bb & -bb
		squareNum := square.FirstSquare()
		bb &= bb - 1

		blockerB := occ & (bishopMagic[squareNum].mask)
		indexB := (blockerB * bishopMagic[squareNum].magic) >> 55
		moveBb := bishopMagicMoves[squareNum][indexB]

		blockerR := occ & (rookMagic[squareNum].mask)
		indexR := (blockerR * rookMagic[squareNum].magic) >> 52
		moveBb |= rookMagicMoves[squareNum][indexR]

		attackBB |= moveBb
	}

	return attackBB
}
