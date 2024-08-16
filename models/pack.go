package models

type Pack struct {
	ID               int    `json:"id"`
	Brand            string `json:"brand"`
	Model            string `json:"model"`
	YearPurchased    int    `json:"year_purchased"`
	CapacityLitres   int    `json:"capacity_litres"`
	WeightEmptyGrams int    `json:"weight_empty_grams"`
	UsageCount       int    `json:"usage_count"`
}

type PackInstance struct {
	ID               int               `json:"id"`
	PackID           int               `json:"pack_id"`
	TotalWeightGrams int               `json:"total_weight_grams"`
	UsageDate        string            `json:"usage_date"`
	Sections         []SectionInstance `json:"sections"`
	UserID           int               `json:"user_id"`
	Notes            string            `json:"notes"`
}
