package models

import "gorm.io/gorm"

// StoredNetworkMetric is gorm.Model that represents a table of networks metrics
type StoredNetworkMetric struct {
	gorm.Model
	ID                string `gorm:"primaryKey;unique;not null"`
	NetworkMetricJSON string `gorm:"not null"`
}
