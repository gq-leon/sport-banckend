package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionAttendance = "attendances"
)

type AttendanceRepository interface {
	Create(ctx context.Context, attendance *Attendance) error
	GetRecordByUserID(ctx context.Context, userID string) ([]Attendance, error)
}

type Attendance struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	Date     string             `bson:"date" json:"date"` // 2024年12月11日 星期三
	Time     string             `bson:"time" json:"time"`
	Type     string             `bson:"type" json:"type"`
	Location string             `bson:"location" json:"location"`
}
