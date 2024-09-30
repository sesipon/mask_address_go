package service

import (
	"mask_adress/service/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskAdress(t *testing.T) {
	service := &Service{}

	tests := []struct {
		text     string
		adress   string
		expected string
	}{
		{"Hello, http://porno залетай и наслаждайся", "http://", "Hello, http://***** залетай и наслаждайся"},
		{"Новый адрес http://www.con-con.ru а старый http://www.gon-gon.ru не работает", "http://", "Новый адрес http://************** а старый http://************** не работает"},
		{"www.lenin.ru", "www.", "www.********"},
		{"No address here", "www", "No address here"},
	}

	for _, test := range tests {
		result := service.MaskAdress(test.text, test.adress)
		assert.Equal(t, test.expected, result, "один для '%s' и '%s'", test.text, test.adress)
	}
}

// короче мок это объект сам мок ничего не тестирует он заменяет продюсер презентер
func TestRun(t *testing.T) {
	mockProducer := new(mocks.Producer)
	mockPresenter := new(mocks.Presenter)

	mockProducer.On("Produce").Return([]string{"item1", "item2", "http://Suka"}, nil)
	mockPresenter.On("Present", []string{"item1", "item2", "http://****"}).Return(nil)

	svc := NewService(mockProducer, mockPresenter)

	// Simulated svc.Run() implementation
	err := svc.Run() // Assuming this calls mockProducer.Produce() inside
	assert.NoError(t, err)

	mockProducer.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}
