//go:build riscv64

package systrap

// golang.org/x/sys/unix does not define PTRACE_SYSEMU on riscv64, but the
// Linux UAPI value is architecture-independent and the generic kernel
// ptrace implementation supports it.
const ptraceSysemu = 31 // PTRACE_SYSEMU
