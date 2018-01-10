package contract

import "fmt"

// Ensures .
type Ensures struct {
	Ok      bool
	Message string
	ensures []*ensure
}

type ensure struct {
	f       func() bool
	message string
}

// Ensure .
func Ensure(f func() bool, message string) *Ensures {
	return &Ensures{
		Ok: true,
		ensures: []*ensure{
			&ensure{
				f:       f,
				message: message,
			},
		},
	}
}

// Ensure .
func (ensures *Ensures) Ensure(f func() bool, message string) *Ensures {
	ensures.ensures = append(ensures.ensures, &ensure{
		f:       f,
		message: message,
	})

	return ensures
}

// Do process ensures
func (ensures *Ensures) Do(err *error) {
	for _, ensure := range ensures.ensures {
		if !ensure.f() {
			err2 := fmt.Errorf(ensure.message)
			err = &err2
			break
		}
	}
}
