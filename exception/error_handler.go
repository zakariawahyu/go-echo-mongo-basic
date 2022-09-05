package exception

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/zakariawahyu/go-echo-mongo-basic/response"
	"net/http"
)

type httpErrorHandler struct {
	statusCode map[error]int
}

func NewHttpErrorHandler(errorStatusCode map[error]int) *httpErrorHandler {
	return &httpErrorHandler{
		statusCode: errorStatusCode,
	}
}

func (self *httpErrorHandler) getStatusCode(err error) int {
	for key, value := range self.statusCode {
		if errors.Is(err, key) {
			return value
		}
	}

	return http.StatusInternalServerError
}

func unwarpRecursive(err error) error {
	var originalErr = err

	for originalErr != nil {
		var internalErr = errors.Unwrap(originalErr)

		if internalErr == nil {
			break
		}
		originalErr = internalErr
	}

	return originalErr
}

func (self *httpErrorHandler) Handler(err error, ctx echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    self.getStatusCode(err),
			Message: unwarpRecursive(err).Error(),
		}
	}

	code := he.Code
	message := he.Message
	if _, ok := he.Message.(string); ok {
		message = response.WebResponse{
			Code:    code,
			Status:  "error",
			Results: err.Error(),
		}
	}

	//Send Response
	if !ctx.Response().Committed {
		if ctx.Request().Method == http.MethodHead {
			err = ctx.NoContent(he.Code)
		} else {
			err = ctx.JSON(code, message)
		}
		if err != nil {
			ctx.Echo().Logger.Error(err)
		}
	}
}
