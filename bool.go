// Package abool provides atomic Boolean type for cleaner code and
// better performance.
package abool

import "sync/atomic"

// AtomicBool is an atomic Boolean
// Its methods are all atomic, thus safe to be called by
// multiple goroutines simultaneously
// Note: When embedding into a struct, one should always use
// *AtomicBool to avoid copy
type AtomicBool int32

// New creates an AtomicBool with default to false
func New() *AtomicBool {
	return new(AtomicBool)
}

// NewBool creates an AtomicBool with given default value
func NewBool(ok bool) *AtomicBool {
	ab := New()
	if ok {
		ab.Set()
	}
	return ab
}

// Set sets the Boolean to true
func (ab *AtomicBool) Set() {
	atomic.StoreInt32((*int32)(ab), 1)
}

// UnSet sets the Boolean to false
func (ab *AtomicBool) UnSet() {
	atomic.StoreInt32((*int32)(ab), 0)
}

// IsSet returns whether the Boolean is true
func (ab *AtomicBool) IsSet() bool {
	return intToBool(atomic.LoadInt32((*int32)(ab)))
}

// SetTo sets the boolean with given Boolean.
func (ab *AtomicBool) SetTo(yes bool) {
	atomic.StoreInt32((*int32)(ab), boolToInt(yes))
}

// Toggle negates boolean atomically and returns the previous boolean value.
func (ab *AtomicBool) Toggle() bool {
	for {
		old := atomic.LoadInt32((*int32)(ab))
		if atomic.CompareAndSwapInt32((*int32)(ab), old, toggleInt(old)) {
			return old == 1
		}
	}
}

// SetToIf sets the Boolean to new only if the Boolean matches the old
// Returns whether the set was done
func (ab *AtomicBool) SetToIf(old, new bool) (set bool) {
	o := boolToInt(old)
	n := boolToInt(new)
	return atomic.CompareAndSwapInt32((*int32)(ab), o, n)
}

// boolToInt convert a boolean to int32
func boolToInt(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// intToBool convert a int32 to boolean
func intToBool(i int32) bool {
	return i == 1
}

// toggleInt toggles the int32
func toggleInt(i int32) int32 {
	if i == 1 {
		return 0
	}
	return 1
}
