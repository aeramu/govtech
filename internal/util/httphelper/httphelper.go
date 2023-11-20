package httphelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alam/govtech/internal/api"
	"github.com/alam/govtech/internal/util/errorhelper"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

const internalServerErrorMessage = "INTERNAL_SERVER_ERROR"

func ReadBody(request *http.Request, result interface{}) error {
	var b bytes.Buffer
	_, err := io.Copy(&b, request.Body)
	if err != nil {
		return errorhelper.WrapWithCode(err, fmt.Sprintf("cannot copy request body to buffer: %+v", b.String()), http.StatusBadRequest)
	}

	err = json.NewDecoder(&b).Decode(&result)
	if err != nil && err != io.EOF {
		return errorhelper.WrapWithCode(err, fmt.Sprintf("cannot convert request body: %+v", b.String()), http.StatusBadRequest)
	}

	return nil
}

func ReadPathVarInt(request *http.Request, name string) int64 {
	str := mux.Vars(request)[name]
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func ReadQueryParamInt(request *http.Request, name string) int64 {
	str := request.URL.Query().Get(name)
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func Write(writer http.ResponseWriter, data interface{}) {
	resp, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("failed marshal http response: %s", err))
	}
	_, err = writer.Write(resp)
	if err != nil {
		panic(fmt.Sprintf("failed write http response: %s", err))
	}
}

func WriteError(writer http.ResponseWriter, err error) {
	log.Println(err)
	errCode := errorhelper.GetCode(err)
	errMsg := err.Error()
	if errCode == http.StatusInternalServerError {
		errMsg = internalServerErrorMessage
	}

	writer.WriteHeader(errorhelper.GetCode(err))

	resp, err := json.Marshal(api.ErrorResponse{
		Success: false,
		Error:   errMsg,
	})
	if err != nil {
		panic(fmt.Sprintf("failed marshal http response: %s", err))
	}
	_, err = writer.Write(resp)
	if err != nil {
		panic(fmt.Sprintf("failed write http response: %s", err))
	}
}
