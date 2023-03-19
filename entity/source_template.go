package entity

type SourceTemplate struct {
	Type   string                 `json:"type"`
	Fields map[string]interface{} `json:"fields"`
}
