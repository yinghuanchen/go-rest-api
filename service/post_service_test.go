package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-rest-api/entity"
)

type MockRepo struct {
	mock.Mock
}

func (mock *MockRepo) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepo) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "post is empty", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepo)

	post := entity.Post{ID: 1, Title: "A", Text: "B"}
	// Set up expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)
	testService := NewPostService(mockRepo)
	result, err := testService.FindAll()

	// Mock assertion
	mockRepo.AssertExpectations(t)

	// Data assertion
	assert.Nil(t, err)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}