package models

type Shipment struct {
	Mawb        string `csv:"MAWB,omitempty" validate:"required"`
	Hawb        string `csv:"HAWB,omitempty" validate:"required"`
	Origin      string `csv:"Origin,omitempty" validate:"required"`
	Destination string `csv:"Destination,omitempty" validate:"required"`
	Content     string `csv:"Content,omitempty" validate:"required"`
	Pieces      int    `csv:"Pieces,omitempty" validate:"required"`
	Weight      int    `csv:"Weight in kg,omitempty" validate:"required"`
	Consigner   string `csv:"Consigner,omitempty" validate:"required"`
	Consignee   string `csv:"Consignee,omitempty" validate:"required"`
}
