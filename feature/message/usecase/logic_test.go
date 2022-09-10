package usecase

import (
	"testing"
	"wa/domain"
	"wa/domain/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddMessage(t *testing.T) {
	repo := new(mocks.MessageData)

	mockData := domain.Message{
		ID:      1,
		Message: "hello...",
		IDUser:  1,
	}

	EmptyData := domain.Message{
		ID:      0,
		Message: "",
		IDUser:  0,
	}
	returnData := mockData
	returnData.ID = 1

	t.Run("Success Add Message", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(mockData).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddMessage(0, mockData)

		assert.Nil(t, err)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})

	t.Run("Error insert", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(EmptyData).Once()
		useCase := New(repo, validator.New())
		res, err := useCase.AddMessage(0, EmptyData)

		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
	})
}

func TestDeleteMessage(t *testing.T) {
	repo := new(mocks.MessageData)

	mockData := domain.Message{
		ID:      1,
		Message: "hello...",
		IDUser:  1,
	}

	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(true).Once()

		usecase := New(repo, validator.New())

		res, err := usecase.DelMessage(mockData.ID)
		assert.Nil(t, err)
		assert.Equal(t, true, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(false).Once()

		usecase := New(repo, validator.New())

		res, _ := usecase.DelMessage(1)
		assert.Equal(t, false, res)
		repo.AssertExpectations(t)
	})
}
