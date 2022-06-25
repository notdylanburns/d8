/*
	Package containing structures and methods relating to the D8 Cpu
*/
package cpu

import "github.com/notdylanburns/d8/emu/pkg/common"

// D8Cpu holds all the internal state of a D8 CPU
type D8Cpu struct {
	pc 		uint16
	sp 		uint16
	li 		uint16
	si, di 	uint16
	tx 		uint16

	flags   uint8
	mode 	uint8

	cycle 	uint8

	read 	common.D8ReadFunc
	write 	common.D8WriteFunc
}

func New(read common.D8ReadFunc, write common.D8WriteFunc) *D8Cpu {
	cpu := new(D8Cpu);
	cpu.read = read
	cpu.write = write

	return cpu
}

func (cpu *D8Cpu) Cycle() {
	cpu.cycle = (cpu.cycle + 1) % 5

	if cpu.cycle == 4 {
		instr := cpu.read(cpu.pc)
		
		cpu.exec(instr)
		
		cpu.cycle = 0
	}

}

func (cpu *D8Cpu) exec(instr uint8) {

}