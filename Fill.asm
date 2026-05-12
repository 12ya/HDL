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

@LOOP
0;JMP