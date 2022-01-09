package domain

import (
	"github.com/castillofranciscodaniel/golang-beers/err"
	mockDomain "github.com/castillofranciscodaniel/golang-beers/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
)

var mockBeerRepository *mockDomain.MockBeerRepository
var beerService DefaultBeerService

func setUp(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockBeerRepository = mockDomain.NewMockBeerRepository(ctrl)
	beerService = DefaultBeerService{beersRepository: mockBeerRepository}
	return func() {
		defer ctrl.Finish()
	}

}

func Test_Post_should_return_errors_if_the_new_beer_can_not_be_created(t *testing.T) {
	tearDown := setUp(t)
	defer tearDown()

	beer, _ := NewBeer(2, "Patagonia", "Norte", "Chile", 740, "CL")

	mockBeerRepository.EXPECT().Post(beer).Return(nil, err.NotFoundError)

	beerService.Post(beer)

	t.Error("Failed while testing the status code")

}
