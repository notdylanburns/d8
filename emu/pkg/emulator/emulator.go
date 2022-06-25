// Contains structures and methods relating to the emulator
package emulator

import (
	"github.com/notdylanburns/d8/emu/internal/logger"
	"github.com/notdylanburns/d8/emu/pkg/cpu"
	"github.com/notdylanburns/d8/emu/pkg/device"
)

type D8Emulator struct {
	cpu 	*cpu.D8Cpu
	devices	[]*device.D8Device
}

func (emu *D8Emulator)Read(addr uint16) uint8 {
	var dev *device.D8Device = nil
	for _, d := range emu.devices {
		if (*d).ReadEnabled(addr) {
			if dev != nil {
				logger.Errorf("Bus conflict while reading at address 0x%04x (%v and %v)", addr, dev, d);
			} else {
				dev = d
			}
		}
	}

	return (*dev).Read(addr)
}

func (emu *D8Emulator)Write(addr uint16, data uint8) {
	var dev *device.D8Device = nil
	for _, d := range emu.devices {
		if (*d).WriteEnabled(addr) {
			if dev != nil {
				// can be suppressed
				logger.Errorf("Bus conflict while writing at address 0x%04x (%v and %v)", addr, dev, d);
			} else {
				dev = d
			}
		}
	}

	(*dev).Write(addr, data)
}

func New() *D8Emulator {
	emu := new(D8Emulator)
	emu.cpu = cpu.New(
		func(addr uint16) uint8 {
			return emu.Read(addr)
		},
		func(addr uint16, data uint8) {
			emu.Write(addr, data)
		},
	)
}