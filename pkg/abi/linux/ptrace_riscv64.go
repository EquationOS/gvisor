// Copyright 2019 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build riscv64
// +build riscv64

package linux

// PtraceRegs mirrors struct user_regs_struct from
// arch/riscv/include/uapi/asm/ptrace.h: pc first, then x1..x31 (no x0).
// Regs[i] holds x(i) for i >= 1; Regs[0] is the program counter.
//
// +marshal
// +stateify savable
type PtraceRegs struct {
	Regs [32]uint64
}

// InstructionPointer returns the address of the next instruction to be
// executed.
func (p *PtraceRegs) InstructionPointer() uint64 {
	return p.Regs[0]
}

// StackPointer returns the address of the Stack pointer.
func (p *PtraceRegs) StackPointer() uint64 {
	return p.Regs[2]
}

// SetStackPointer sets the stack pointer to the specified value.
func (p *PtraceRegs) SetStackPointer(sp uint64) {
	p.Regs[2] = sp
}
