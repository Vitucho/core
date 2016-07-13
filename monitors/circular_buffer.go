package monitors

import "time"

type ValueWithTimestamp struct {
	Value     int
	Timestamp int64
}

type CircularBuffer struct {
	values   []ValueWithTimestamp
	next     int
	capacity int
}

func NewCircularBuffer(capacity int) CircularBuffer {
	sl := make([]ValueWithTimestamp, 0, capacity)
	return CircularBuffer{sl, 0, capacity}
}

func (buf *CircularBuffer) All(value int) bool {
	ok := true
	for _, val := range buf.values {
		ok = ok && (val.Value == value)
	}
	return ok
}

func (buf *CircularBuffer) Append(val int) {

	// hit the last value, start from zero.
	if buf.next == cap(buf.values) {
		buf.next = 0
	}

	// extend length if necesary.
	if buf.next == len(buf.values) {
		extended := make([]ValueWithTimestamp, len(buf.values)+1, buf.capacity)
		for i := range buf.values {
			extended[i] = buf.values[i]
		}
		buf.values = extended
	}

	buf.values[buf.next] = ValueWithTimestamp{val, time.Now().Unix()}
	buf.next++
}

func (buf *CircularBuffer) GetValues() []ValueWithTimestamp {
	ret := make([]ValueWithTimestamp, len(buf.values), cap(buf.values))
	copy(ret, buf.values)
	return ret
}
