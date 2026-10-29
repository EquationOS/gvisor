//go:build riscv64

package ptrace

// golang.org/x/sys/unix does not define PTRACE_SYSEMU* on riscv64, but the
// Linux UAPI values are architecture-independent and the generic kernel
// ptrace implementation supports them.
const (
	ptraceSysemu           = 31 // PTRACE_SYSEMU
	ptraceSysemuSinglestep = 32 // PTRACE_SYSEMU_SINGLESTEP
)
