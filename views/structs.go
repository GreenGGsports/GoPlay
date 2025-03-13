// views/structs.go
package views

type Response struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}

type PostRequest struct {
	Name string `"json:" Name`
	Todo string `"json": todo`
}
