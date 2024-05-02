// Copyright 2021 OpenSSF Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: clients/repo_client.go

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/ossf/scorecard/v4/clients"
)

// MockRepoClient is a mock of RepoClient interface.
type MockRepoClient struct {
	ctrl     *gomock.Controller
	recorder *MockRepoClientMockRecorder
}

// MockRepoClientMockRecorder is the mock recorder for MockRepoClient.
type MockRepoClientMockRecorder struct {
	mock *MockRepoClient
}

// NewMockRepoClient creates a new mock instance.
func NewMockRepoClient(ctrl *gomock.Controller) *MockRepoClient {
	mock := &MockRepoClient{ctrl: ctrl}
	mock.recorder = &MockRepoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoClient) EXPECT() *MockRepoClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRepoClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRepoClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepoClient)(nil).Close))
}

// GetBranch mocks base method.
func (m *MockRepoClient) GetBranch(branch string) (*clients.BranchRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch", branch)
	ret0, _ := ret[0].(*clients.BranchRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockRepoClientMockRecorder) GetBranch(branch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockRepoClient)(nil).GetBranch), branch)
}

// GetCreatedAt mocks base method.
func (m *MockRepoClient) GetCreatedAt() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreatedAt")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreatedAt indicates an expected call of GetCreatedAt.
func (mr *MockRepoClientMockRecorder) GetCreatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatedAt", reflect.TypeOf((*MockRepoClient)(nil).GetCreatedAt))
}

// GetDefaultBranch mocks base method.
func (m *MockRepoClient) GetDefaultBranch() (*clients.BranchRef, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultBranch")
	ret0, _ := ret[0].(*clients.BranchRef)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultBranch indicates an expected call of GetDefaultBranch.
func (mr *MockRepoClientMockRecorder) GetDefaultBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultBranch", reflect.TypeOf((*MockRepoClient)(nil).GetDefaultBranch))
}

// GetDefaultBranchName mocks base method.
func (m *MockRepoClient) GetDefaultBranchName() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultBranchName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultBranchName indicates an expected call of GetDefaultBranchName.
func (mr *MockRepoClientMockRecorder) GetDefaultBranchName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultBranchName", reflect.TypeOf((*MockRepoClient)(nil).GetDefaultBranchName))
}

// GetFileReader mocks base method.
func (m *MockRepoClient) GetFileReader(filename string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileReader", filename)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileReader indicates an expected call of GetFileReader.
func (mr *MockRepoClientMockRecorder) GetFileReader(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileReader", reflect.TypeOf((*MockRepoClient)(nil).GetFileReader), filename)
}

// GetOrgRepoClient mocks base method.
func (m *MockRepoClient) GetOrgRepoClient(arg0 context.Context) (clients.RepoClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgRepoClient", arg0)
	ret0, _ := ret[0].(clients.RepoClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgRepoClient indicates an expected call of GetOrgRepoClient.
func (mr *MockRepoClientMockRecorder) GetOrgRepoClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgRepoClient", reflect.TypeOf((*MockRepoClient)(nil).GetOrgRepoClient), arg0)
}

// InitRepo mocks base method.
func (m *MockRepoClient) InitRepo(repo clients.Repo, commitSHA string, commitDepth int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitRepo", repo, commitSHA, commitDepth)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitRepo indicates an expected call of InitRepo.
func (mr *MockRepoClientMockRecorder) InitRepo(repo, commitSHA, commitDepth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitRepo", reflect.TypeOf((*MockRepoClient)(nil).InitRepo), repo, commitSHA, commitDepth)
}

// IsArchived mocks base method.
func (m *MockRepoClient) IsArchived() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsArchived")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsArchived indicates an expected call of IsArchived.
func (mr *MockRepoClientMockRecorder) IsArchived() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsArchived", reflect.TypeOf((*MockRepoClient)(nil).IsArchived))
}

// ListCheckRunsForRef mocks base method.
func (m *MockRepoClient) ListCheckRunsForRef(ref string) ([]clients.CheckRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCheckRunsForRef", ref)
	ret0, _ := ret[0].([]clients.CheckRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCheckRunsForRef indicates an expected call of ListCheckRunsForRef.
func (mr *MockRepoClientMockRecorder) ListCheckRunsForRef(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCheckRunsForRef", reflect.TypeOf((*MockRepoClient)(nil).ListCheckRunsForRef), ref)
}

// ListCommits mocks base method.
func (m *MockRepoClient) ListCommits() ([]clients.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits")
	ret0, _ := ret[0].([]clients.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockRepoClientMockRecorder) ListCommits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockRepoClient)(nil).ListCommits))
}

// ListContributors mocks base method.
func (m *MockRepoClient) ListContributors() ([]clients.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContributors")
	ret0, _ := ret[0].([]clients.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContributors indicates an expected call of ListContributors.
func (mr *MockRepoClientMockRecorder) ListContributors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContributors", reflect.TypeOf((*MockRepoClient)(nil).ListContributors))
}

// ListFiles mocks base method.
func (m *MockRepoClient) ListFiles(predicate func(string) (bool, error)) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", predicate)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockRepoClientMockRecorder) ListFiles(predicate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockRepoClient)(nil).ListFiles), predicate)
}

// ListIssues mocks base method.
func (m *MockRepoClient) ListIssues() ([]clients.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIssues")
	ret0, _ := ret[0].([]clients.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIssues indicates an expected call of ListIssues.
func (mr *MockRepoClientMockRecorder) ListIssues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIssues", reflect.TypeOf((*MockRepoClient)(nil).ListIssues))
}

// ListLicenses mocks base method.
func (m *MockRepoClient) ListLicenses() ([]clients.License, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLicenses")
	ret0, _ := ret[0].([]clients.License)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLicenses indicates an expected call of ListLicenses.
func (mr *MockRepoClientMockRecorder) ListLicenses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLicenses", reflect.TypeOf((*MockRepoClient)(nil).ListLicenses))
}

// ListProgrammingLanguages mocks base method.
func (m *MockRepoClient) ListProgrammingLanguages() ([]clients.Language, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProgrammingLanguages")
	ret0, _ := ret[0].([]clients.Language)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProgrammingLanguages indicates an expected call of ListProgrammingLanguages.
func (mr *MockRepoClientMockRecorder) ListProgrammingLanguages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProgrammingLanguages", reflect.TypeOf((*MockRepoClient)(nil).ListProgrammingLanguages))
}

// ListReleases mocks base method.
func (m *MockRepoClient) ListReleases() ([]clients.Release, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReleases")
	ret0, _ := ret[0].([]clients.Release)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReleases indicates an expected call of ListReleases.
func (mr *MockRepoClientMockRecorder) ListReleases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReleases", reflect.TypeOf((*MockRepoClient)(nil).ListReleases))
}

// ListSboms mocks base method.
func (m *MockRepoClient) ListSBOMs() ([]clients.SBOM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSboms")
	ret0, _ := ret[0].([]clients.SBOM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSboms indicates an expected call of ListSboms.
func (mr *MockRepoClientMockRecorder) ListSboms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSboms", reflect.TypeOf((*MockRepoClient)(nil).ListSBOMs))
}

// ListStatuses mocks base method.
func (m *MockRepoClient) ListStatuses(ref string) ([]clients.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStatuses", ref)
	ret0, _ := ret[0].([]clients.Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStatuses indicates an expected call of ListStatuses.
func (mr *MockRepoClientMockRecorder) ListStatuses(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStatuses", reflect.TypeOf((*MockRepoClient)(nil).ListStatuses), ref)
}

// ListSuccessfulWorkflowRuns mocks base method.
func (m *MockRepoClient) ListSuccessfulWorkflowRuns(filename string) ([]clients.WorkflowRun, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSuccessfulWorkflowRuns", filename)
	ret0, _ := ret[0].([]clients.WorkflowRun)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSuccessfulWorkflowRuns indicates an expected call of ListSuccessfulWorkflowRuns.
func (mr *MockRepoClientMockRecorder) ListSuccessfulWorkflowRuns(filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSuccessfulWorkflowRuns", reflect.TypeOf((*MockRepoClient)(nil).ListSuccessfulWorkflowRuns), filename)
}

// ListWebhooks mocks base method.
func (m *MockRepoClient) ListWebhooks() ([]clients.Webhook, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWebhooks")
	ret0, _ := ret[0].([]clients.Webhook)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWebhooks indicates an expected call of ListWebhooks.
func (mr *MockRepoClientMockRecorder) ListWebhooks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebhooks", reflect.TypeOf((*MockRepoClient)(nil).ListWebhooks))
}

// LocalPath mocks base method.
func (m *MockRepoClient) LocalPath() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalPath")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocalPath indicates an expected call of LocalPath.
func (mr *MockRepoClientMockRecorder) LocalPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalPath", reflect.TypeOf((*MockRepoClient)(nil).LocalPath))
}

// Search mocks base method.
func (m *MockRepoClient) Search(request clients.SearchRequest) (clients.SearchResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", request)
	ret0, _ := ret[0].(clients.SearchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockRepoClientMockRecorder) Search(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockRepoClient)(nil).Search), request)
}

// SearchCommits mocks base method.
func (m *MockRepoClient) SearchCommits(request clients.SearchCommitsOptions) ([]clients.Commit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCommits", request)
	ret0, _ := ret[0].([]clients.Commit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCommits indicates an expected call of SearchCommits.
func (mr *MockRepoClientMockRecorder) SearchCommits(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCommits", reflect.TypeOf((*MockRepoClient)(nil).SearchCommits), request)
}

// URI mocks base method.
func (m *MockRepoClient) URI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URI")
	ret0, _ := ret[0].(string)
	return ret0
}

// URI indicates an expected call of URI.
func (mr *MockRepoClientMockRecorder) URI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URI", reflect.TypeOf((*MockRepoClient)(nil).URI))
}
