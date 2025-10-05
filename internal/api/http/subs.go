package http

import (
	"net/http"
	"subs-service/internal/api/http/response"
	"subs-service/internal/api/http/types"
	"subs-service/internal/config"
	"subs-service/internal/usecases"
	"subs-service/pkg/http/handlers"

	"github.com/go-chi/chi/v5"
)

type SubHandler struct {
	subSvc  usecases.SubService
	pathCfg config.PathConfig
	svcCfg  config.ServiceConfig
	dataCfg config.DataConfig
}

func NewSubHandler(
	subSvc usecases.SubService,
	pathCfg config.PathConfig,
	svcCfg config.ServiceConfig,
	dataCfg config.DataConfig,
) *SubHandler {
	return &SubHandler{
		subSvc:  subSvc,
		pathCfg: pathCfg,
		svcCfg:  svcCfg,
		dataCfg: dataCfg,
	}
}

func (h *SubHandler) WithSubHandlers() handlers.RouterOption {
	return func(r chi.Router) {
		r.Get(h.pathCfg.GetSub, h.getSubHandler)
		r.Post(h.pathCfg.PostSub, h.postSubHandler)
		r.Put(h.pathCfg.PutSub, h.putSubHandler)
		r.Delete(h.pathCfg.DeleteSub, h.deleteSubHandler)

		r.Get(h.pathCfg.ListSubs, h.listSubsHandler)
		r.Get(h.pathCfg.GetSummary, h.getSummaryHandler)
	}
}

// @Summary 	Get subscription by id
// @Tags 		subs
// @Produce 	json
// @Param 		id 		path 	string true "Subcription's id"
// @Success 	200 {object} 	domain.Sub "Successfully got sub"
// @Failure 	400 {string} 	string "Bad request"
// @Failure 	404 {string} 	string "Object not found"
// @Failure 	500 {string} 	string "Internal error"
// @Router		/subs/{id} 		[get]
func (h *SubHandler) getSubHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetSubRequest(r)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.subSvc.GetSub(r.Context(), req.Id)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, res, http.StatusOK)
}

// @Summary 	Create new subscription
// @Description Для параметров подписки по умолчанию установлены следующие ограничения:
// @Description - имя сервиса должно быть непустым и не длиннее 50 символов;
// @Description - стоимость подписки должна быть положительной, но не более 100.000.
// @Tags 		subs
// @Accept  	json
// @Produce 	json
// @Param 		sub 	body 	domain.Sub true "Sub details"
// @Success 	201 {object} 	domain.Sub "Successfully created sub"
// @Failure 	400 {string} 	string "Bad request"
// @Failure 	500 {string} 	string "Internal error"
// @Router		/subs 			[post]
func (h *SubHandler) postSubHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePostSubRequest(r, h.dataCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.subSvc.PostSub(r.Context(), &req.Sub)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, res, http.StatusCreated)
}

// @Summary 	Update subscription's data by id
// @Tags 		subs
// @Accept 		json
// @Produce 	json
// @Param 		id 				path 	string true "Sub's id"
// @Param 		sub 			body 	domain.Sub true "Sub details"
// @Success 	200 {object} 			domain.Sub "Successfully updated sub"
// @Failure 	400 {string} 			string "Bad request"
// @Failure 	404 {string} 			string "Object not found"
// @Failure 	500 {string} 			string "Internal error"
// @Router		/subs/{id}				[put]
func (h *SubHandler) putSubHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreatePutSubRequest(r, h.dataCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.subSvc.PutSub(r.Context(), req.Id, &req.Sub)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, res, http.StatusOK)
}

// @Summary 	Delete subscription by id
// @Tags 		subs
// @Produce 	json
// @Param 		id 				path 	string true "Sub's id"
// @Success 	200 {object} 			domain.Sub "Successfully deleted sub"
// @Failure 	400 {string} 			string "Bad request"
// @Failure 	404 {string} 			string "Object not found"
// @Failure 	500 {string} 			string "Internal error"
// @Router		/subs/{id}				[delete]
func (h *SubHandler) deleteSubHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateDeleteSubRequest(r)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.subSvc.DeleteSub(r.Context(), req.Id)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, res, http.StatusOK)
}

// @Summary 	Get user's subscriptions list
// @Description Параметр user_id обязателен для получения списка подписок. Опционально поддерживается фильтрация по названию сервиса.
// @Description Также поддерживается keyset пагинация - опционально можно указать размер страницы (по умолчанию 20) и токен для получения следующей страницы (поле next_page_token в теле предыдущего запроса).
// @Tags 		subs
// @Produce 	json
// @Param 		user_id 		query 	string true "User's id"
// @Param 		service_name 	query 	string false "Service name"
// @Param 		page_size 		query 	int false "Page size"
// @Param 		page_token 		query 	int false "Page token (for keyset pagination)"
// @Success 	200 {object} 			types.ListSubsResponse "Successfully got subs list"
// @Failure 	400 {string} 			string "Bad request"
// @Failure 	500 {string} 			string "Internal error"
// @Router		/subs					[get]
func (h *SubHandler) listSubsHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateListSubsRequest(r, h.dataCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.subSvc.ListSubs(r.Context(), req.Opts)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, types.CreateListSubsResponse(res), http.StatusOK)
}

// @Summary 	Get summary of user's subscriptions (e.g. total price)
// @Description Параметр user_id обязателен для получения суммарной стоимости подписок. Опционально поддерживается фильтрация по названию сервиса.
// @Tags 		subs
// @Produce 	json
// @Param 		user_id 		query 	string true "User's id"
// @Param 		service_name 	query 	string false "Service name"
// @Success 	200 {object} 			domain.Summary "Successfully got summary"
// @Failure 	400 {string} 			string "Bad request"
// @Failure 	500 {string} 			string "Internal error"
// @Router 		/subs/summary 			[get]
func (h *SubHandler) getSummaryHandler(w http.ResponseWriter, r *http.Request) {
	req, err := types.CreateGetSummaryRequest(r, h.dataCfg)
	if err != nil {
		response.ProcessCreatingRequestError(w, err, h.svcCfg.DebugMode)
		return
	}

	res, err := h.subSvc.GetSummary(r.Context(), req.Opts)
	if err != nil {
		response.ProcessError(w, err, h.svcCfg.DebugMode)
		return
	}

	response.WriteResponse(w, res, http.StatusOK)
}
