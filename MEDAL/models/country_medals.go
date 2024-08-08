package models

type CountryMedals struct {
	Name        string `json:"name"`
	GoldCount   int    `json:"gold_count"`
	SilverCount int    `json:"silver_count"`
	BronzeCount int    `json:"bronze_count"`
}
