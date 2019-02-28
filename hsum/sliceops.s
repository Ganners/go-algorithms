#include "textflag.h"

// SumFloatAVX(f []float64) float64
TEXT ·SumFloatAVX(SB),NOSPLIT,$0-32
	MOVQ f+0(FP), AX
	MOVQ len+8(FP), BX

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

	VMOVUPD 0(AX)(DX * 1), Y2
	VMOVUPD 32(AX)(DX * 1), Y3
	VHADDPD Y0, Y2, Y0
	VHADDPD Y1, Y3, Y1

	CMPQ CX, $0
	JG sumNext

	VHADDPD Y0, Y1, Y1
	VPERM2F128 $1, Y1, Y1, Y2
	VHADDPD Y1, Y2, Y2
	VHADDPD Y2, Y2, Y2

	MOVQ X2, ret+24(FP)

	RET

// SumFloatAVX(f []float32) float32
TEXT ·SumFloat32AVX(SB),NOSPLIT,$0-32
	MOVQ f+0(FP), AX
	MOVQ len+8(FP), BX

	VPXOR Y0, Y0, Y0
	VPXOR Y1, Y1, Y1
	VPXOR Y2, Y2, Y2
	VPXOR Y3, Y3, Y3
	VPXOR X2, X2, X2

	MOVQ BX, CX
	SHRQ $4, CX

sumNext:
	SUBQ $1, CX

	MOVQ CX, DX
	SHLQ $6, DX

	// Y2 = A, B, C, D, E, F, G, H
	// Y3 = I, J, K, L, M, N, O, P
	VMOVUPS 0(AX)(DX * 1), Y2
	VMOVUPS 32(AX)(DX * 1), Y3

	VHADDPS Y0, Y2, Y0
	VHADDPS Y1, Y3, Y1

	CMPQ CX, $0
	JG sumNext // grab more to sum up

	// Horizontal Add: A + B | I + J | C + D | K + L | E + F | M + N | G + H | O + P
	// Permute = E + F | M + N | G + H | O + P | A + B | I + J | C + D | K + L
	//
	// Then we have the following:
	//
	// A + B | I + J | C + D | K + L | E + F | M + N | G + H | O + P
	// E + F | M + N | G + H | O + P | A + B | I + J | C + D | K + L
	//
	// Add: A + B + E + F | I + J + M + N | C + D + G + H | K + L + O + P | ...
	//
	// Then we end up with two similar packed singles, so we need to rotate the
	// first and second two singles:
	//
	// Shuffle: K + L + O + P | C + D + G + H | I + J + M + N | A + B + E + F | ...
	//
	// Then we can close them all up
	//
	// Add V: A + B + E + F + K + L + O + P | I + J + M + N + C + D + G + H | ...
	// Add H: A + B + E + F + K + L + O + P + I + J + M + N + C + D + G + H | ...
	//
	// Then we just pick the lower single from X2 to return
	VHADDPS Y0, Y1, Y1
	VPERM2F128 $1, Y1, Y1, Y2
	VADDPS Y1, Y2, Y2
	VSHUFPS $0x1b, Y2, Y2, Y1
	VADDPS Y1, Y2, Y2
	VHADDPS Y2, Y2, Y2

	MOVQ X2, ret+24(FP)
	RET

// IDK(a, b []float64) []float64
TEXT ·IDK(SB),NOSPLIT,$0-64
	// slice A
	MOVQ a+0(FP), AX

	// slice B
	MOVQ a+24(FP), BX

	VMOVUPD (AX), Y0
	VMOVUPD (BX), Y1

	// To sum them
	// VADDPD Y0, Y1, Y1

	// To divide them
	// VDIVPD Y1, Y0, Y1

	// To multiply them
	// VMULPD Y0, Y1, Y1

	// To compute the max
	// VMAXPD Y0, Y1, Y1

	// To compute the min
	// VMINPD Y0, Y1, Y1

	// To round (.5 is ceil)
	// VDIVPD Y1, Y0, Y1
	// VADDPD Y1, Y1, Y1
	// VADDPD Y1, Y1, Y1
	// VROUNDPD $0, Y1, Y1

	// floor
	// VDIVPD Y1, Y0, Y1
	// VADDPD Y1, Y1, Y1
	// VADDPD Y1, Y1, Y1
	// VROUNDPD $1, Y1, Y1

	// ceil
	// VDIVPD Y1, Y0, Y1
	// VADDPD Y1, Y1, Y1
	// VADDPD Y1, Y1, Y1
	// VROUNDPD $2, Y1, Y1

	// sqrt
	// VSQRTPD Y1, Y1

	// make some integers n convert to floats
	// MOVQ $10, 0(CX)
	// MOVQ $20, 4(CX)
	// MOVQ $30, 8(CX)
	// MOVQ $40, 12(CX)
	// VCVTDQ2PD (CX), Y1

	// how do we do this :thinking:
	// RDRAND (AX)

	VMOVUPD Y1, (AX)

	MOVQ AX, ret+48(FP)
	MOVQ $4, ret+56(FP)
	MOVQ $4, ret+64(FP)
	RET
