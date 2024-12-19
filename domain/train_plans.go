package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTrainPlan = "train_plans" // 计划

	TrainPlanFormat   = "2006-01-02"
	TrainPlanStrength = "strength"
	TrainPlanCardio   = "cardio"
)

var TrainPlanCategory = map[string]string{
	TrainPlanStrength: "力量训练",
	TrainPlanCardio:   "有氧运动",
}

type TrainPlanRepository interface {
	Create(ctx context.Context, plan *TrainPlan) error
	Update(ctx context.Context, id string, plan *TrainPlan) error
	Delete(ctx context.Context, id string) error
	GetByDate(ctx context.Context, userID string, date string) ([]TrainPlan, error)
	GetByID(ctx context.Context, ID string) (TrainPlan, error)
}

type TrainPlan struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Date      string             `bson:"date" json:"date"` // YYYY-mm-dd
	Name      string             `bson:"name" json:"name"`
	Category  string             `bson:"category" json:"category"`   // strength:力量 cardio:有氧
	Reps      int32              `bson:"reps" json:"reps"`           // 次数
	Weight    float64            `bson:"weight" json:"weight"`       // 重量，kg
	Distance  float64            `bson:"distance" json:"distance"`   // 距离，km
	Duration  int32              `bson:"duration" json:"duration"`   // 时间，m
	Sets      int32              `bson:"sets" json:"sets"`           // 组数
	Completed int32              `bson:"completed" json:"completed"` // 完成组数
}

func (t *TrainPlan) GetID() string {
	return t.ID.Hex()
}

func (t *TrainPlan) GetType() string {
	if category, ok := TrainPlanCategory[t.Category]; ok {
		return category
	}
	return "未知训练"
}
