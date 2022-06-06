// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
	externalRef0 "github.com/gojek/xp/common/api/schema"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// BadRequest defines model for BadRequest.
type BadRequest externalRef0.Error

// CreateExperimentSuccess defines model for CreateExperimentSuccess.
type CreateExperimentSuccess struct {
	Data externalRef0.Experiment `json:"data"`
}

// CreateProjectSettingsSuccess defines model for CreateProjectSettingsSuccess.
type CreateProjectSettingsSuccess struct {
	Data externalRef0.ProjectSettings `json:"data"`
}

// GetExperimentHistorySuccess defines model for GetExperimentHistorySuccess.
type GetExperimentHistorySuccess struct {
	Data externalRef0.ExperimentHistory `json:"data"`
}

// GetExperimentSuccess defines model for GetExperimentSuccess.
type GetExperimentSuccess struct {
	Data externalRef0.Experiment `json:"data"`
}

// GetProjectExperimentVariablesSuccess defines model for GetProjectExperimentVariablesSuccess.
type GetProjectExperimentVariablesSuccess struct {
	Data []string `json:"data"`
}

// GetProjectSettingsSuccess defines model for GetProjectSettingsSuccess.
type GetProjectSettingsSuccess struct {
	Data externalRef0.ProjectSettings `json:"data"`
}

// GetSegmentersSuccess defines model for GetSegmentersSuccess.
type GetSegmentersSuccess struct {
	Data []externalRef0.Segmenter `json:"data"`
}

// InternalServerError defines model for InternalServerError.
type InternalServerError externalRef0.Error

// ListExperimentHistorySuccess defines model for ListExperimentHistorySuccess.
type ListExperimentHistorySuccess struct {
	Data   []externalRef0.ExperimentHistory `json:"data"`
	Paging *externalRef0.Paging             `json:"paging,omitempty"`
}

// ListExperimentsSuccess defines model for ListExperimentsSuccess.
type ListExperimentsSuccess struct {
	Data   []externalRef0.Experiment `json:"data"`
	Paging *externalRef0.Paging      `json:"paging,omitempty"`
}

// ListProjectsSuccess defines model for ListProjectsSuccess.
type ListProjectsSuccess struct {
	Data []externalRef0.Project `json:"data"`
}

// ListSegmentersSuccess defines model for ListSegmentersSuccess.
type ListSegmentersSuccess struct {
	Data []externalRef0.Segmenter `json:"data"`
}

// NotFound defines model for NotFound.
type NotFound externalRef0.Error

// UpdateExperimentSuccess defines model for UpdateExperimentSuccess.
type UpdateExperimentSuccess struct {
	Data externalRef0.Experiment `json:"data"`
}

// UpdateProjectSettingsSuccess defines model for UpdateProjectSettingsSuccess.
type UpdateProjectSettingsSuccess struct {
	Data externalRef0.ProjectSettings `json:"data"`
}

// CreateExperimentRequestBody defines model for CreateExperimentRequestBody.
type CreateExperimentRequestBody struct {
	Description *string                            `json:"description"`
	EndTime     time.Time                          `json:"end_time"`
	Interval    *int32                             `json:"interval"`
	Name        string                             `json:"name"`
	Segment     externalRef0.ExperimentSegment     `json:"segment"`
	StartTime   time.Time                          `json:"start_time"`
	Status      externalRef0.ExperimentStatus      `json:"status"`
	Tier        *externalRef0.ExperimentTier       `json:"tier,omitempty"`
	Treatments  []externalRef0.ExperimentTreatment `json:"treatments"`
	Type        externalRef0.ExperimentType        `json:"type"`
	UpdatedBy   *string                            `json:"updated_by,omitempty"`
}

// CreateProjectSettingsRequestBody defines model for CreateProjectSettingsRequestBody.
type CreateProjectSettingsRequestBody struct {
	EnableS2idClustering *bool                          `json:"enable_s2id_clustering,omitempty"`
	RandomizationKey     string                         `json:"randomization_key"`
	Segmenters           externalRef0.ProjectSegmenters `json:"segmenters"`

	// Object containing information to define a valid treatment schema
	TreatmentSchema *externalRef0.TreatmentSchema `json:"treatment_schema,omitempty"`
	ValidationUrl   *string                       `json:"validation_url,omitempty"`
}

// UpdateExperimentRequestBody defines model for UpdateExperimentRequestBody.
type UpdateExperimentRequestBody struct {
	Description *string                            `json:"description"`
	EndTime     time.Time                          `json:"end_time"`
	Interval    *int32                             `json:"interval"`
	Segment     externalRef0.ExperimentSegment     `json:"segment"`
	StartTime   time.Time                          `json:"start_time"`
	Status      externalRef0.ExperimentStatus      `json:"status"`
	Tier        *externalRef0.ExperimentTier       `json:"tier,omitempty"`
	Treatments  []externalRef0.ExperimentTreatment `json:"treatments"`
	Type        externalRef0.ExperimentType        `json:"type"`
	UpdatedBy   *string                            `json:"updated_by,omitempty"`
}

// UpdateProjectSettingsRequestBody defines model for UpdateProjectSettingsRequestBody.
type UpdateProjectSettingsRequestBody struct {
	EnableS2idClustering *bool                          `json:"enable_s2id_clustering,omitempty"`
	RandomizationKey     string                         `json:"randomization_key"`
	Segmenters           externalRef0.ProjectSegmenters `json:"segmenters"`

	// Object containing information to define a valid treatment schema
	TreatmentSchema *externalRef0.TreatmentSchema `json:"treatment_schema,omitempty"`
	ValidationUrl   *string                       `json:"validation_url,omitempty"`
}

// ListExperimentsParams defines parameters for ListExperiments.
type ListExperimentsParams struct {
	Status *externalRef0.ExperimentStatus `json:"status,omitempty"`

	// Used together with the start_time, to filter experiments that are at least partially running in the input range.
	EndTime   *time.Time                   `json:"end_time,omitempty"`
	Tier      *externalRef0.ExperimentTier `json:"tier,omitempty"`
	Type      *externalRef0.ExperimentType `json:"type,omitempty"`
	Name      *string                      `json:"name,omitempty"`
	UpdatedBy *string                      `json:"updated_by,omitempty"`

	// Search experiment name and description for a partial match of the search text
	Search *string `json:"search,omitempty"`

	// Result page number. It defaults to 1.
	Page *int32 `json:"page,omitempty"`

	// Number of items on each page. It defaults to 10.
	PageSize *int32 `json:"page_size,omitempty"`

	// Used together with the end_time, to filter experiments that are at least partially running in the input range.
	StartTime *time.Time              `json:"start_time,omitempty"`
	Segment   *map[string]interface{} `json:"segment,omitempty"`

	// controls whether or not weak segmenter matches (experiments where the segmenter is optional) should be returned
	IncludeWeakMatch *bool `json:"include_weak_match,omitempty"`
}

// ListExperimentHistoryParams defines parameters for ListExperimentHistory.
type ListExperimentHistoryParams struct {

	// Result page number. It defaults to 1.
	Page *int32 `json:"page,omitempty"`

	// Number of items on each page. It defaults to 10.
	PageSize *int32 `json:"page_size,omitempty"`
}

// CreateExperimentJSONRequestBody defines body for CreateExperiment for application/json ContentType.
type CreateExperimentJSONRequestBody CreateExperimentRequestBody

// UpdateExperimentJSONRequestBody defines body for UpdateExperiment for application/json ContentType.
type UpdateExperimentJSONRequestBody UpdateExperimentRequestBody

// CreateProjectSettingsJSONRequestBody defines body for CreateProjectSettings for application/json ContentType.
type CreateProjectSettingsJSONRequestBody CreateProjectSettingsRequestBody

// UpdateProjectSettingsJSONRequestBody defines body for UpdateProjectSettings for application/json ContentType.
type UpdateProjectSettingsJSONRequestBody UpdateProjectSettingsRequestBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List info of all projects set up for Experimentation
	// (GET /projects)
	ListProjects(w http.ResponseWriter, r *http.Request)
	// Get all parameters required for generating treatments for the given project
	// (GET /projects/{project_id}/experiment-variables)
	GetProjectExperimentVariables(w http.ResponseWriter, r *http.Request, projectId int64)
	// Get experiments for a project w.r.t. query params
	// (GET /projects/{project_id}/experiments)
	ListExperiments(w http.ResponseWriter, r *http.Request, projectId int64, params ListExperimentsParams)
	// Create a new experiment for a project
	// (POST /projects/{project_id}/experiments)
	CreateExperiment(w http.ResponseWriter, r *http.Request, projectId int64)
	// Get details of an experiment with the given experiment_id and project_id
	// (GET /projects/{project_id}/experiments/{experiment_id})
	GetExperiment(w http.ResponseWriter, r *http.Request, projectId int64, experimentId int64)
	// Update an experiment with the given experiment_id and project_id
	// (PUT /projects/{project_id}/experiments/{experiment_id})
	UpdateExperiment(w http.ResponseWriter, r *http.Request, projectId int64, experimentId int64)
	// Disable an experiment with the given experiment_id and project_id
	// (PUT /projects/{project_id}/experiments/{experiment_id}/disable)
	DisableExperiment(w http.ResponseWriter, r *http.Request, projectId int64, experimentId int64)
	// Enable an experiment with the given experiment_id and project_id
	// (PUT /projects/{project_id}/experiments/{experiment_id}/enable)
	EnableExperiment(w http.ResponseWriter, r *http.Request, projectId int64, experimentId int64)
	// List an experiment's historical versions
	// (GET /projects/{project_id}/experiments/{experiment_id}/history)
	ListExperimentHistory(w http.ResponseWriter, r *http.Request, projectId int64, experimentId int64, params ListExperimentHistoryParams)
	// List an experiment's historical versions
	// (GET /projects/{project_id}/experiments/{experiment_id}/history/{version})
	GetExperimentHistory(w http.ResponseWriter, r *http.Request, projectId int64, experimentId int64, version int64)
	// Get all segmenter configurations required for generating experiments for the given project
	// (GET /projects/{project_id}/segmenters)
	GetSegmenters(w http.ResponseWriter, r *http.Request, projectId int64)
	// Get the settings for the given project
	// (GET /projects/{project_id}/settings)
	GetProjectSettings(w http.ResponseWriter, r *http.Request, projectId int64)
	// Set up new project for Experimentation
	// (POST /projects/{project_id}/settings)
	CreateProjectSettings(w http.ResponseWriter, r *http.Request, projectId int64)
	// Update the settings for the given project
	// (PUT /projects/{project_id}/settings)
	UpdateProjectSettings(w http.ResponseWriter, r *http.Request, projectId int64)
	// List all segmenter configurations registered with XP
	// (GET /segmenters)
	ListSegmenters(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ListProjects operation middleware
func (siw *ServerInterfaceWrapper) ListProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListProjects(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetProjectExperimentVariables operation middleware
func (siw *ServerInterfaceWrapper) GetProjectExperimentVariables(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProjectExperimentVariables(w, r, projectId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListExperiments operation middleware
func (siw *ServerInterfaceWrapper) ListExperiments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListExperimentsParams

	// ------------- Optional query parameter "status" -------------
	if paramValue := r.URL.Query().Get("status"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "status", r.URL.Query(), &params.Status)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter status: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time" -------------
	if paramValue := r.URL.Query().Get("end_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "end_time", r.URL.Query(), &params.EndTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter end_time: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "tier" -------------
	if paramValue := r.URL.Query().Get("tier"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "tier", r.URL.Query(), &params.Tier)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter tier: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "type" -------------
	if paramValue := r.URL.Query().Get("type"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "type", r.URL.Query(), &params.Type)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter type: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "name" -------------
	if paramValue := r.URL.Query().Get("name"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "name", r.URL.Query(), &params.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "updated_by" -------------
	if paramValue := r.URL.Query().Get("updated_by"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "updated_by", r.URL.Query(), &params.UpdatedBy)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter updated_by: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "search" -------------
	if paramValue := r.URL.Query().Get("search"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "search", r.URL.Query(), &params.Search)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter search: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page_size: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_time" -------------
	if paramValue := r.URL.Query().Get("start_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "start_time", r.URL.Query(), &params.StartTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter start_time: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "segment" -------------
	if paramValue := r.URL.Query().Get("segment"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "segment", r.URL.Query(), &params.Segment)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter segment: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "include_weak_match" -------------
	if paramValue := r.URL.Query().Get("include_weak_match"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "include_weak_match", r.URL.Query(), &params.IncludeWeakMatch)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter include_weak_match: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListExperiments(w, r, projectId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateExperiment operation middleware
func (siw *ServerInterfaceWrapper) CreateExperiment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateExperiment(w, r, projectId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetExperiment operation middleware
func (siw *ServerInterfaceWrapper) GetExperiment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "experiment_id" -------------
	var experimentId int64

	err = runtime.BindStyledParameter("simple", false, "experiment_id", chi.URLParam(r, "experiment_id"), &experimentId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter experiment_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetExperiment(w, r, projectId, experimentId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateExperiment operation middleware
func (siw *ServerInterfaceWrapper) UpdateExperiment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "experiment_id" -------------
	var experimentId int64

	err = runtime.BindStyledParameter("simple", false, "experiment_id", chi.URLParam(r, "experiment_id"), &experimentId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter experiment_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateExperiment(w, r, projectId, experimentId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DisableExperiment operation middleware
func (siw *ServerInterfaceWrapper) DisableExperiment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "experiment_id" -------------
	var experimentId int64

	err = runtime.BindStyledParameter("simple", false, "experiment_id", chi.URLParam(r, "experiment_id"), &experimentId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter experiment_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DisableExperiment(w, r, projectId, experimentId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// EnableExperiment operation middleware
func (siw *ServerInterfaceWrapper) EnableExperiment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "experiment_id" -------------
	var experimentId int64

	err = runtime.BindStyledParameter("simple", false, "experiment_id", chi.URLParam(r, "experiment_id"), &experimentId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter experiment_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.EnableExperiment(w, r, projectId, experimentId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListExperimentHistory operation middleware
func (siw *ServerInterfaceWrapper) ListExperimentHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "experiment_id" -------------
	var experimentId int64

	err = runtime.BindStyledParameter("simple", false, "experiment_id", chi.URLParam(r, "experiment_id"), &experimentId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter experiment_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ListExperimentHistoryParams

	// ------------- Optional query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page_size: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListExperimentHistory(w, r, projectId, experimentId, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetExperimentHistory operation middleware
func (siw *ServerInterfaceWrapper) GetExperimentHistory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "experiment_id" -------------
	var experimentId int64

	err = runtime.BindStyledParameter("simple", false, "experiment_id", chi.URLParam(r, "experiment_id"), &experimentId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter experiment_id: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "version" -------------
	var version int64

	err = runtime.BindStyledParameter("simple", false, "version", chi.URLParam(r, "version"), &version)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter version: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetExperimentHistory(w, r, projectId, experimentId, version)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetSegmenters operation middleware
func (siw *ServerInterfaceWrapper) GetSegmenters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSegmenters(w, r, projectId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetProjectSettings operation middleware
func (siw *ServerInterfaceWrapper) GetProjectSettings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetProjectSettings(w, r, projectId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateProjectSettings operation middleware
func (siw *ServerInterfaceWrapper) CreateProjectSettings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateProjectSettings(w, r, projectId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// UpdateProjectSettings operation middleware
func (siw *ServerInterfaceWrapper) UpdateProjectSettings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "project_id" -------------
	var projectId int64

	err = runtime.BindStyledParameter("simple", false, "project_id", chi.URLParam(r, "project_id"), &projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter project_id: %s", err), http.StatusBadRequest)
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UpdateProjectSettings(w, r, projectId)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// ListSegmenters operation middleware
func (siw *ServerInterfaceWrapper) ListSegmenters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{""})

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ListSegmenters(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects", wrapper.ListProjects)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects/{project_id}/experiment-variables", wrapper.GetProjectExperimentVariables)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects/{project_id}/experiments", wrapper.ListExperiments)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/projects/{project_id}/experiments", wrapper.CreateExperiment)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects/{project_id}/experiments/{experiment_id}", wrapper.GetExperiment)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/projects/{project_id}/experiments/{experiment_id}", wrapper.UpdateExperiment)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/projects/{project_id}/experiments/{experiment_id}/disable", wrapper.DisableExperiment)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/projects/{project_id}/experiments/{experiment_id}/enable", wrapper.EnableExperiment)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects/{project_id}/experiments/{experiment_id}/history", wrapper.ListExperimentHistory)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects/{project_id}/experiments/{experiment_id}/history/{version}", wrapper.GetExperimentHistory)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects/{project_id}/segmenters", wrapper.GetSegmenters)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/projects/{project_id}/settings", wrapper.GetProjectSettings)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/projects/{project_id}/settings", wrapper.CreateProjectSettings)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/projects/{project_id}/settings", wrapper.UpdateProjectSettings)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/segmenters", wrapper.ListSegmenters)
	})

	return r
}