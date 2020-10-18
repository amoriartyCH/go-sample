package models

type Product struct {
	ID 				string 		`json:"id,omitempty`
	name			string		`json:"product_name`
	description		string		`json:"product_description`
	price			float32		`json:"product_price`
}
