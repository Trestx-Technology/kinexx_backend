package goalDB

import "kinexx_backend/pkg/entity"

type GoalService interface {
	MakeGoal(goal *entity.GoalDB) (string, error)
	GetGoal(goalID string) (entity.GoalDB, error)
	GetGoalCreatedByMe(userID string) ([]entity.GoalDB, error)
	DeleteGoal(goalID string) error
	PauseGoal(status, goalID string) (string, error)
	EditGoal(goal entity.GoalDB, goalID string) (string, error)
}
