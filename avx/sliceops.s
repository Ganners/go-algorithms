#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

// Sum(f []int) int
TEXT ·Sum(SB),NOSPLIT,$0
	MOVQ f+0(FP), AX
	MOVL len+8(FP), DX
	XORQ BX, BX
	XORQ CX, CX

ADD_PREV:
	DECQ DX

	MOVQ (AX)(DX * 8), CX
	ADDQ CX, BX

	CMPQ DX, $0
	JNE ADD_PREV

RETURN:
	MOVQ BX, ret+24(FP)
	RET


// SumFloat(f []float64) float64
TEXT ·SumFloat(SB),NOSPLIT,$0
	MOVQ f+0(FP), AX
	MOVL len+8(FP), DX
	XORPS X0, X0
	XORL CX, CX

ADD_PREV:
	DECL DX

	MOVSD (AX)(DX * 8), X1
	ADDSD X1, X0

	CMPL DX, $0
	JNE ADD_PREV

RETURN:
	MOVSD X0, ret+24(FP)
	RET


// SumFloatAVX(f []float64) float64
TEXT ·SumFloatAVX(SB),NOSPLIT,$0
	MOVQ f+0(FP), AX
	MOVL len+8(FP), DX

	VMOVUPD 0(AX), Y0
	VMOVUPD 16(AX), Y1

	VHADDPD Y0, Y1, Y1
	VMOVUPD Y1, ret+24(FP)

	RET
