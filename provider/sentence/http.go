package sentence

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/avast/retry-go"
	"github.com/valyala/fasthttp"
)

func httpGetDoTimeout(url string, timeout time.Duration) (body []byte, err error) {
	err = retry.Do(
		func() error {
			req := &fasthttp.Request{}
			req.SetRequestURI(url)
			req.Header.SetMethod(http.MethodGet)
			req.Header.SetContentType("application/json; charset=utf-8")

			client := &fasthttp.Client{
				TLSConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			}

			resp := &fasthttp.Response{}
			if err = client.DoTimeout(req, resp, timeout); err != nil {
				return err
			}

			if resp.StatusCode() != http.StatusOK {
				return fmt.Errorf("response status code is %d", resp.StatusCode())
			}

			body = resp.Body()

			return nil
		})
	return
}
