package models

type CountryRanking struct {
	Name        string `json:"name"`
	TotalMedals int    `json:"total_medals"`
	GoldCount   int    `json:"gold_count"`
	SilverCount int    `json:"silver_count"`
	BronzeCount int    `json:"bronze_count"`
	Rank        int    `json:"rank"`
}
