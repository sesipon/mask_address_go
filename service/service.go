package service

import "sync"

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

	for i := 0; i <= len(buf)-len(example); i++ {
		match := true
		for j := 0; j < len(example); j++ {
			if buf[i+j] != example[j] {
				match = false
			}
		}

		if match {

			for k := i + len(example); k < len(buf) && buf[k] != ' '; k++ {
				buf[k] = '*'
			}
		}
	}

	return string(buf)
}

// RUN - запуск сервиса

func (s *Service) Run() error {

	data, err := s.prod.Produce()
	if err != nil {
		return err
	}

	// многопточность
	textChannel := make(chan string)
	resultsChannel := make(chan string)
	counter := make(chan struct{}, 10)

	for i := 0; i < 10; i++ {
		go func() {
			for text := range textChannel {
				counter <- struct{}{}

				maskedText := s.MaskAdress(text, "http://")

				resultsChannel <- maskedText

				<-counter
			}
		}()
	}

	for _, line := range data {
		textChannel <- line
	}
	close(textChannel)

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for range resultsChannel {
			}
		}()
	}
	wg.Wait()

	var maskedMessages []string
	for i := 0; i < len(data); i++ {
		maskedMessages = append(maskedMessages, <-resultsChannel)
	}

	return s.pres.Present(maskedMessages)
}
