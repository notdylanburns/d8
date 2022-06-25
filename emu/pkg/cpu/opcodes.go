package cpu

// nop
func _00(cpu *D8Cpu) {

}

// lnk ##
func _01(cpu *D8Cpu) {
	cpu.read()
}

var OPCODES [256]func(*D8Cpu) = [256]func(*D8Cpu){
	_00, _01, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
	_00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00, _00,
}