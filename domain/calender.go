package domain

import "context"

type MonthWorkoutsRequest struct {
	Date string `json:"date"`
}

type WorkoutsResponse struct {
	Date      string     `json:"date"`
	Exercises []Exercise `json:"exercises"`
}

type Exercise struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Reps      int32   `json:"reps"`
	Weight    float64 `json:"weight"`
	Distance  float64 `json:"distance"`
	Duration  int32   `json:"duration"`
	Sets      int32   `json:"sets"`
	Completed int32   `json:"completed"`
	Type      string  `json:"type"`
}

type CalenderUseCase interface {
	GetMonthWorkouts(ctx context.Context, userID string, date string) ([]TrainPlan, error)
	GetDayWorkouts(ctx context.Context, userID string, date string) ([]TrainPlan, error)
}

type CalenderRepository interface {
	GetTrainPlansByDate(ctx context.Context, userID, startDate, endDate string) ([]TrainPlan, error)
}
