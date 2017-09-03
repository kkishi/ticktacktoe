// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto (interfaces: TickTackToe_GameServer,TickTackToe_GameClient)

// Package mock_ticktacktoe_proto is a generated GoMock package.
package mock_ticktacktoe_proto

import (
	gomock "github.com/golang/mock/gomock"
	ticktacktoe_proto "github.com/kkishi/ticktacktoe/proto/ticktacktoe_proto"
	context "golang.org/x/net/context"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockTickTackToe_GameServer is a mock of TickTackToe_GameServer interface
type MockTickTackToe_GameServer struct {
	ctrl     *gomock.Controller
	recorder *MockTickTackToe_GameServerMockRecorder
}

// MockTickTackToe_GameServerMockRecorder is the mock recorder for MockTickTackToe_GameServer
type MockTickTackToe_GameServerMockRecorder struct {
	mock *MockTickTackToe_GameServer
}

// NewMockTickTackToe_GameServer creates a new mock instance
func NewMockTickTackToe_GameServer(ctrl *gomock.Controller) *MockTickTackToe_GameServer {
	mock := &MockTickTackToe_GameServer{ctrl: ctrl}
	mock.recorder = &MockTickTackToe_GameServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTickTackToe_GameServer) EXPECT() *MockTickTackToe_GameServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockTickTackToe_GameServer) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTickTackToe_GameServerMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).Context))
}

// Recv mocks base method
func (m *MockTickTackToe_GameServer) Recv() (*ticktacktoe_proto.Request, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ticktacktoe_proto.Request)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockTickTackToe_GameServerMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockTickTackToe_GameServer) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTickTackToe_GameServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockTickTackToe_GameServer) Send(arg0 *ticktacktoe_proto.Response) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockTickTackToe_GameServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockTickTackToe_GameServer) SendHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockTickTackToe_GameServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockTickTackToe_GameServer) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTickTackToe_GameServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockTickTackToe_GameServer) SetHeader(arg0 metadata.MD) error {
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockTickTackToe_GameServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockTickTackToe_GameServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockTickTackToe_GameServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockTickTackToe_GameServer)(nil).SetTrailer), arg0)
}

// MockTickTackToe_GameClient is a mock of TickTackToe_GameClient interface
type MockTickTackToe_GameClient struct {
	ctrl     *gomock.Controller
	recorder *MockTickTackToe_GameClientMockRecorder
}

// MockTickTackToe_GameClientMockRecorder is the mock recorder for MockTickTackToe_GameClient
type MockTickTackToe_GameClientMockRecorder struct {
	mock *MockTickTackToe_GameClient
}

// NewMockTickTackToe_GameClient creates a new mock instance
func NewMockTickTackToe_GameClient(ctrl *gomock.Controller) *MockTickTackToe_GameClient {
	mock := &MockTickTackToe_GameClient{ctrl: ctrl}
	mock.recorder = &MockTickTackToe_GameClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTickTackToe_GameClient) EXPECT() *MockTickTackToe_GameClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockTickTackToe_GameClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockTickTackToe_GameClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockTickTackToe_GameClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockTickTackToe_GameClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).Context))
}

// Header mocks base method
func (m *MockTickTackToe_GameClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockTickTackToe_GameClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).Header))
}

// Recv mocks base method
func (m *MockTickTackToe_GameClient) Recv() (*ticktacktoe_proto.Response, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*ticktacktoe_proto.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockTickTackToe_GameClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockTickTackToe_GameClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockTickTackToe_GameClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockTickTackToe_GameClient) Send(arg0 *ticktacktoe_proto.Request) error {
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockTickTackToe_GameClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).Send), arg0)
}

// SendMsg mocks base method
func (m *MockTickTackToe_GameClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockTickTackToe_GameClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockTickTackToe_GameClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockTickTackToe_GameClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockTickTackToe_GameClient)(nil).Trailer))
}