//go:build !riscv64

package ptrace

import "golang.org/x/sys/unix"

const (
	ptraceSysemu           = unix.PTRACE_SYSEMU
	ptraceSysemuSinglestep = unix.PTRACE_SYSEMU_SINGLESTEP
)
