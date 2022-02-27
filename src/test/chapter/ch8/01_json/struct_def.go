package json

type BasicInfo struct {
	Name string `01_json:"name"`
	Age  int    `01_json:"age"`
}
type JobInfo struct {
	Skills []string `01_json:"skills"`
}
type Employee struct {
	BasicInfo BasicInfo `01_json:"basic_info"`
	JobInfo   JobInfo   `01_json:"job_info"`
}
