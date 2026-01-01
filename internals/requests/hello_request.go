package requests

type HelloRequest struct {
	Name  string `json:"name" query:"name" validate:"omitempty,min=3"`
}
