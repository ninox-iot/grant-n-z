package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/usecase/service"
)

type PolicyHandlerImpl struct {
	RequestHandler    RequestHandler
	PolicyService service.PolicyService
}

func NewPolicyHandlerHandler() PolicyHandler {
	log.Logger.Info("Inject `RequestHandler`, `PolicyService` to `PolicyHandler`")
	return PolicyHandlerImpl{
		RequestHandler: NewRequestHandler(),
		PolicyService: service.NewPolicyService(),
	}
}

func (ph PolicyHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		ph.Get(w, r)
	case http.MethodPost:
		ph.Post(w, r)
	case http.MethodPut:
		ph.Put(w, r)
	case http.MethodDelete:
		ph.Delete(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (ph PolicyHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("GET policies")
	id := r.URL.Query().Get(entity.PolicyId.String())

	roleMemberEntities, err := ph.PolicyService.Get(id)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(roleMemberEntities)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func (ph PolicyHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	log.Logger.Info("POST policies")
	var policyEntity *entity.Policy

	body, err := ph.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	_ = json.Unmarshal(body, &policyEntity)
	if err := ph.RequestHandler.ValidateHttpRequest(w, policyEntity); err != nil {
		return
	}

	policyEntity, err = ph.PolicyService.InsertPolicy(policyEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(policyEntity)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func (ph PolicyHandlerImpl) Put(w http.ResponseWriter, r *http.Request) {
}

func (ph PolicyHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
