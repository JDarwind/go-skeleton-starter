package httpkit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	w           http.ResponseWriter
	r           *http.Request
	status      int
	contentType string
}

type responsePayload struct {
	Result bool `json:"result"`
	Data   any  `json:"data,omitempty"`
	Errors any  `json:"errors,omitempty"`
}

func NewResponse(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		w:           w,
		r:           r,
		status:      http.StatusOK,
		contentType: "application/json",
	}
}

func (res *Response) ContentType(ct string) *Response {
	res.contentType = ct
	return res
}

func (res *Response) Status(code int) *Response {
	res.status = code
	return res
}

func (res *Response) Header(key, value string) *Response {
	res.w.Header().Set(key, value)
	return res
}

func (res *Response) RemoveHeader(key string) *Response {
	res.w.Header().Del(key)
	return res
}

func (res *Response) Success(data any) {
	res.w.Header().Set("Content-Type", res.contentType)

	payload := responsePayload{
		Result: true,
		Data:   normalizePayload(data),
	}

	res.w.WriteHeader(res.status)

	if res.contentType == "application/json" {
		_ = json.NewEncoder(res.w).Encode(payload)
		return
	}

	writeRaw(res.w, payload.Data)
}

func (res *Response) Error(err any, status ...int) {
	code := res.status
	if len(status) > 0 {
		code = status[0]
	}

	res.w.Header().Set("Content-Type", res.contentType)

	payload := responsePayload{
		Result: false,
		Errors: normalizePayload(err),
	}

	res.w.WriteHeader(code)

	if res.contentType == "application/json" {
		_ = json.NewEncoder(res.w).Encode(payload)
		return
	}

	writeRaw(res.w, payload.Errors)
}

func writeRaw(w http.ResponseWriter, v any) {
	switch x := v.(type) {
	case []byte:
		_, _ = w.Write(x)
	case string:
		_, _ = w.Write([]byte(x))
	default:
		fmt.Fprint(w, v)
	}
}

func normalizePayload(v any) any {
	switch v.(type) {
	case string, []byte, map[string]any, map[string][]string, []any:
		return v
	default:
		return v
	}
}
