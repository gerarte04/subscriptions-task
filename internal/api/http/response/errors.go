package response

import (
	"errors"
	"log"
	"net/http"
	"subs-service/internal/repository"
	pkgErrors "subs-service/pkg/errors"
)

var (
	ErrInternal = errors.New("internal server error")

	errorCodes = map[error]int{
		repository.ErrInvalidSubData: http.StatusBadRequest,
		repository.ErrNoSubIDExists: http.StatusNotFound,
	}
)

func ProcessCreatingRequestError(w http.ResponseWriter, err error, debugMode bool) {
	log.Print(err.Error())

	if !debugMode {
		err = pkgErrors.UnwrapAll(err)
	}

	http.Error(w, err.Error(), http.StatusBadRequest)
}

func ProcessError(w http.ResponseWriter, err error, debugMode bool) {
	log.Print(err.Error())

	if !debugMode {
		err = pkgErrors.UnwrapAll(err)
	}

	code := http.StatusInternalServerError

	if docCode, ok := errorCodes[pkgErrors.UnwrapAll(err)]; ok {
		code = docCode
	} else if !debugMode {
		err = ErrInternal
	}

	http.Error(w, err.Error(), code)
}
