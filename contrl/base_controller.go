package contrl

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/lunelabs/go-gen-rest/resp"
	"io/ioutil"
	"net"
	"net/http"
)

type BaseController struct {
}

func (c *BaseController) GetRequestObject(w http.ResponseWriter, r *http.Request, model interface{}) error {
	content, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(content))

	if err := json.Unmarshal(content, &model); err != nil {
		c.WriteErrorResponse(w, err.Error(), resp.ErrorCodeBadRequest, http.StatusBadRequest)

		return err
	}

	val := validator.New()

	if err := val.Struct(model); err != nil {
		c.WriteErrorResponse(w, err.Error(), resp.ErrorCodeBadRequest, http.StatusBadRequest)

		return err
	}

	return nil
}

func (c *BaseController) GetJsonKeys(r *http.Request) ([]string, error) {
	dynamic := make(map[string]interface{})
	content, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(content))

	if err := json.Unmarshal(content, &dynamic); err != nil {
		return []string{}, err
	}

	keys := make([]string, 0, len(dynamic))

	for k := range dynamic {
		keys = append(keys, k)
	}

	return keys, nil
}

func (c *BaseController) GetRequestFilter(w http.ResponseWriter, r *http.Request, model interface{}) error {
	simpleMap := map[string]interface{}{}

	for k := range r.URL.Query() {
		simpleMap[k] = r.URL.Query().Get(k)
	}

	encodedJson, err := json.Marshal(simpleMap)

	if err != nil {
		c.WriteErrorResponse(w, err.Error(), resp.ErrorCodeBadRequest, http.StatusBadRequest)

		return err
	}

	if err := json.Unmarshal(encodedJson, &model); err != nil {
		c.WriteErrorResponse(w, err.Error(), resp.ErrorCodeBadRequest, http.StatusBadRequest)

		return err
	}

	val := validator.New()

	if err := val.Struct(model); err != nil {
		c.WriteErrorResponse(w, err.Error(), resp.ErrorCodeBadRequest, http.StatusBadRequest)

		return err
	}

	return nil
}

func (c *BaseController) WriteErrorResponse(
	w http.ResponseWriter,
	errorMessage string,
	errorCode string,
	httpCode int,
) {
	resp.WriteErrorResponse(
		w,
		errorMessage,
		errorCode,
		httpCode,
	)
}

func (c *BaseController) WriteJsonResponse(w http.ResponseWriter, r interface{}) {
	resp.WriteJsonResponse(w, r)
}

func (c *BaseController) WriteJsonResponseWithCode(w http.ResponseWriter, r interface{}, code int) {
	resp.WriteJsonResponse(w, r)
}

func (c *BaseController) GetRequestIp(r *http.Request) string {
	ipAddress := r.Header.Get("X-Real-Ip")

	if len(ipAddress) == 0 {
		ipAddress = r.RemoteAddr
	}

	if ip, _, err := net.SplitHostPort(ipAddress); err == nil {
		return ip
	}

	return ipAddress
}
