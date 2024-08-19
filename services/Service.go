package services

type Service struct {
	prod Producer
	pres Presenter
	out  string // Путь к результату
}

func NewService(prod Producer, pres Presenter, out string) *Service {
	return &Service{prod: prod, pres: pres, out: out}
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
				break
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
func (s *Service) Run(inputPath string) error {
	// получаем данные из Producera
	data, err := s.prod.Produce()
	if err != nil {
		return err
	}
	for i, line := range data {
		data[i] = s.MaskAdress(line, "http://")
	}

	err = s.pres.Present(data)
	if err != nil {
		return err
	}

	return nil

}
