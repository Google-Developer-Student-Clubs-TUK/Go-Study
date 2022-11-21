package tukorea

type (
	login struct {
		Id  string `json:"id"`
		Pwd string `json:"pwd"`
	}

	tukoreaLogin struct {
		usr_id  string
		usr_pwd string
	}
)

type (
	timetable struct {
		subject string
	}
)

type ApiResponse[T any] struct {
	resultCode int
	msg        string
	result     T
}
