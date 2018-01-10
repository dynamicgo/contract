package contract

// Status .
type Status struct {
	Ok      bool
	Message string
}

func (status *Status) Error() string {
	return status.Message
}

// Require .
func (status *Status) Require(ok bool, message string) *Status {
	if !status.Ok {
		return status
	}

	if !ok {
		status.Ok = ok
		status.Message = message
	}

	return status
}

// Require start method input check
func Require(ok bool, message string) *Status {
	status := &Status{
		Ok: ok,
	}

	if !status.Ok {
		status.Message = message
	}

	return status
}
