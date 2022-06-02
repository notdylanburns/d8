# D8 Processesor Specification

The D8 CPU architecture is an 8-bit architecture featuring four 8-bit general purpose registers - `a`, `b`, `c`, and `d`, a 16-bit program counter and stack pointer (`pc` and `sp` respectively). Furthermore, several other 16-bit address registers are present, namely the return address register (`ra`), the source/destination index (`si` and `di`) and the transfer register (`tx`) which exists as two 8-bit registers (`th` and `tl`), and can be used to move values between the system's main 8-bit bus and the 16-bit address bus.

The CPU uses pipelining to increase efficiency. The instruction pipeline has four stages, each of which lasts four cycles of the CPU clock. On the fifth cycle, each stage of the pipeline is shifted into the next. Using this method, the CPU can effectively execute four instructions in 20 cycles, or one instruction every five cycles if they weren't parellelised.

ctrl bit outputs from pipeline
[0:1] = cycle
[2:9] = instruction
[10:13] = mode
[14:17] = flags

eg for instruction 0x00, in mode 0, with flags=0b0000
0 0000: inc pc,
0 0001: load const, inc pc
0 0002: load const, inc pc
0 0003: nop

pipeline:
    stage 1:
        0:
            - inc pc
        1: 
            - load const
            - load mode byte
            - inc pc
        2:
            - load const
            - inc pc
        3:

