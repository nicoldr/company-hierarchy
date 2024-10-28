package models

type Department struct {
    ID   int    `json:"id"`	
    Name string `json:"name"`		// Department name
    Flags int    `json:"flags"`		// Codified states in bits (active, deleted, approved)
    ParentID *int `json:"parent_id"`	// Parent department ID, can be null
}
