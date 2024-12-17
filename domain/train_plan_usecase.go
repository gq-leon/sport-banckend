package domain

import "context"

type AddPlanRequest struct {
	Name      string  `json:"name"`
	Date      string  `json:"date"`
	Category  string  `bson:"category"` // strength:力量 cardio:有氧
	Sets      int32   `json:"sets"`
	Reps      int32   `json:"reps"`
	Weight    float64 `json:"weight"`
	Duration  int32   `json:"duration"`
	Distance  float64 `json:"distance"`
	Completed int32   `bson:"completed"` // 完成组数
}

type DelPlanRequest struct {
	Id string `json:"id"`
}

type UpdateProgressPlanRequest struct {
	Id   string `json:"id"`
	Type uint   `json:"type"` // 1:完成  2:取消
}

func (req *UpdateProgressPlanRequest) IsCompleted() bool {
	return req.Type == 1
}

type TodayWorkoutResponse struct {
	Date      string      `json:"date"`
	Exercises []TrainPlan `json:"exercises"`
}

type PlanResponse struct{}

type TrainPlanUseCase interface {
	Create(ctx context.Context, data TrainPlan) error
	GetPlansByDate(ctx context.Context, userID string, date string) ([]TrainPlan, error)
	Delete(ctx context.Context, ID string) error
	GetPlanByID(ctx context.Context, ID string) (TrainPlan, error)
	UpdateByID(ctx context.Context, ID string, data *TrainPlan) error
}
