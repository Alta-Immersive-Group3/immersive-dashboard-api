package service

import (
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/feedback"
	"github.com/ALTA-Immersive-Group3/immersive-dahsboard-api/features/mentee"
	"github.com/go-playground/validator/v10"
)

type feedbackService struct {
	feedbackData feedback.FeedbackDataInterface
	validate     *validator.Validate
}

func New(repo feedback.FeedbackDataInterface) feedback.FeedbackServiceInterface {
	return &feedbackService{
		feedbackData: repo,
		validate:     validator.New(),
	}
}

func (service *feedbackService) Create(input feedback.Core) (feedback.Core, error) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return feedback.Core{}, errValidate
	}

	id, errInsert := service.feedbackData.Insert(input)
	if errInsert != nil {
		return feedback.Core{}, errInsert
	}

	return service.GetById(id)
}

func (service *feedbackService) GetAll() ([]feedback.Core, error) {
	data, err := service.feedbackData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (service *feedbackService) GetAllByMenteeId(idMentee uint64) ([]feedback.Core, mentee.Core, error) {
	dataFeedbacks, dataMentee, err := service.feedbackData.SelectAllByMenteeId(idMentee)
	if err != nil {
		return nil, mentee.Core{}, err
	}
	return dataFeedbacks, dataMentee, err
}

func (service *feedbackService) GetById(id uint64) (feedback.Core, error) {
	data, err := service.feedbackData.SelectById(id)
	if err != nil {
		return feedback.Core{}, err
	}
	return data, err
}

func (service *feedbackService) UpdateById(id uint64, input feedback.Core) (feedback.Core, error) {
	errUpdate := service.feedbackData.UpdateById(id, input)
	if errUpdate != nil {
		return feedback.Core{}, errUpdate
	}

	return service.GetById(id)
}

func (service *feedbackService) DeleteById(id uint64) error {
	errUpdate := service.feedbackData.DeleteById(id)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
