// Code generated by MockGen. DO NOT EDIT.
// Source: ./coins.go

// Package handlers is a generated GoMock package.
package handlers

import (
	entities "crypto-viewer/src/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCoinsUseCase is a mock of CoinsUseCase interface.
type MockCoinsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCoinsUseCaseMockRecorder
}

// MockCoinsUseCaseMockRecorder is the mock recorder for MockCoinsUseCase.
type MockCoinsUseCaseMockRecorder struct {
	mock *MockCoinsUseCase
}

// NewMockCoinsUseCase creates a new mock instance.
func NewMockCoinsUseCase(ctrl *gomock.Controller) *MockCoinsUseCase {
	mock := &MockCoinsUseCase{ctrl: ctrl}
	mock.recorder = &MockCoinsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoinsUseCase) EXPECT() *MockCoinsUseCaseMockRecorder {
	return m.recorder
}

// GetCoins mocks base method.
func (m *MockCoinsUseCase) GetCoins(params map[string]string) (entities.CoinsData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoins", params)
	ret0, _ := ret[0].(entities.CoinsData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoins indicates an expected call of GetCoins.
func (mr *MockCoinsUseCaseMockRecorder) GetCoins(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoins", reflect.TypeOf((*MockCoinsUseCase)(nil).GetCoins), params)
}

// MockSaveDataUseCase is a mock of SaveDataUseCase interface.
type MockSaveDataUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockSaveDataUseCaseMockRecorder
}

// MockSaveDataUseCaseMockRecorder is the mock recorder for MockSaveDataUseCase.
type MockSaveDataUseCaseMockRecorder struct {
	mock *MockSaveDataUseCase
}

// NewMockSaveDataUseCase creates a new mock instance.
func NewMockSaveDataUseCase(ctrl *gomock.Controller) *MockSaveDataUseCase {
	mock := &MockSaveDataUseCase{ctrl: ctrl}
	mock.recorder = &MockSaveDataUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaveDataUseCase) EXPECT() *MockSaveDataUseCaseMockRecorder {
	return m.recorder
}

// SaveCoins mocks base method.
func (m *MockSaveDataUseCase) SaveCoins(coinsData entities.CoinsData) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCoins", coinsData)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveCoins indicates an expected call of SaveCoins.
func (mr *MockSaveDataUseCaseMockRecorder) SaveCoins(coinsData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCoins", reflect.TypeOf((*MockSaveDataUseCase)(nil).SaveCoins), coinsData)
}