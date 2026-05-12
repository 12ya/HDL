/**
    *LOOP:
    * read keyboard
    * if no key -> color = 0
    * if key    -> color = -1
    * reset addr to SCREEN
    * write color from SCREEN until KBD
    * repeat
*/

(LOOP)
@KBD
D=M

@WHITE
D;JEQ

@BLACK
0;JMP

(WHITE)
@color
M=0
@DRAW
0;JMP

(BLACK)
@color
M=-1
@DRAW
0;JMP

(DRAW)
@SCREEN
D=A
@addr
M=D

(DRAW_LOOP)
@color
D=M
@addr
A=M
M=D

@addr
M=M+1

@addr
D=M
@KBD
D=D-A
@DRAW_LOOP
D;JLT

@LOOP
0;JMP