package models

type Section struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UsageCount  int    `json:"usage_count"`
}

type SectionInstance struct {
	ID                      int            `json:"id"`
	SectionID               int            `json:"section_id"`
	PackInstanceID          int            `json:"pack_instance_id"`
	TotalWeightGrams        int            `json:"total_weight_grams"`
	PercentageOfTotalWeight float64        `json:"percentage_of_total_weight"`
	Items                   []ItemInstance `json:"items"`
	Notes                   string         `json:"notes"`
}
