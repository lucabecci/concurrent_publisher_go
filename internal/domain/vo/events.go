package vo

type Event string

const (
	Synchronize Event = "Synchronize"
	Restore     Event = "Restore"
)

func (e Event) IsValid() bool {
	switch e {
	case Synchronize, Restore:
		return true
	}
	return false
}
