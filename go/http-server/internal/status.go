package internal

type Status int

const (
	StatusOk                  Status = 200
	StatusCreated             Status = 201
	StatusNotFound            Status = 404
	StatusInternalServerError Status = 500
)

func (s Status) ToString() string {
	switch s {
	case StatusOk:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusNotFound:
		return "Not Found"
	case StatusInternalServerError:
		return "Internal Server Error"
	default:
		return ""
	}
}
