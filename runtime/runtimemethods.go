package main

import (
	"runtime"
	"runtime/debug"
)

func main() {

	// Gosched yields the processor, allowing other goroutines to run. It does not
	// suspend the current goroutine, so execution resumes automatically.
	runtime.Gosched()

	// GC runs a garbage collection and blocks the caller until the
	// garbage collection is complete. It may also block the entire
	// program.
	runtime.GC()

	// A negative percentage effectively disables garbage collection, unless
	// the memory limit is reached.
	// See SetMemoryLimit for more details
	debug.SetGCPercent(-1)

	// задает **общее** количество памяти, которое Go runtime может использовать.
	// GOMEMLIMIT или debug.SetMemoryLimit из пакета runtime
	// Когда общее значение памяти приближается к GOMEMLIMIT - запускается сборщик мусора
	//
	//**Спираль смерти**
	//
	//По мере стремления "живой" памяти кучи к GOMEMLIMIT, GC будет запускаться все чаще и чаще.
	//
	//Если не будет никаких ограничений, сборщик мусора Go в будет работать непрерывно.
	//
	//**Мягкое управление памятью**
	//
	//- \-  Go не предоставляет 100% гарантий соблюдения ограничения памяти GOMEMLIMIT
	//- \-  Предел использования процессорного времени: 50% с окном CPU в 2 * GOMAXPROCS секунды
	debug.SetMemoryLimit(111)
}
