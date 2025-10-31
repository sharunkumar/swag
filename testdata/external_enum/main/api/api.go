package api
package api

import (
	"github.com/swaggo/swag/testdata/external_enum/external"
)

// Task represents a task with external enum fields
type Task struct {
	ID       int              `json:"id"`
	Name     string           `json:"name"`
	Status   external.Status  `json:"status"`
	Priority external.Priority `json:"priority"`
}

// GetTask example
// @Summary Get a task with external enum fields
// @Description Returns a task that uses enums from external package
// @ID get_task
// @Accept  json
// @Produce  json
// @Success 200 {object} Task "Successfully retrieved task"
// @Router /task [get]
func GetTask() Task {
	return Task{}
}
