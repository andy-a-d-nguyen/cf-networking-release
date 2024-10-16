package handlers

import (
	"net/http"

	"code.cloudfoundry.org/lager/v3"
	"code.cloudfoundry.org/policy-server/api"
	"code.cloudfoundry.org/policy-server/store"
)

//counterfeiter:generate -o fakes/policy_cleaner.go --fake-name PolicyCleaner . policyCleaner
type policyCleaner interface {
	DeleteStalePolicies() ([]store.Policy, error)
}

//counterfeiter:generate -o fakes/error_response.go --fake-name ErrorResponse . errorResponse
type errorResponse interface {
	InternalServerError(lager.Logger, http.ResponseWriter, error, string)
	BadRequest(lager.Logger, http.ResponseWriter, error, string)
	NotFound(lager.Logger, http.ResponseWriter, error, string)
	NotAcceptable(lager.Logger, http.ResponseWriter, error, string)
	Forbidden(lager.Logger, http.ResponseWriter, error, string)
	Unauthorized(lager.Logger, http.ResponseWriter, error, string)
}

type PoliciesCleanup struct {
	PolicyMapper  api.PolicyMapper
	PolicyCleaner policyCleaner
	ErrorResponse errorResponse
}

func NewPoliciesCleanup(writer api.PolicyMapper, policyCleaner policyCleaner, errorResponse errorResponse) *PoliciesCleanup {
	return &PoliciesCleanup{
		PolicyMapper:  writer,
		PolicyCleaner: policyCleaner,
		ErrorResponse: errorResponse,
	}
}

func (h *PoliciesCleanup) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logger := getLogger(req)
	logger = logger.Session("cleanup-policies")

	c2cPolicies, err := h.PolicyCleaner.DeleteStalePolicies()
	if err != nil {
		h.ErrorResponse.InternalServerError(logger, w, err, "policies cleanup failed")
		return
	}

	for i := range c2cPolicies {
		c2cPolicies[i].Source.Tag = ""
		c2cPolicies[i].Destination.Tag = ""
	}

	bytes, err := h.PolicyMapper.AsBytes(c2cPolicies)
	if err != nil {
		h.ErrorResponse.InternalServerError(logger, w, err, "map policy as bytes failed")
		return
	}

	w.WriteHeader(http.StatusOK)
	// #nosec G104 - ignore errors writing http responses to avoid spamming logs during a DoS
	w.Write(bytes)
}
