#include "textflag.h"

// sumFloatAVX(f []float64) float64
TEXT ·sumFloatAVX(SB),NOSPLIT,$0-32
	MOVQ f+0(FP), AX
	MOVQ len+8(FP), BX


	// clear 256 bit registers that we'll be using
	VPXOR Y0, Y0, Y0
	VPXOR Y1, Y1, Y1
	VPXOR Y2, Y2, Y2

	VMOVUPD (AX), Y0
	VMOVUPD (32)(AX), Y1

	VADDPD Y0, Y1, Y1
	VPERM2F128 $1, Y1, Y1, Y2
	VADDPD Y1, Y2, Y2
	VHADDPD Y2, Y2, Y2

	MOVQ X2, ret+24(FP)

	RET
