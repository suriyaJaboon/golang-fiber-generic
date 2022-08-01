package handles

import (
	"encoding/json"
	"errors"
	"fg/dtos"
	"fg/x"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	NOTFOUND     = x.NOTFOUND
	VALIDATION   = x.VALIDATION
	UNAUTHORIZED = x.UNAUTHORIZED
)

type Context struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
}

type (
	ErrorValidator struct {
		Code           int
		Opt            string
		ErrorResponses []*ErrorValidation
	}

	ErrorValidation struct {
		Tag     string `json:"tag"`
		Field   string `json:"field"`
		Message string `json:"message"`
	}
)

func (e *ErrorValidator) Error() string {
	bt, err := json.Marshal(e.ErrorResponses)
	if err != nil {
		return err.Error()
	}

	return e.Opt + " " + strconv.Itoa(e.Code) + ": " + string(bt)
}

func ErrValidator(errs []*ErrorValidation) error {
	return &ErrorValidator{Code: fiber.StatusBadRequest, Opt: VALIDATION, ErrorResponses: errs}
}

func ErrBodyParser() error {
	return &dtos.Error{Code: fiber.StatusNotAcceptable, Opt: VALIDATION, Err: errors.New("BodyParser")}
}

var (
	Fc = fiber.Config{
		ServerHeader: "0.0.1",
		BodyLimit:    10 * 1024 * 1024, // 10 MB
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  10 * time.Second,
		ErrorHandler: fiberErrorHandler,
	}

	//fiberFilterHandler = func(c *fiber.Ctx) bool {
	//	var routes = []string{pkg.BasePath, pkg.Health, pkg.Login}
	//	for _, route := range routes {
	//		if route == filepath.Base(c.Path()) {
	//			return true
	//		}
	//		if filepath.Base(c.Path()) == pkg.Product && c.Method() == fiber.MethodPost {
	//			return true
	//		}
	//	}
	//	return false
	//}

	fiberErrorHandler = func(c *fiber.Ctx, err error) error {
		log.Println(err)

		var code = fiber.StatusInternalServerError
		var res = &Context{
			Code:    "SERVER-ERROR",
			Message: fiber.ErrInternalServerError.Message,
		}

		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
			res.Message = e.Message
		}

		if e, ok := err.(*dtos.Error); ok {
			if e.Code > 0 {
				code = e.Code
				if e.Opt == NOTFOUND {
					code = fiber.StatusNotFound
				}
			} else {
				if e.Opt == UNAUTHORIZED {
					code = fiber.StatusUnauthorized
				}
			}

			res.Code = e.Opt
			res.Message = e.Err.Error()

			if e.Err == mongo.ErrNoDocuments {
				res.Message = fiber.ErrInternalServerError.Message
			}
			if e.Opt == UNAUTHORIZED {
				res.Message = fiber.ErrUnauthorized.Message
			}
		}

		if e, ok := err.(*ErrorValidator); ok {
			code = e.Code
			res.Code = e.Opt
			res.Message = e.ErrorResponses
		}

		return c.Status(code).JSON(res)
	}
)
