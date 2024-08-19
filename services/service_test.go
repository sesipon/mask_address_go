package services

import (

	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

)


type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) Produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}


type MockPresenter struct {
	mock.Mock
}


func (m *MockPresenter) Present(data []string) error {
	args := m.Called(data)
	return args.Error(0)
}


type ServiceTestSuite struct {
	suite.Suite
	service *Service
	mockProd *MockProducer
	mockPres *MockPresenter
}


func (s *ServiceTestSuite) SetupTest() {
	s.mockProd = &MockProducer{}
	s.mockPres = &MockPresenter{}
	s.service = NewService(s.mockProd, s.mockPres, "test_output.txt")
}


func (s *ServiceTestSuite) TestRun() {
	expectedData := []string{"data1", "http://www.lenigrad.ru", "data3"}
	maskedData := []string{"data1", "http://***************", "data3"} // Ожидаемый результат

	// Задаем ожидания для моков
	s.mockProd.On("Produce").Return(expectedData, nil)
	s.mockPres.On("Present", maskedData).Return(nil)

	// Выполняем метод Run сервиса
	err := s.service.Run("test_input.txt")
	s.NoError(err)

	// Проверяем, что моки были вызваны
	s.mockProd.AssertExpectations(s.T())
	s.mockPres.AssertExpectations(s.T())
}

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
