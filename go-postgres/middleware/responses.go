package middleware

type basic_response struct {
	Message string `json:"message,omitempty"`
}

type error_response struct {
	Err_type string `json:"err_type,omitempty"`
	Message  string `json:"message,omitempty"`
}

type with_id_response struct {
	ID      int64  `json:"id,string,omitempty"`
	Message string `json:"message,omitempty"`
}

type flag_response struct {
	F int `json:"f,string"`
}

type directory struct {
	Name    string     `json:"name"`
	Columns []string   `json:"columns"`
	Data    [][]string `json:"data"`
}
