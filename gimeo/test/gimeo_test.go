package gimeo_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/julianedialkova/gimeo/gimeo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Gimeo", func() {

	var client *Client
	var request *http.Request
	var username = "username"
	var secret = "secret"
	var token = "token"

	Context("When calling BuildRequestURL", func() {

		It("should build a proper uri when the path begins with /", func() {
			result := BuildRequestURL("/hello")
			Ω(result).To(Equal(DefaultRequest.Hostname + "/hello"))
		})

		It("should build a proper uri when the path doesn't begin with /", func() {
			result := BuildRequestURL("hello")
			Ω(result).To(Equal(DefaultRequest.Hostname + "/hello"))
		})
	})

	Context("When calling ApplyDefaults", func() {

		BeforeEach(func() {
			request, _ = http.NewRequest("GET", DefaultRequest.Hostname, nil)
		})

		Context("and the user has a token", func() {

			var token = "token"

			BeforeEach(func() {
				client = Vimeo(username, secret, token)
			})

			It("should set the correct default values", func() {
				client.ApplyDefaults(request)
				expectedToken := fmt.Sprintf("bearer %s", token)
				Ω(request.Header.Get("Authorization")).To(Equal(expectedToken))
				for key, value := range *DefaultRequest.Headers {
					Ω(request.Header.Get(key)).To(Equal(value))
				}
			})
		})

		Context("and the user has no token", func() {

			BeforeEach(func() {
				client = Vimeo(username, secret, "")
			})

			It("should set the correct default values", func() {
				client.ApplyDefaults(request)

				auth := []byte(fmt.Sprintf("%s:%s", username, secret))
				token := base64.StdEncoding.EncodeToString(auth)
				expectedToken := fmt.Sprintf("basic %s", token)

				Ω(request.Header.Get("Authorization")).To(Equal(expectedToken))
				for key, value := range *DefaultRequest.Headers {
					Ω(request.Header.Get(key)).To(Equal(value))
				}
			})
		})
	})

	Context("When making a Get request", func() {

		var result string
		var statusCode int
		server := ghttp.NewServer()
		DefaultRequest.Hostname = server.URL()

		Context("when the request succeeds", func() {
			BeforeEach(func() {
				client = Vimeo(username, secret, token)
				result = "{ response: correct }"
				statusCode = http.StatusOK
				// result := []byte("Correct!")

				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/test"),
						// ghttp.VerifyHeader(http.Header{
						//     "X-Sprocket-API-Version": []string{"1.0"},
						// }),
						ghttp.RespondWithJSONEncoded(statusCode, result),
					),
				)
			})

			It("should return the correct response", func() {
				resp, err := client.Get("/test", nil)
				Ω(err).Should(BeNil())
				// Ω(resp).Should(Equal(result)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(body))
				fmt.Println("{ response: correct }")
				Ω(body).Should(Equal([]byte(result)))

			})
		})
	})
})
