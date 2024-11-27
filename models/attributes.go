package models

import "time"

type Attribute struct {
    ID        int       `json:"id"`                       // Unique identifier
    Name      string    `json:"name"`                     // Attribute name
    DataType  string    `json:"data_type"`                // Data type of the attribute
    CreatedAt time.Time `json:"created_at"`               // Timestamp of creation
    UpdatedAt time.Time `json:"updated_at"`               // Timestamp of last update
}


type CreateAttribute struct {
	Name      string    `json:"name"`                     // Attribute name
    DataType  string    `json:"data_type"`                // Data type of the attribute
}

type UpdateAttribute struct {
	ID            int       `json:"id"`
	Name      string    `json:"name"`                     // Attribute name
    DataType  string    `json:"data_type"`                // Data type of the attribute
}

type GetListAttributeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListAttributeResponse struct {
	Count int     `json:"count"`
	Attributes []*Attribute `json:"attributes"`
}
