# D8 Processesor Specification

The D8 CPU architecture is an 8-bit architecture featuring four 8-bit general purpose registers - `a`, `b`, `c`, and `d`, a 16-bit program counter and stack pointer (`pc` and `sp` respectively). Furthermore, several other 16-bit address registers are present, namely the return address register (`ra`), the source/destination index (`si` and `di`) and the transfer register (`tx`) which exists as two 8-bit registers (`th` and `tl`), and can be used to move values between the system's main 8-bit bus and the 16-bit address bus.

The CPU uses pipelining to increase efficiency. The instruction pipeline has four stages, each of which lasts four cycles of the CPU clock. On the fifth cycle, each stage of the pipeline is shifted into the next. Using this method, the CPU can effectively execute four instructions in 20 cycles, or one instruction every five cycles if they weren't parellelised.

## Registers

| Name              | Symbol | Bits | ++ | -- | Load |
|-------------------|--------|------|----|----|------|
| Program Counter   | pc     | 16   | ✓  |    | ✓    |
| Stack Pointer     | sp     | 16   | ✓  | ✓  | ✓    |
| Return Address    | ra     | 16   |    |    | ✓    |
| Source Index      | si     | 16   | ✓  | ✓  | ✓    |
| Destination Index | di     | 16   | ✓  | ✓  | ✓    |
| Transfer Register | tx     | 16   |    |    | ✓    |
| Transfer Lower    | tl     | 8    |    |    | ✓    |
| Transfer Higher   | th     | 8    |    |    | ✓    |
| A Register        | a      | 8    | ✓  | ✓  | ✓    |    
| B Register        | b      | 8    | ✓  | ✓  | ✓    |  
| C Register        | c      | 8    | ✓  | ✓  | ✓    |  
| D Register        | d      | 8    | ✓  | ✓  | ✓    |  

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


Control Lines: 23
    AU_WN0
    AU_WN1
    AU_WN2
    AU_RN0
    AU_RN1
    AU_RN2
    AU_LOAD_TH
    AU_LOAD_TL
    PC_CE#
    SP_U/D#
    SP_CE#
    RA_CE#
    SI_U/D#
    SI_CE#
    DI_U/D#
    DI_CE#
    CR_ASSERT_MAIN
    CR_LOAD_LOW
    CR_LOAD_HIGH
    GPR_W_N0
    GPR_W_N1
    GPR_R_N0
    GPR_R_N1
    GPR_WE_L
    GPR_OE_L
    GPR_LOAD_L
    GPR_U_DL
    GPR_CE_L

mov a,#:
    0: AU_W_N0=0, AU_W_N1=1, AU_W_N2=1,


Constant Register:
    Control Lines: 3
        CR_ASSERT_MAIN: Asserts the low 8 bits of the constant register onto the main bus
        CR_LOAD_LOW: Loads the low 8 bits of the constant register from the main bus
        CR_LOAD_HIGH: Loads the high 8 bits of the constant register from the main bus

Address Unit:
    Control Lines: 16
        W_N[3]: The device to load into
            000: none
            001: none
            010: di
            011: si
            100: ra
            101: sp
            110: pc
            111: tx

        R_N[3]: The device to output onto the address bus
            000: none
            001: const16
            010: di
            011: si
            100: ra
            101: sp
            110: pc
            111: tx

        LOAD_TL: Causes tl to load from the main bus (incompatible with W_N = 111)
        LOAD_TH: Causes th to load from the main bus (incompatible with W_N = 111)
        PC_CE#: PC count enable
        SP_U/D#: Sets the count direction of SP
        SP_CE#: SP count enable
        RA_CE#: RA count enable
        SI_U/D#: Sets the count direction of SI
        SI_CE#: SI count enable
        DI_U/D#: Sets the count direction of DI
        DI_CE#: DI count enable

GPR Unit:
    Control Lines: 9
        W_N[2]: The device to write / modify
            00: a
            01: b
            10: c
            11: d

        R_N[2]: The device to output onto the main bus
            00: a
            01: b
            10: c
            11: d

        WE_L: Write enable
        OE_L: Output enable
        LOAD_L: Causes the device specified by W_N to load
        U_DL: Sets the count direction of the device specified by W_N
        CE_L: Enables counting for the device specified by W_N 