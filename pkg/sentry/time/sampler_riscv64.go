// Copyright 2018 The gVisor Authors.
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

package time

import (
	"encoding/binary"
	"os"
)

// defaultTimebaseFreq is the RISC-V timebase used when the device tree is
// not available (e.g. inside the sandbox stub's mount namespace). qemu's
// virt machine and the RISC-V privileged architecture recommended value
// are both 10MHz.
const defaultTimebaseFreq = 10000000

// getCNTFRQ returns the RISC-V timebase frequency.
//
// The device-tree property is a big-endian 32-bit value. When the file is
// missing or unreadable, fall back to the standard 10MHz timebase rather
// than 0 (a zero frequency zeroed defaultOverheadCycles and maxOverheadCycles,
// which made lowOverheadSample loop forever).
func getCNTFRQ() uint32 {
	buf, err := os.ReadFile("/proc/device-tree/cpus/timebase-frequency")
	if err != nil || len(buf) < 4 {
		return defaultTimebaseFreq
	}
	freq := binary.BigEndian.Uint32(buf[:4])
	if freq == 0 {
		return defaultTimebaseFreq
	}
	return freq
}

// getDefaultArchOverheadCycles returns a default syscall overhead estimate
// scaled by the timebase frequency, mirroring the arm64 calibration.
func getDefaultArchOverheadCycles() TSCValue {
	cntfrq := uint64(getCNTFRQ())
	overheadCycles := 1000 * cntfrq / 1000000000
	if overheadCycles == 0 {
		overheadCycles = 1
	}
	return TSCValue(overheadCycles)
}

// defaultOverheadCycles is the default estimated syscall overhead in cycles.
// It is further refined as syscalls are made.
var defaultOverheadCycles = getDefaultArchOverheadCycles()

// maxOverheadCycles is the maximum allowed syscall overhead in cycles.
var maxOverheadCycles = 100 * defaultOverheadCycles
