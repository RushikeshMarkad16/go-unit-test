package services

import (
	"errors"
	mocks "example/mocks/repositories"
	"example/models"
	"testing"
)

func TestInsert(t *testing.T) {
	repo := mocks.ProductRepositoryInterface{}
	service := NewProductService(&repo)

	//subtest for success
	t.Run("success", func(t *testing.T) {
		repo.Mock.On("Add", models.Product{
			ID:    "QW56H2J",
			Name:  "Mac",
			Price: 126000,
			Stock: 63,
		}).Return(nil).Once()

		err := service.Insert("QW56H2J", "Mac", 126000, 63)

		if err != nil {
			t.Errorf("Failed to insert %v ", err.Error())
		}
	})

	//subtest for failure
	t.Run("failure", func(t *testing.T) {
		repo.Mock.On("Add", models.Product{
			ID:    "QW56H2J",
			Name:  "Mac",
			Price: 126000,
			Stock: 63,
		}).Return(errors.New("Failed to add")).Once()

		err := service.Insert("QW56H2J", "Mac", 126000, 63)

		if err == nil {
			t.Errorf("Expecting error, got no eror")
		}
	})
}
