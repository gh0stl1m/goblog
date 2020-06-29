package delivery_test

import (
	"net/http/httptest"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	delivery "github.com/gh0stl1m/goblog/accountservice/delivery/http"
	"github.com/gh0stl1m/goblog/accountservice/domains/fixtures"
	"github.com/gh0stl1m/goblog/accountservice/domains/mocks"
)

var _ = Describe("Handlers", func() {
	Describe("AccountHandlers", func() {
		var mockAccountUseCases mocks.MockAccountUseCases
		BeforeEach(func() {
			mockAccountUseCases = mocks.MockAccountUseCases{}
		})

		Context("Given a valid ID", func() {
			It("When ID exists then it must be returned", func() {
				accountID := strconv.Itoa(fixtures.GenerateID())
				// Arrange
				mockAccountUseCases.On("GetByID", accountID).
					Return(fixtures.GenerateAccount(accountID), nil)
				req := httptest.NewRequest("GET", "/accounts/"+accountID, nil)
				resp := httptest.NewRecorder()

				// Act
				delivery.NewRouter().ServeHTTP(resp, req)

				// Asserts
				Expect(resp.Code).To(Equal(200))
			})
		})
	})

})
