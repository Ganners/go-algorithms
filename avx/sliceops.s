#include "textflag.h"

// SumFloatAVX(f []float64) float64
TEXT ·SumFloatAVX(SB),NOSPLIT,$0-32
	MOVQ f+0(FP), AX
	MOVQ len+8(FP), BX

	// clear 256 bit registers that we'll be using
	VPXOR Y0, Y0, Y0
	VPXOR Y1, Y1, Y1
	VPXOR Y2, Y2, Y2
	VPXOR X5, X5, X5

	MOVQ BX, CX
	SHRQ $3, CX

sumNext:
	SUBQ $1, CX

	MOVQ CX, DX
	SHLQ $6, DX

	VMOVUPD 0(AX)(DX * 1), Y0
	VMOVUPD 32(AX)(DX * 1), Y1

	// Y0 = 1, 2, 3, 4
	// Y1 = 5, 6, 7, 8

	// HADD(Y0, Y1) = 1 + 2 | 5 + 6 | 3 + 4 | 7 + 8
	//              = 3 | 11 | 7 | 15
	VHADDPD Y0, Y1, Y1

	// PERM2F128(1, Y1, Y1) = 15 | 7 | 11 | 3
	VPERM2F128 $1, Y1, Y1, Y2

	// HADD(Y1, Y2) = 3 + 11 | 15 + 7 | 7 + 15 | 11 + 3
	//              = 14 | 22 | 22 | 14
	VHADDPD Y1, Y2, Y2
	VHADDPD Y2, Y2, Y2
	ADDSD X2, X5

	CMPQ CX, $0
	JG sumNext

	MOVQ X5, ret+24(FP)

	RET

// SumFloatAVX(f []float32) float32
TEXT ·SumFloat32AVX(SB),NOSPLIT,$0-32
	MOVQ f+0(FP), AX
	MOVQ len+8(FP), BX

	// clear 256 bit registers that we'll be using
	VPXOR Y0, Y0, Y0
	VPXOR Y1, Y1, Y1
	VPXOR Y2, Y2, Y2
	VPXOR Y3, Y3, Y3
	VPXOR Y4, Y4, Y4
	VPXOR Y5, Y5, Y5
	VPXOR Y6, Y6, Y6
	VPXOR Y7, Y7, Y7
	VPXOR Y8, Y8, Y8
	VPXOR X5, X5, X5
	VPXOR X6, X6, X6
	VPXOR X7, X7, X7

	MOVQ BX, CX
	SHRQ $4, CX // we can do this 16 at a time boi

sumNext:
	SUBQ $1, CX

	MOVQ CX, DX
	SHLQ $6, DX

	VMOVUPS 0(AX)(DX * 1), Y0
	VMOVUPS 32(AX)(DX * 1), Y1

	// Y0 = 1, 2, 3, 4, 5, 6, 7, 8
	// Y1 = 9, 10, 11, 12, 13, 14, 15, 16

	// Y0 = A, B, C, D, E, F, G, H
	// Y1 = I, J, K, L, M, N, O, P

	// HADD(Y0, Y1) = 1 + 2 | 9 + 10 | 3 + 4 | 11 + 12 | 5 + 6 | 13 + 14 | 7 + 8 | 15 + 16
	//              = 3 | 19 | 7 | 23 | 11 | 27 | 15 | 31

	// HADD(Y1, Y0) = A + B | I + J | C + D | K + L | E + F | M + N | G + H | O + P
	VHADDPS Y1, Y0, Y1

	// PERM2F128 = E + F | M + N | G + H | O + P | A + B | I + J | C + D | K + L
	VPERM2F128 $1, Y1, Y1, Y2

	// A + B | I + J | C + D | K + L | E + F | M + N | G + H | O + P
	// E + F | M + N | G + H | O + P | A + B | I + J | C + D | K + L
	//
	// Add: A + B + E + F | I + J + M + N | C + D + G + H | K + L + O + P | ...
	VADDPS Y1, Y2, Y3

	// Shuffle: K + L + O + P | C + D + G + H | I + J + M + N | A + B + E + F | ...
	VSHUFPS $0x1b, Y3, Y3, Y4

	// A + B + E + F + K + L + O + P | I + J + M + N + C + D + G + H | ...
	VADDPS Y3, Y4, Y5

	// A + B + E + F + K + L + O + P + I + J + M + N + C + D + G + H | ...
	VHADDPS Y5, Y5, Y6

	// accumulate
	ADDSS X6, X7

	CMPQ CX, $0
	JG sumNext

	MOVQ X7, ret+24(FP)

	RET
