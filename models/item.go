package models

type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	WeightGrams int    `json:"weight_grams"`
	UsageCount  int    `json:"usage_count"`
}

type ItemInstance struct {
	ID                int    `json:"id"`
	ItemID            int    `json:"item_id"`
	SectionInstanceID int    `json:"section_instance_id"`
	Notes             string `json:"notes"`
}
