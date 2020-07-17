// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nukosuke/go-zendesk/zendesk (interfaces: API)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	zendesk "github.com/nukosuke/go-zendesk/zendesk"
	reflect "reflect"
)

// Client is a mock of API interface
type Client struct {
	ctrl     *gomock.Controller
	recorder *ClientMockRecorder
}

// ClientMockRecorder is the mock recorder for Client
type ClientMockRecorder struct {
	mock *Client
}

// NewClient creates a new mock instance
func NewClient(ctrl *gomock.Controller) *Client {
	mock := &Client{ctrl: ctrl}
	mock.recorder = &ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Client) EXPECT() *ClientMockRecorder {
	return m.recorder
}

// CreateAutomation mocks base method
func (m *Client) CreateAutomation(arg0 context.Context, arg1 zendesk.Automation) (zendesk.Automation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAutomation", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Automation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAutomation indicates an expected call of CreateAutomation
func (mr *ClientMockRecorder) CreateAutomation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAutomation", reflect.TypeOf((*Client)(nil).CreateAutomation), arg0, arg1)
}

// CreateBrand mocks base method
func (m *Client) CreateBrand(arg0 context.Context, arg1 zendesk.Brand) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBrand", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBrand indicates an expected call of CreateBrand
func (mr *ClientMockRecorder) CreateBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBrand", reflect.TypeOf((*Client)(nil).CreateBrand), arg0, arg1)
}

// CreateDynamicContentItem mocks base method
func (m *Client) CreateDynamicContentItem(arg0 context.Context, arg1 zendesk.DynamicContentItem) (zendesk.DynamicContentItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDynamicContentItem", arg0, arg1)
	ret0, _ := ret[0].(zendesk.DynamicContentItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDynamicContentItem indicates an expected call of CreateDynamicContentItem
func (mr *ClientMockRecorder) CreateDynamicContentItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDynamicContentItem", reflect.TypeOf((*Client)(nil).CreateDynamicContentItem), arg0, arg1)
}

// CreateGroup mocks base method
func (m *Client) CreateGroup(arg0 context.Context, arg1 zendesk.Group) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup
func (mr *ClientMockRecorder) CreateGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*Client)(nil).CreateGroup), arg0, arg1)
}

// CreateOrganization mocks base method
func (m *Client) CreateOrganization(arg0 context.Context, arg1 zendesk.Organization) (zendesk.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization
func (mr *ClientMockRecorder) CreateOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*Client)(nil).CreateOrganization), arg0, arg1)
}

// CreateSLAPolicy mocks base method
func (m *Client) CreateSLAPolicy(arg0 context.Context, arg1 zendesk.SLAPolicy) (zendesk.SLAPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSLAPolicy", arg0, arg1)
	ret0, _ := ret[0].(zendesk.SLAPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSLAPolicy indicates an expected call of CreateSLAPolicy
func (mr *ClientMockRecorder) CreateSLAPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSLAPolicy", reflect.TypeOf((*Client)(nil).CreateSLAPolicy), arg0, arg1)
}

// CreateTarget mocks base method
func (m *Client) CreateTarget(arg0 context.Context, arg1 zendesk.Target) (zendesk.Target, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTarget", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Target)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTarget indicates an expected call of CreateTarget
func (mr *ClientMockRecorder) CreateTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTarget", reflect.TypeOf((*Client)(nil).CreateTarget), arg0, arg1)
}

// CreateTicket mocks base method
func (m *Client) CreateTicket(arg0 context.Context, arg1 zendesk.Ticket) (zendesk.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicket indicates an expected call of CreateTicket
func (mr *ClientMockRecorder) CreateTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*Client)(nil).CreateTicket), arg0, arg1)
}

// CreateTicketField mocks base method
func (m *Client) CreateTicketField(arg0 context.Context, arg1 zendesk.TicketField) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicketField", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicketField indicates an expected call of CreateTicketField
func (mr *ClientMockRecorder) CreateTicketField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicketField", reflect.TypeOf((*Client)(nil).CreateTicketField), arg0, arg1)
}

// CreateTicketForm mocks base method
func (m *Client) CreateTicketForm(arg0 context.Context, arg1 zendesk.TicketForm) (zendesk.TicketForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicketForm", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicketForm indicates an expected call of CreateTicketForm
func (mr *ClientMockRecorder) CreateTicketForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicketForm", reflect.TypeOf((*Client)(nil).CreateTicketForm), arg0, arg1)
}

// CreateTrigger mocks base method
func (m *Client) CreateTrigger(arg0 context.Context, arg1 zendesk.Trigger) (zendesk.Trigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrigger", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrigger indicates an expected call of CreateTrigger
func (mr *ClientMockRecorder) CreateTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrigger", reflect.TypeOf((*Client)(nil).CreateTrigger), arg0, arg1)
}

// CreateUser mocks base method
func (m *Client) CreateUser(arg0 context.Context, arg1 zendesk.User) (zendesk.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(zendesk.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *ClientMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*Client)(nil).CreateUser), arg0, arg1)
}

// DeleteAutomation mocks base method
func (m *Client) DeleteAutomation(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAutomation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAutomation indicates an expected call of DeleteAutomation
func (mr *ClientMockRecorder) DeleteAutomation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAutomation", reflect.TypeOf((*Client)(nil).DeleteAutomation), arg0, arg1)
}

// DeleteBrand mocks base method
func (m *Client) DeleteBrand(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBrand", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBrand indicates an expected call of DeleteBrand
func (mr *ClientMockRecorder) DeleteBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBrand", reflect.TypeOf((*Client)(nil).DeleteBrand), arg0, arg1)
}

// DeleteGroup mocks base method
func (m *Client) DeleteGroup(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup
func (mr *ClientMockRecorder) DeleteGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*Client)(nil).DeleteGroup), arg0, arg1)
}

// DeleteOrganization mocks base method
func (m *Client) DeleteOrganization(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganization", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganization indicates an expected call of DeleteOrganization
func (mr *ClientMockRecorder) DeleteOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganization", reflect.TypeOf((*Client)(nil).DeleteOrganization), arg0, arg1)
}

// DeleteSLAPolicy mocks base method
func (m *Client) DeleteSLAPolicy(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSLAPolicy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSLAPolicy indicates an expected call of DeleteSLAPolicy
func (mr *ClientMockRecorder) DeleteSLAPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSLAPolicy", reflect.TypeOf((*Client)(nil).DeleteSLAPolicy), arg0, arg1)
}

// DeleteTarget mocks base method
func (m *Client) DeleteTarget(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTarget", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTarget indicates an expected call of DeleteTarget
func (mr *ClientMockRecorder) DeleteTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTarget", reflect.TypeOf((*Client)(nil).DeleteTarget), arg0, arg1)
}

// DeleteTicket mocks base method
func (m *Client) DeleteTicket(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicket indicates an expected call of DeleteTicket
func (mr *ClientMockRecorder) DeleteTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicket", reflect.TypeOf((*Client)(nil).DeleteTicket), arg0, arg1)
}

// DeleteTicketField mocks base method
func (m *Client) DeleteTicketField(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicketField", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicketField indicates an expected call of DeleteTicketField
func (mr *ClientMockRecorder) DeleteTicketField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicketField", reflect.TypeOf((*Client)(nil).DeleteTicketField), arg0, arg1)
}

// DeleteTicketForm mocks base method
func (m *Client) DeleteTicketForm(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicketForm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTicketForm indicates an expected call of DeleteTicketForm
func (mr *ClientMockRecorder) DeleteTicketForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicketForm", reflect.TypeOf((*Client)(nil).DeleteTicketForm), arg0, arg1)
}

// DeleteTrigger mocks base method
func (m *Client) DeleteTrigger(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrigger", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrigger indicates an expected call of DeleteTrigger
func (mr *ClientMockRecorder) DeleteTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrigger", reflect.TypeOf((*Client)(nil).DeleteTrigger), arg0, arg1)
}

// DeleteUpload mocks base method
func (m *Client) DeleteUpload(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUpload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUpload indicates an expected call of DeleteUpload
func (mr *ClientMockRecorder) DeleteUpload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUpload", reflect.TypeOf((*Client)(nil).DeleteUpload), arg0, arg1)
}

// GetAttachment mocks base method
func (m *Client) GetAttachment(arg0 context.Context, arg1 int64) (zendesk.Attachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachment", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Attachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttachment indicates an expected call of GetAttachment
func (mr *ClientMockRecorder) GetAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachment", reflect.TypeOf((*Client)(nil).GetAttachment), arg0, arg1)
}

// GetAutomation mocks base method
func (m *Client) GetAutomation(arg0 context.Context, arg1 int64) (zendesk.Automation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutomation", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Automation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAutomation indicates an expected call of GetAutomation
func (mr *ClientMockRecorder) GetAutomation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutomation", reflect.TypeOf((*Client)(nil).GetAutomation), arg0, arg1)
}

// GetAutomations mocks base method
func (m *Client) GetAutomations(arg0 context.Context, arg1 *zendesk.AutomationListOptions) ([]zendesk.Automation, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAutomations", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.Automation)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAutomations indicates an expected call of GetAutomations
func (mr *ClientMockRecorder) GetAutomations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAutomations", reflect.TypeOf((*Client)(nil).GetAutomations), arg0, arg1)
}

// GetBrand mocks base method
func (m *Client) GetBrand(arg0 context.Context, arg1 int64) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBrand", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBrand indicates an expected call of GetBrand
func (mr *ClientMockRecorder) GetBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBrand", reflect.TypeOf((*Client)(nil).GetBrand), arg0, arg1)
}

// GetDynamicContentItems mocks base method
func (m *Client) GetDynamicContentItems(arg0 context.Context) ([]zendesk.DynamicContentItem, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDynamicContentItems", arg0)
	ret0, _ := ret[0].([]zendesk.DynamicContentItem)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDynamicContentItems indicates an expected call of GetDynamicContentItems
func (mr *ClientMockRecorder) GetDynamicContentItems(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDynamicContentItems", reflect.TypeOf((*Client)(nil).GetDynamicContentItems), arg0)
}

// GetGroup mocks base method
func (m *Client) GetGroup(arg0 context.Context, arg1 int64) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup
func (mr *ClientMockRecorder) GetGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*Client)(nil).GetGroup), arg0, arg1)
}

// GetGroups mocks base method
func (m *Client) GetGroups(arg0 context.Context) ([]zendesk.Group, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroups", arg0)
	ret0, _ := ret[0].([]zendesk.Group)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetGroups indicates an expected call of GetGroups
func (mr *ClientMockRecorder) GetGroups(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*Client)(nil).GetGroups), arg0)
}

// GetLocales mocks base method
func (m *Client) GetLocales(arg0 context.Context) ([]zendesk.Locale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocales", arg0)
	ret0, _ := ret[0].([]zendesk.Locale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocales indicates an expected call of GetLocales
func (mr *ClientMockRecorder) GetLocales(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocales", reflect.TypeOf((*Client)(nil).GetLocales), arg0)
}

// GetMultipleTickets mocks base method
func (m *Client) GetMultipleTickets(arg0 context.Context, arg1 []int64) ([]zendesk.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMultipleTickets", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMultipleTickets indicates an expected call of GetMultipleTickets
func (mr *ClientMockRecorder) GetMultipleTickets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMultipleTickets", reflect.TypeOf((*Client)(nil).GetMultipleTickets), arg0, arg1)
}

// GetOrganization mocks base method
func (m *Client) GetOrganization(arg0 context.Context, arg1 int64) (zendesk.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganization", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganization indicates an expected call of GetOrganization
func (mr *ClientMockRecorder) GetOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganization", reflect.TypeOf((*Client)(nil).GetOrganization), arg0, arg1)
}

// GetSLAPolicies mocks base method
func (m *Client) GetSLAPolicies(arg0 context.Context, arg1 *zendesk.SLAPolicyListOptions) ([]zendesk.SLAPolicy, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSLAPolicies", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.SLAPolicy)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSLAPolicies indicates an expected call of GetSLAPolicies
func (mr *ClientMockRecorder) GetSLAPolicies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSLAPolicies", reflect.TypeOf((*Client)(nil).GetSLAPolicies), arg0, arg1)
}

// GetSLAPolicy mocks base method
func (m *Client) GetSLAPolicy(arg0 context.Context, arg1 int64) (zendesk.SLAPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSLAPolicy", arg0, arg1)
	ret0, _ := ret[0].(zendesk.SLAPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSLAPolicy indicates an expected call of GetSLAPolicy
func (mr *ClientMockRecorder) GetSLAPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSLAPolicy", reflect.TypeOf((*Client)(nil).GetSLAPolicy), arg0, arg1)
}

// GetTarget mocks base method
func (m *Client) GetTarget(arg0 context.Context, arg1 int64) (zendesk.Target, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTarget", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Target)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTarget indicates an expected call of GetTarget
func (mr *ClientMockRecorder) GetTarget(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTarget", reflect.TypeOf((*Client)(nil).GetTarget), arg0, arg1)
}

// GetTargets mocks base method
func (m *Client) GetTargets(arg0 context.Context) ([]zendesk.Target, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargets", arg0)
	ret0, _ := ret[0].([]zendesk.Target)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTargets indicates an expected call of GetTargets
func (mr *ClientMockRecorder) GetTargets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargets", reflect.TypeOf((*Client)(nil).GetTargets), arg0)
}

// GetTicket mocks base method
func (m *Client) GetTicket(arg0 context.Context, arg1 int64) (zendesk.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicket", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket
func (mr *ClientMockRecorder) GetTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*Client)(nil).GetTicket), arg0, arg1)
}

// GetTicketField mocks base method
func (m *Client) GetTicketField(arg0 context.Context, arg1 int64) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketField", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicketField indicates an expected call of GetTicketField
func (mr *ClientMockRecorder) GetTicketField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketField", reflect.TypeOf((*Client)(nil).GetTicketField), arg0, arg1)
}

// GetTicketFields mocks base method
func (m *Client) GetTicketFields(arg0 context.Context) ([]zendesk.TicketField, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketFields", arg0)
	ret0, _ := ret[0].([]zendesk.TicketField)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTicketFields indicates an expected call of GetTicketFields
func (mr *ClientMockRecorder) GetTicketFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketFields", reflect.TypeOf((*Client)(nil).GetTicketFields), arg0)
}

// GetTicketForm mocks base method
func (m *Client) GetTicketForm(arg0 context.Context, arg1 int64) (zendesk.TicketForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketForm", arg0, arg1)
	ret0, _ := ret[0].(zendesk.TicketForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicketForm indicates an expected call of GetTicketForm
func (mr *ClientMockRecorder) GetTicketForm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketForm", reflect.TypeOf((*Client)(nil).GetTicketForm), arg0, arg1)
}

// GetTicketForms mocks base method
func (m *Client) GetTicketForms(arg0 context.Context, arg1 *zendesk.TicketFormListOptions) ([]zendesk.TicketForm, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicketForms", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.TicketForm)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTicketForms indicates an expected call of GetTicketForms
func (mr *ClientMockRecorder) GetTicketForms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicketForms", reflect.TypeOf((*Client)(nil).GetTicketForms), arg0, arg1)
}

// GetTickets mocks base method
func (m *Client) GetTickets(arg0 context.Context, arg1 *zendesk.TicketListOptions) ([]zendesk.Ticket, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTickets", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.Ticket)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTickets indicates an expected call of GetTickets
func (mr *ClientMockRecorder) GetTickets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTickets", reflect.TypeOf((*Client)(nil).GetTickets), arg0, arg1)
}

// GetTrigger mocks base method
func (m *Client) GetTrigger(arg0 context.Context, arg1 int64) (zendesk.Trigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrigger", arg0, arg1)
	ret0, _ := ret[0].(zendesk.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrigger indicates an expected call of GetTrigger
func (mr *ClientMockRecorder) GetTrigger(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrigger", reflect.TypeOf((*Client)(nil).GetTrigger), arg0, arg1)
}

// GetTriggers mocks base method
func (m *Client) GetTriggers(arg0 context.Context, arg1 *zendesk.TriggerListOptions) ([]zendesk.Trigger, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTriggers", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.Trigger)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTriggers indicates an expected call of GetTriggers
func (mr *ClientMockRecorder) GetTriggers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTriggers", reflect.TypeOf((*Client)(nil).GetTriggers), arg0, arg1)
}

// GetUser mocks base method
func (m *Client) GetUser(arg0 context.Context, arg1 int64) (zendesk.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(zendesk.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *ClientMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*Client)(nil).GetUser), arg0, arg1)
}

// GetUserFields mocks base method
func (m *Client) GetUserFields(arg0 context.Context, arg1 *zendesk.UserFieldListOptions) ([]zendesk.UserField, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFields", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.UserField)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserFields indicates an expected call of GetUserFields
func (mr *ClientMockRecorder) GetUserFields(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFields", reflect.TypeOf((*Client)(nil).GetUserFields), arg0, arg1)
}

// GetUsers mocks base method
func (m *Client) GetUsers(arg0 context.Context, arg1 *zendesk.UserListOptions) ([]zendesk.User, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].([]zendesk.User)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUsers indicates an expected call of GetUsers
func (mr *ClientMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*Client)(nil).GetUsers), arg0, arg1)
}

// Search mocks base method
func (m *Client) Search(arg0 context.Context, arg1 *zendesk.SearchOptions) (zendesk.SearchResults, zendesk.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1)
	ret0, _ := ret[0].(zendesk.SearchResults)
	ret1, _ := ret[1].(zendesk.Page)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Search indicates an expected call of Search
func (mr *ClientMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*Client)(nil).Search), arg0, arg1)
}

// UpdateAutomation mocks base method
func (m *Client) UpdateAutomation(arg0 context.Context, arg1 int64, arg2 zendesk.Automation) (zendesk.Automation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAutomation", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Automation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAutomation indicates an expected call of UpdateAutomation
func (mr *ClientMockRecorder) UpdateAutomation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAutomation", reflect.TypeOf((*Client)(nil).UpdateAutomation), arg0, arg1, arg2)
}

// UpdateBrand mocks base method
func (m *Client) UpdateBrand(arg0 context.Context, arg1 int64, arg2 zendesk.Brand) (zendesk.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBrand", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBrand indicates an expected call of UpdateBrand
func (mr *ClientMockRecorder) UpdateBrand(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBrand", reflect.TypeOf((*Client)(nil).UpdateBrand), arg0, arg1, arg2)
}

// UpdateGroup mocks base method
func (m *Client) UpdateGroup(arg0 context.Context, arg1 int64, arg2 zendesk.Group) (zendesk.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGroup indicates an expected call of UpdateGroup
func (mr *ClientMockRecorder) UpdateGroup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*Client)(nil).UpdateGroup), arg0, arg1, arg2)
}

// UpdateOrganization mocks base method
func (m *Client) UpdateOrganization(arg0 context.Context, arg1 int64, arg2 zendesk.Organization) (zendesk.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrganization", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganization indicates an expected call of UpdateOrganization
func (mr *ClientMockRecorder) UpdateOrganization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganization", reflect.TypeOf((*Client)(nil).UpdateOrganization), arg0, arg1, arg2)
}

// UpdateSLAPolicy mocks base method
func (m *Client) UpdateSLAPolicy(arg0 context.Context, arg1 int64, arg2 zendesk.SLAPolicy) (zendesk.SLAPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSLAPolicy", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.SLAPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSLAPolicy indicates an expected call of UpdateSLAPolicy
func (mr *ClientMockRecorder) UpdateSLAPolicy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSLAPolicy", reflect.TypeOf((*Client)(nil).UpdateSLAPolicy), arg0, arg1, arg2)
}

// UpdateTarget mocks base method
func (m *Client) UpdateTarget(arg0 context.Context, arg1 int64, arg2 zendesk.Target) (zendesk.Target, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Target)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTarget indicates an expected call of UpdateTarget
func (mr *ClientMockRecorder) UpdateTarget(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTarget", reflect.TypeOf((*Client)(nil).UpdateTarget), arg0, arg1, arg2)
}

// UpdateTicket mocks base method
func (m *Client) UpdateTicket(arg0 context.Context, arg1 int64, arg2 zendesk.Ticket) (zendesk.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicket", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicket indicates an expected call of UpdateTicket
func (mr *ClientMockRecorder) UpdateTicket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicket", reflect.TypeOf((*Client)(nil).UpdateTicket), arg0, arg1, arg2)
}

// UpdateTicketField mocks base method
func (m *Client) UpdateTicketField(arg0 context.Context, arg1 int64, arg2 zendesk.TicketField) (zendesk.TicketField, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicketField", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.TicketField)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicketField indicates an expected call of UpdateTicketField
func (mr *ClientMockRecorder) UpdateTicketField(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicketField", reflect.TypeOf((*Client)(nil).UpdateTicketField), arg0, arg1, arg2)
}

// UpdateTicketForm mocks base method
func (m *Client) UpdateTicketForm(arg0 context.Context, arg1 int64, arg2 zendesk.TicketForm) (zendesk.TicketForm, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicketForm", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.TicketForm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicketForm indicates an expected call of UpdateTicketForm
func (mr *ClientMockRecorder) UpdateTicketForm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicketForm", reflect.TypeOf((*Client)(nil).UpdateTicketForm), arg0, arg1, arg2)
}

// UpdateTrigger mocks base method
func (m *Client) UpdateTrigger(arg0 context.Context, arg1 int64, arg2 zendesk.Trigger) (zendesk.Trigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrigger", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.Trigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrigger indicates an expected call of UpdateTrigger
func (mr *ClientMockRecorder) UpdateTrigger(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrigger", reflect.TypeOf((*Client)(nil).UpdateTrigger), arg0, arg1, arg2)
}

// UpdateUser mocks base method
func (m *Client) UpdateUser(arg0 context.Context, arg1 int64, arg2 zendesk.User) (zendesk.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *ClientMockRecorder) UpdateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*Client)(nil).UpdateUser), arg0, arg1, arg2)
}

// UploadAttachment mocks base method
func (m *Client) UploadAttachment(arg0 context.Context, arg1, arg2 string) zendesk.UploadWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAttachment", arg0, arg1, arg2)
	ret0, _ := ret[0].(zendesk.UploadWriter)
	return ret0
}

// UploadAttachment indicates an expected call of UploadAttachment
func (mr *ClientMockRecorder) UploadAttachment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAttachment", reflect.TypeOf((*Client)(nil).UploadAttachment), arg0, arg1, arg2)
}
