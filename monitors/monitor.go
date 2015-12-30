package monitors

const (
	UN = iota
	OK
	NOK
)

type BaseMonitor struct {
	title       string
	description string
	tripped     bool
}

func NewBaseMonitor(title, description string) *BaseMonitor {
	return &BaseMonitor{title, description, false}
}

func (m *BaseMonitor) Trip() {
	m.tripped = true
}

func (m *BaseMonitor) Untrip() {
	m.tripped = false
}

func (m *BaseMonitor) IsTripped() bool {
	return m.tripped
}

type Monitor interface {
	Stater
	Describer
	Tripper
}

type Tripper interface {
	Trip()
	Untrip()
	IsTripped() bool
}

type Stater interface {
	Check() int
	Values() []ValueWithTimestamp
}

type Describer interface {
	Name() string
	Description() string
}

type Group struct {
	Name     string
	Monitors []Monitor
}

func AllFailed(m Monitor) bool {
	for i := range m.Values() {
		if m.Values()[i].Value != NOK {
			return false
		}
	}
	return true
}
