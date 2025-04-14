package model

import (
	"github.com/google/uuid"
)

type MachineProduct struct {
	ID            uuid.UUID            `json:"id" bson:"-"`
	MongoId       string               `bson:"_id" json:"-"`
	Name          string               `json:"name" bson:"name"`
	Manufacturer  string               `json:"manufacturer" bson:"manufacturer"`
	ShortDesc     string               `json:"shortDescription" bson:"short_description"`
	Description   string               `json:"description" bson:"description"`
	Price         int                  `json:"price" bson:"price"`
	Attributes    []Attribute          `json:"attributes" bson:"attributes"`
	AttributesMap map[string]Attribute `json:"attributesMap" bson:"attributes_map"`
	ImageSet      []Image              `json:"imageSet" bson:"image_set"`
	Quantity      int                  `json:"quantity" bson:"quantity"`
	Category      Category             `json:"category" bson:"category"`
	SubCategory   Category             `json:"subCategory" bson:"sub_category"`
	IsDigital     bool                 `json:"isDigital" bson:"is_digital"`
}

type Attribute struct {
	Name           string      `json:"name" bson:"name"`
	Category       string      `json:"category" bson:"category"`
	Type           int         `json:"type" bson:"type"` // 1=string, 2=number, 3=bool
	Order          int         `json:"order" bson:"order"`
	Requirement    int         `json:"requirement" bson:"requirement"`
	IsTranslatable bool        `json:"isTranslatable" bson:"is_translatable"`
	Value          interface{} `json:"value" bson:"value"`
	ValueI18n      interface{} `json:"value_i18n" bson:"value_i18n"`
}

type Image struct {
	ID            string `json:"id" bson:"_id"`
	Title         string `json:"title" bson:"title"`
	FileName      string `json:"fileName" bson:"file_name"`
	URL           string `json:"url" bson:"url"`
	ImageWidth    int    `json:"imageWidth" bson:"image_width"`
	ImageHeight   int    `json:"imageHeight" bson:"image_height"`
	ImageType     int    `json:"imageType" bson:"image_type"`
	TargetDevices []int  `json:"targetDevices" bson:"target_devices"`
}

type Category struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Priority    int    `json:"priority" bson:"priority"`
	Color       string `json:"color" bson:"color"`
	CreatedAt   string `json:"createdAt" bson:"created_at"`
	UpdatedAt   string `json:"updatedAt" bson:"updated_at"`
	Type        int    `json:"type" bson:"type"`
}
