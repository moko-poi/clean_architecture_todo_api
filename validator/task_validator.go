package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/moko-poi/go-todo-api/model"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task,
		validation.Field(
			// タイトルに対するバリデーション
			&task.Title,
			// 存在するか
			validation.Required.Error("title is required"),
			// 文字数の数
			validation.RuneLength(1, 10).Error("limited max 10 char"),
		),
	)
}
