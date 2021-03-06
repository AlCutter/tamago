// ARM processor support
// https://github.com/f-secure-foundry/tamago
//
// Copyright (c) F-Secure Corporation
// https://foundry.f-secure.com
//
// Use of this source code is governed by the license
// that can be found in the LICENSE file.

package arm

// defined in irq.s
func irq_enable()
func irq_disable()

// InterruptsEnable enables IRQ and FIQ interrupts.
func (cpu *CPU) InterruptsEnable() {
	irq_enable()
}

// InterruptsDisable disables IRQ and FIQ interrupts.
func (cpu *CPU) InterruptsDisable() {
	irq_disable()
}
