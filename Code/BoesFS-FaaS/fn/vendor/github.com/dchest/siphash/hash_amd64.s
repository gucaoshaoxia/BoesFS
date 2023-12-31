// +build amd64,!appengine,!gccgo

// This is a translation of the gcc output of FloodyBerry's pure-C public
// domain siphash implementation at https://github.com/floodyberry/siphash
// func Hash(k0, k1 uint64, b []byte) uint64
TEXT	·Hash(SB),4,$0-48
	MOVQ	k0+0(FP),CX
	MOVQ	$0x736F6D6570736575,R9
	MOVQ	k1+8(FP),DI
	MOVQ	$0x6C7967656E657261,BX
	MOVQ	$0x646F72616E646F6D,AX
	MOVQ	b_len+24(FP),DX
	MOVQ	DX,R11
	MOVQ	DX,R10
	XORQ	CX,R9
	XORQ	CX,BX
	MOVQ	$0x7465646279746573,CX
	XORQ	DI,AX
	XORQ	DI,CX
	SHLQ	$0x38,R11
	XORQ	DI,DI
	MOVQ	b_base+16(FP),SI
	ANDQ	$0xFFFFFFFFFFFFFFF8,R10
	JE	afterLoop
	XCHGQ	AX,AX
loopBody:
	MOVQ	0(SI)(DI*1),R8
	ADDQ	AX,R9
	RORQ	$0x33,AX
	XORQ	R9,AX
	RORQ	$0x20,R9
	ADDQ	$0x8,DI
	XORQ	R8,CX
	ADDQ	CX,BX
	RORQ	$0x30,CX
	XORQ	BX,CX
	ADDQ	AX,BX
	RORQ	$0x2F,AX
	ADDQ	CX,R9
	RORQ	$0x2B,CX
	XORQ	BX,AX
	XORQ	R9,CX
	RORQ	$0x20,BX
	ADDQ	AX,R9
	ADDQ	CX,BX
	RORQ	$0x33,AX
	RORQ	$0x30,CX
	XORQ	R9,AX
	XORQ	BX,CX
	RORQ	$0x20,R9
	ADDQ	AX,BX
	ADDQ	CX,R9
	RORQ	$0x2F,AX
	RORQ	$0x2B,CX
	XORQ	BX,AX
	RORQ	$0x20,BX
	XORQ	R9,CX
	XORQ	R8,R9
	CMPQ	R10,DI
	JA	loopBody
afterLoop:
	SUBQ	R10,DX

	CMPQ	DX,$0x7
	JA	afterSwitch

	// no support for jump tables

	CMPQ	DX,$0x7
	JE	sw7

	CMPQ	DX,$0x6
	JE	sw6

	CMPQ	DX,$0x5
	JE	sw5

	CMPQ	DX,$0x4
	JE	sw4

	CMPQ	DX,$0x3
	JE	sw3

	CMPQ	DX,$0x2
	JE	sw2

	CMPQ	DX,$0x1
	JE	sw1

	JMP	afterSwitch

sw7:	MOVBQZX	6(SI)(DI*1),DX
	SHLQ	$0x30,DX
	ORQ	DX,R11
sw6:	MOVBQZX	0x5(SI)(DI*1),DX
	SHLQ	$0x28,DX
	ORQ	DX,R11
sw5:	MOVBQZX	0x4(SI)(DI*1),DX
	SHLQ	$0x20,DX
	ORQ	DX,R11
sw4:	MOVBQZX	0x3(SI)(DI*1),DX
	SHLQ	$0x18,DX
	ORQ	DX,R11
sw3:	MOVBQZX	0x2(SI)(DI*1),DX
	SHLQ	$0x10,DX
	ORQ	DX,R11
sw2:	MOVBQZX	0x1(SI)(DI*1),DX
	SHLQ	$0x8,DX
	ORQ	DX,R11
sw1:	MOVBQZX	0(SI)(DI*1),DX
	ORQ	DX,R11
afterSwitch:
	LEAQ	(AX)(R9*1),SI
	XORQ	R11,CX
	RORQ	$0x33,AX
	ADDQ	CX,BX
	MOVQ	CX,DX
	XORQ	SI,AX
	RORQ	$0x30,DX
	RORQ	$0x20,SI
	LEAQ	0(BX)(AX*1),CX
	XORQ	BX,DX
	RORQ	$0x2F,AX
	ADDQ	DX,SI
	RORQ	$0x2B,DX
	XORQ	CX,AX
	XORQ	SI,DX
	RORQ	$0x20,CX
	ADDQ	AX,SI
	RORQ	$0x33,AX
	ADDQ	DX,CX
	XORQ	SI,AX
	RORQ	$0x30,DX
	RORQ	$0x20,SI
	XORQ	CX,DX
	ADDQ	AX,CX
	RORQ	$0x2F,AX
	ADDQ	DX,SI
	XORQ	CX,AX
	RORQ	$0x2B,DX
	RORQ	$0x20,CX
	XORQ	SI,DX
	XORQ	R11,SI
	XORB	$0xFF,CL
	ADDQ	AX,SI
	RORQ	$0x33,AX
	ADDQ	DX,CX
	RORQ	$0x30,DX
	XORQ	SI,AX
	XORQ	CX,DX
	RORQ	$0x20,SI
	ADDQ	AX,CX
	ADDQ	DX,SI
	RORQ	$0x2F,AX
	RORQ	$0x2B,DX
	XORQ	CX,AX
	XORQ	SI,DX
	RORQ	$0x20,CX
	ADDQ	AX,SI
	ADDQ	DX,CX
	RORQ	$0x33,AX
	RORQ	$0x30,DX
	XORQ	SI,AX
	RORQ	$0x20,SI
	XORQ	CX,DX
	ADDQ	AX,CX
	RORQ	$0x2F,AX
	ADDQ	DX,SI
	RORQ	$0x2B,DX
	XORQ	CX,AX
	XORQ	SI,DX
	RORQ	$0x20,CX
	ADDQ	AX,SI
	ADDQ	DX,CX
	RORQ	$0x33,AX
	RORQ	$0x30,DX
	XORQ	CX,DX
	XORQ	SI,AX
	RORQ	$0x20,SI
	ADDQ	DX,SI
	ADDQ	AX,CX
	RORQ	$0x2F,AX
	XORQ	CX,AX
	RORQ	$0x2B,DX
	RORQ	$0x20,CX
	XORQ	SI,DX
	ADDQ	AX,SI
	RORQ	$0x33,AX
	ADDQ	DX,CX
	XORQ	SI,AX
	RORQ	$0x30,DX
	XORQ	CX,DX
	ADDQ	AX,CX
	RORQ	$0x2F,AX
	XORQ	CX,AX
	RORQ	$0x2B,DX
	RORQ	$0x20,CX
	XORQ	DX,AX
	XORQ	CX,AX
	MOVQ	AX,ret+40(FP)
	RET
