package service

type Service struct {
	prod Producer
	pres Presenter
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.0 --name Producer
type Producer interface {
	Produce() ([]string, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.45.0 --name Presenter
type Presenter interface {
	Present([]string) error
}

func NewService(prod Producer, pres Presenter) *Service {
	return &Service{prod: prod, pres: pres}
}

func (s *Service) MaskAdress(text string, adress string) string {
	buf := []byte(text)
	example := []byte(adress)

	// Ищем вхождения example
	for i := 0; i <= len(buf)-len(example); i++ {
		match := true
		for j := 0; j < len(example); j++ {
			if buf[i+j] != example[j] {
				match = false
			}
		}

		if match {
			// Маскируем символы после example
			for k := i + len(example); k < len(buf) && buf[k] != ' '; k++ {
				buf[k] = '*'
			}
		}
	}

	return string(buf)
}

// RUN - запуск сервиса

func (s *Service) Run() error {
	// получаем данные из Producer
	data, err := s.prod.Produce()
	if err != nil {
		return err
	}

	var maskedMessages []string
	for _, line := range data {
		maskedMessages = append(maskedMessages, s.MaskAdress(line, "http://"))
	}

	return s.pres.Present(maskedMessages)
}
