package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"privy/internal/repository"
	mock_repo "privy/mock/repository"
	m "privy/models"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		repository repository.Repository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				repository: mockRepository,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.repository)
			_, ok := got.(Handler)
			if !ok {
				t.Errorf("Not Handler interface")
			}
		})
	}
}
func Test_handler_GetListOfCakes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		method string
		path   string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "Success",
			args: args{
				method: http.MethodGet,
				path:   "/cakes?limit=10&offset=0",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetListOfCakes(gomock.Any(), 10, 0).Return([]m.Cake{
					{Id: 1, Title: "title"},
				}, nil)
			},
		},
		{
			name: "Repository error",
			args: args{
				method: http.MethodGet,
				path:   "/cakes?limit=10&offset=0",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetListOfCakes(gomock.Any(), 10, 0).Return(nil, errors.New("repository error"))
			},
		},
		{
			name: "no offset",
			args: args{
				method: http.MethodGet,
				path:   "/cakes?limit=10",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetListOfCakes(gomock.Any(), 10, 0).Return([]m.Cake{
					{Id: 1, Title: "title"},
				}, nil)
			},
		},
		{
			name: "no limit",
			args: args{
				method: http.MethodGet,
				path:   "/cakes?offset=0",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetListOfCakes(gomock.Any(), 100, 0).Return([]m.Cake{
					{Id: 1, Title: "title"},
				}, nil)
			},
		},
		{
			name: "limit not integer",
			args: args{
				method: http.MethodGet,
				path:   "/cakes?limit=not_number&offset=0",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {
				return
			},
		},
		{
			name: "offset not integer",
			args: args{
				method: http.MethodGet,
				path:   "/cakes?limit=100&offset=not_number",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {
				return
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				repository: mockRepository,
			}
			if err := h.GetListOfCakes(c); err != nil {
				t.Errorf("handler.GetListOfCakes() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_GetDetailsOfCake(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		method string
		path   string
		id     string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "Success",
			args: args{
				method: http.MethodGet,
				path:   "/cakes",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().GetDetailsOfCake(gomock.Any(), 1).
					Return(m.Cake{Id: 1, Title: "title", Description: "description", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, nil)
			},
		},
		{
			name: "id not valid",
			args: args{
				method: http.MethodGet,
				path:   "/cakes",
				id:     "abc",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "repository error",
			args: args{
				method: http.MethodGet,
				path:   "/cakes",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().GetDetailsOfCake(gomock.Any(), 1).
					Return(m.Cake{Id: 1, Title: "title", Description: "description", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, errors.New("Repository Error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest("GET", "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			tt.mock()

			h := &handler{
				repository: mockRepository,
			}
			if err := h.GetDetailsOfCake(c); err != nil {
				t.Errorf("handler.GetListOfCakes() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_InsertCake(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		method string
		path   string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "Success",
			args: args{
				method: http.MethodPost,
				path:   "/cakes?title=judul&description=deskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().InsertCake(gomock.Any(), m.Cake{}).Return(m.Cake{Id: 1, Title: "title"}, nil)
			},
		},
		{
			name: "Repository error",
			args: args{
				method: http.MethodPost,
				path:   "/cakes?title=judul&description=deskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().InsertCake(gomock.Any(), m.Cake{}).Return(m.Cake{}, errors.New("repository error"))
			},
		},
		{
			name: "Empty title",
			args: args{
				method: http.MethodPost,
				path:   "/cakes?title=&description=deskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Empty description",
			args: args{
				method: http.MethodPost,
				path:   "/cakes?title=judul&description=&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Invalid rating",
			args: args{
				method: http.MethodPost,
				path:   "/cakes?title=judul&description=deskripsi&rating=abc&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Invalid image",
			args: args{
				method: http.MethodPost,
				path:   "/cakes?title=judul&description=deskripsi&rating=9.8&image=1234",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.mock()

			h := &handler{
				repository: mockRepository,
			}
			if err := h.InsertCake(c); err != nil {
				t.Errorf("handler.GetListOfCakes() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_UpdateCake(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		method string
		path   string
		id     string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "Success",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=newjudul&description=newdeskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCake(
					gomock.Any(), m.Cake{}).
					Return(m.Cake{Id: 1, Title: "newjudul", Description: "newdeskripsi", Rating: 9.8, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, nil)
			},
		},
		{
			name: "Success with empty title",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=&description=newdeskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCake(
					gomock.Any(), m.Cake{}).
					Return(m.Cake{Id: 1, Title: "newjudul", Description: "newdeskripsi", Rating: 9.8, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, nil)
			},
		},
		{
			name: "Success with empty description",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=newjudul&description=&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCake(
					gomock.Any(), m.Cake{}).
					Return(m.Cake{Id: 1, Title: "newjudul", Description: "newdeskripsi", Rating: 9.8, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, nil)
			},
		},
		{
			name: "Success with empty rating",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=hello&description=newdeskripsi&rating=&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCake(
					gomock.Any(), m.Cake{}).
					Return(m.Cake{Id: 1, Title: "newjudul", Description: "newdeskripsi", Rating: 0, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, nil)
			},
		},
		{
			name: "Success with empty image",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=hello&description=newdeskripsi&rating=&image=",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCake(
					gomock.Any(), m.Cake{}).
					Return(m.Cake{Id: 1, Title: "newjudul", Description: "newdeskripsi", Rating: 0, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, nil)
			},
		},
		{
			name: "id wrong format",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=newjudul&description=newdeskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "abc",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "title wrong format",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=#&description=newdeskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "rating wrong format",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=newtitle&description=newdeskripsi&rating=abc&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "image wrong format",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=newtitle&description=newdeskripsi&rating=abc&image=httttps://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Internal Server Error",
			args: args{
				method: http.MethodPatch,
				path:   "/cakes?title=newjudul&description=newdeskripsi&rating=9.8&image=https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().UpdateCake(
					gomock.Any(), m.Cake{}).
					Return(m.Cake{Id: 1, Title: "newjudul", Description: "newdeskripsi", Rating: 9.8, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"}, errors.New("internal server error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			tt.mock()

			h := &handler{
				repository: mockRepository,
			}
			if err := h.UpdateCake(c); err != nil {
				t.Errorf("handler.UpdateCake() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
func Test_handler_DeleteCake(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mock_repo.NewMockRepository(ctrl)

	type args struct {
		method string
		path   string
		id     string
	}
	type wants struct {
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		wants wants
		mock  func()
	}{
		{
			name: "Success",
			args: args{
				method: http.MethodDelete,
				path:   "/cakes",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusOK,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteCake(gomock.Any(), 1).
					Return(nil)
			},
		},
		{
			name: "id wrong format",
			args: args{
				method: http.MethodDelete,
				path:   "/cakes",
				id:     "abc",
			},
			wants: wants{
				statusCode: http.StatusBadRequest,
			},
			mock: func() {},
		},
		{
			name: "Internal Server Error",
			args: args{
				method: http.MethodDelete,
				path:   "/cakes",
				id:     "1",
			},
			wants: wants{
				statusCode: http.StatusInternalServerError,
			},
			mock: func() {
				mockRepository.EXPECT().DeleteCake(gomock.Any(), 1).
					Return(errors.New("internal server error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetPath("/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			tt.mock()

			h := &handler{
				repository: mockRepository,
			}
			if err := h.DeleteCake(c); err != nil {
				t.Errorf("handler.GetListOfCakes() error = %v", err)
			}

			assert.Equal(t, tt.wants.statusCode, rec.Code)
		})
	}
}
