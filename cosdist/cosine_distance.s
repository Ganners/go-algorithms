#include "textflag.h"

// CosineDistanceAVX(a, b []float32) float32
TEXT ·CosineDistanceAVX(SB),NOSPLIT,$64
	MOVQ a+0(FP), AX
	MOVQ a_len+8(FP), BX
	MOVQ b+24(FP), CX
	MOVQ b_len+32(FP), DX

	CMPQ BX, DX
	JE compute

	// if they are not equal, return -1
	MOVSS $(-1.0), X0
	MOVD X0, ret+48(FP)
	RET

compute:
	VZEROALL
	SHRQ $3, BX

dotProducts:
	SUBQ $1, BX
	MOVQ BX, DX

	// multiply BX by 32
	SHLQ $5, DX

	// copy 8 packed singles from slice into Y0 and Y1
	// we'll work from right to left in the slice, shifting BX down to 0
	VMOVUPS 0(AX)(DX * 1), Y0
	VMOVUPS 0(CX)(DX * 1), Y1

	// dot product into lower and upper of y2 then sum lower and upper
	VDPPS $0xF7, Y0, Y0, Y2
	VPERM2F128 $1, Y2, Y2, Y8
	VADDPS Y2, Y8, Y2

	// dot product into lower and upper of y3 then sum lower and upper
	VDPPS $0xF7, Y1, Y1, Y3
	VPERM2F128 $1, Y3, Y3, Y8
	VADDPS Y3, Y8, Y3

	// dot product into lower and upper of y4 then sum lower and upper
	VDPPS $0xF7, Y0, Y1, Y4
	VPERM2F128 $1, Y4, Y4, Y8
	VADDPS Y4, Y8, Y4

	// add to totals computed so far
	VADDPS Y2, Y5, Y5
	VADDPS Y3, Y6, Y6
	VADDPS Y4, Y7, Y7

	CMPQ BX, $0
	JG dotProducts

	// multiply a dot a * b dot b
	VMULPS X5, X6, X0

	// take the square root of above (the norm)
	VSQRTPS X0, X0

	// divide X7 by the norm to leave the cosine in x0
	VDIVPS X0, X7, X0

	MOVD X0, ret+48(FP)
	RET

// DotAVX(a, b []float32) float32
TEXT ·DotAVX(SB),NOSPLIT,$64
	MOVQ a+0(FP), AX
	MOVQ a_len+8(FP), BX
	MOVQ b+24(FP), CX
	MOVQ b_len+32(FP), DX

	CMPQ BX, DX
	JE compute

	// if they are not equal, return -1
	MOVSS $(-1.0), X0
	MOVD X0, ret+48(FP)
	RET

compute:
	VZEROALL
	SHRQ $3, BX

dotProducts:
	SUBQ $1, BX
	MOVQ BX, DX

	// multiply BX by 32
	SHLQ $5, DX

	// copy 8 packed singles from slice into Y0 and Y1
	// we'll work from right to left in the slice, shifting BX down to 0
	VMOVUPS 0(AX)(DX * 1), Y0
	VMOVUPS 0(CX)(DX * 1), Y1

	// dot product into lower and upper of y4 then sum lower and upper
	VDPPS $0xF7, Y0, Y1, Y4
	VPERM2F128 $1, Y4, Y4, Y8
	VADDPS Y4, Y8, Y4

	// add to totals computed so far
	VADDPS Y4, Y7, Y7

	CMPQ BX, $0
	JG dotProducts

	MOVD X7, ret+48(FP)
	RET
