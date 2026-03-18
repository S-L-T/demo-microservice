package presentation_http

import (
	"errors"
	"github.com/S-L-T/demo-microservice/domain/entity"
	"github.com/S-L-T/demo-microservice/domain/use_case"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type UserRepositoryMock struct{}

func (u UserRepositoryMock) AddUser(_ entity.User) (string, error) {
	return "2d2eb499-3785-4301-b2fa-3c3706b5d1b2", nil
}

func (u UserRepositoryMock) UpdateUser(_ entity.User) error {
	return nil
}

func (u UserRepositoryMock) DeleteUser(_ string) error {
	return nil
}

func (u UserRepositoryMock) GetPaginatedUsers(_ entity.Filter, _ uint64, _ uint64) ([]entity.User, error) {
	tc, _ := time.Parse(time.RFC3339, "2022-06-22T12:44:07Z")
	tu, _ := time.Parse(time.RFC3339, "2022-06-24T12:44:07Z")
	return []entity.User{
		{
			ID:        "2d2eb499-3785-4301-b2fa-3c3706b5d1b1",
			FirstName: "TestFirstName1",
			LastName:  "TestLastName1",
			Nickname:  "TestNickname1",
			Password:  "TestPassword1",
			Email:     "TestEmail1",
			Country:   "TestCountry1",
			CreatedAt: tc,
			UpdatedAt: tu,
		},
		{
			ID:        "87bdbbe0-9568-4455-943d-5128f92c8f89",
			FirstName: "TestFirstName2",
			LastName:  "TestLastName2",
			Nickname:  "TestNickname2",
			Password:  "TestPassword2",
			Email:     "TestEmail2",
			Country:   "TestCountry2",
			CreatedAt: tc,
			UpdatedAt: tu,
		},
	}, nil
}

type UserRepositoryMockWithErrors struct{}

func (u UserRepositoryMockWithErrors) AddUser(_ entity.User) (string, error) {
	return "", errors.New("dummy error")
}

func (u UserRepositoryMockWithErrors) UpdateUser(_ entity.User) error {
	return errors.New("dummy error")
}

func (u UserRepositoryMockWithErrors) DeleteUser(_ string) error {
	return errors.New("dummy error")
}

func (u UserRepositoryMockWithErrors) GetPaginatedUsers(_ entity.Filter, _ uint64, _ uint64) ([]entity.User, error) {
	return nil, errors.New("dummy error")
}

func TestServer_healthcheck(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	s := NewServer(use_case.NewUserUseCase(UserRepositoryMock{}))
	s.Router.ServeHTTP(rr, req)

	expectedStatus := http.StatusOK
	if rr.Code != expectedStatus {
		t.Errorf("handler returned unexpected status: got %v want %v",
			rr.Code, expectedStatus)
	}
}

func TestServer_user(t *testing.T) {
	type args struct {
		method          string
		hasErrors       bool
		expectedStatus  int
		expectedMessage string
		payload         string
		useCase        use_case.User
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "GET no errors",
			args: args{
				method:          http.MethodGet,
				expectedStatus:  http.StatusOK,
				expectedMessage: `{"users":[{"id":"2d2eb499-3785-4301-b2fa-3c3706b5d1b1","first_name":"TestFirstName1","last_name":"TestLastName1","nickname":"TestNickname1","password":"TestPassword1","email":"TestEmail1","country":"TestCountry1","created_at":"2022-06-22T12:44:07Z","updated_at":"2022-06-24T12:44:07Z"},{"id":"87bdbbe0-9568-4455-943d-5128f92c8f89","first_name":"TestFirstName2","last_name":"TestLastName2","nickname":"TestNickname2","password":"TestPassword2","email":"TestEmail2","country":"TestCountry2","created_at":"2022-06-22T12:44:07Z","updated_at":"2022-06-24T12:44:07Z"}]}`,
				payload:         `{}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "PUT no errors",
			args: args{
				method:          http.MethodPut,
				expectedStatus:  http.StatusOK,
				expectedMessage: `{"id":"2d2eb499-3785-4301-b2fa-3c3706b5d1b2"}`,
				payload:         `{}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "DELETE no errors",
			args: args{
				method:          http.MethodDelete,
				expectedStatus:  http.StatusOK,
				expectedMessage: `{}`,
				payload:         `{"id":"2d2eb499-3785-4301-b2fa-3c3706b5d1b2"}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "PATCH no errors",
			args: args{
				method:          http.MethodPatch,
				expectedStatus:  http.StatusOK,
				expectedMessage: `{"user":{}}`,
				payload:         `{}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "GET with errors",
			args: args{
				method:          http.MethodGet,
				expectedStatus:  http.StatusInternalServerError,
				expectedMessage: `{"error":"dummy error"}`,
				payload:         `{}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMockWithErrors{}),
			},
		},
		{
			name: "PUT with errors",
			args: args{
				method:          http.MethodPut,
				expectedStatus:  http.StatusInternalServerError,
				expectedMessage: `{"error":"dummy error"}`,
				payload:         `{}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMockWithErrors{}),
			},
		},
		{
			name: "DELETE with errors",
			args: args{
				method:          http.MethodDelete,
				expectedStatus:  http.StatusInternalServerError,
				expectedMessage: `{"error":"dummy error"}`,
				payload:         `{}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMockWithErrors{}),
			},
		},
		{
			name: "PATCH with errors",
			args: args{
				method:          http.MethodPatch,
				expectedStatus:  http.StatusInternalServerError,
				expectedMessage: `{"user":{},"error":"dummy error"}`,
				payload:         `{}`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMockWithErrors{}),
			},
		},
		{
			name: "GET bad request",
			args: args{
				method:          http.MethodGet,
				expectedStatus:  http.StatusBadRequest,
				expectedMessage: `{"error":"json: cannot unmarshal string into Go value of type entity.GetPaginatedUsersRequest"}`,
				payload:         `"bad request"`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "PUT bad request",
			args: args{
				method:          http.MethodPut,
				expectedStatus:  http.StatusBadRequest,
				expectedMessage: `{"error":"json: cannot unmarshal string into Go value of type entity.AddUserRequest"}`,
				payload:         `"bad request"`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "DELETE bad request",
			args: args{
				method:          http.MethodDelete,
				expectedStatus:  http.StatusBadRequest,
				expectedMessage: `{"error":"json: cannot unmarshal string into Go value of type entity.DeleteUserRequest"}`,
				payload:         `"bad request"`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "PATCH bad request",
			args: args{
				method:          http.MethodPatch,
				expectedStatus:  http.StatusBadRequest,
				expectedMessage: `{"user":{},"error":"json: cannot unmarshal string into Go value of type entity.UpdateUserRequest"}`,
				payload:         `"bad request"`,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "OPTIONS",
			args: args{
				method:          http.MethodOptions,
				expectedStatus:  http.StatusOK,
				expectedMessage: ``,
				payload:         ``,
				useCase:        use_case.NewUserUseCase(UserRepositoryMockWithErrors{}),
			},
		},
		{
			name: "HEAD",
			args: args{
				method:          http.MethodHead,
				expectedStatus:  http.StatusMethodNotAllowed,
				expectedMessage: ``,
				payload:         ``,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "POST",
			args: args{
				method:          http.MethodPost,
				expectedStatus:  http.StatusMethodNotAllowed,
				expectedMessage: ``,
				payload:         ``,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "CONNECT",
			args: args{
				method:          http.MethodConnect,
				expectedStatus:  http.StatusMethodNotAllowed,
				expectedMessage: ``,
				payload:         ``,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
		{
			name: "TRACE",
			args: args{
				method:          http.MethodTrace,
				expectedStatus:  http.StatusMethodNotAllowed,
				expectedMessage: ``,
				payload:         ``,
				useCase:        use_case.NewUserUseCase(UserRepositoryMock{}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.args.method, "/user", strings.NewReader(tt.args.payload))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			s := NewServer(tt.args.useCase)
			s.Router.ServeHTTP(rr, req)
			if rr.Code != tt.args.expectedStatus {
				t.Errorf(
					"handler returned unexpected status: got %v want %v",
					rr.Code,
					tt.args.expectedStatus,
				)
			}

			if rr.Body.String() != tt.args.expectedMessage {
				t.Errorf(
					"handler returned unexpected message: got %v want %v",
					rr.Body.String(),
					tt.args.expectedMessage,
				)
			}
		})
	}
}
