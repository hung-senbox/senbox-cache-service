package qr

import (
	"time"
)

type LibMode struct {
	ID                  string            `bson:"id" json:"id"`
	StudentId           string            `bson:"student_id" json:"student_id"`
	StudentSbtId        string            `bson:"student_sbt_id" json:"student_sbt_id"`
	TermId              string            `bson:"term_id" json:"term_id"`
	TermWeek            uint              `bson:"term_week" json:"term_week"`
	SbtCode             string            `bson:"sbt_code" json:"sbt_code"`
	ScanLocationBoxId   string            `bson:"scan_location_box_id" json:"scan_location_box_id"`
	ScanLocationBoxType string            `bson:"scan_location_box_type" json:"scan_location_box_type"`
	ScanStatus          LibModeScanStatus `bson:"scan_status" json:"scan_status"`
	ScannedBy           string            `bson:"scanned_by" json:"scanned_by"`
	CreatedAt           time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt           time.Time         `bson:"updated_at" json:"updated_at"`
}

type LibModeScanStatus struct {
	SbtStatus             string `bson:"sbt_status" json:"sbt_status"`
	SbtItemStatus         string `bson:"sbt_item_status" json:"sbt_item_status"`
	ScanLocationBoxStatus string `bson:"scan_location_box_status" json:"scan_location_box_status"`
}
