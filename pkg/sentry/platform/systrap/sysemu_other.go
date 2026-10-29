//go:build !riscv64

package systrap

import "golang.org/x/sys/unix"

const ptraceSysemu = unix.PTRACE_SYSEMU
