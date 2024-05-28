package entities

type Status string

const (
	Waiting   Status = "Waiting"
	Started   Status = "Started"
	Completed Status = "Completed"
	Blocked   Status = "Blocked"
)

func (s Status) IsValid() bool {
	switch s {
	case Waiting, Started, Completed, Blocked:
		return true
	}
	return false
}
