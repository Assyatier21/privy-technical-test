package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	m "privy/models"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, _, _ := sqlmock.New()

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.db)
			_, ok := got.(Repository)
			if !ok {
				t.Errorf("Not Repository interface")
			}
		})
	}
}
func Test_repository_GetListOfCakes(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx    context.Context
		limit  int
		offset int
	}
	tests := []struct {
		name    string
		args    args
		want    []m.Cake
		wantErr bool
		mock    func()
	}{
		{
			name: "Success",
			args: args{
				ctx:    ctx,
				limit:  10,
				offset: 0,
			},
			want: []m.Cake{
				{Id: 1, Title: "title", Description: "description", Rating: 10, Image: "https://www.abc.com/abc.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: "2022-12-01 20:29:00"},
				{Id: 2, Title: "title2", Description: "description2", Rating: 20, Image: "https://www.abc.com/abc.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: "2022-12-01 20:29:00"},
			},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow(1, "title", "description", 10, "https://www.abc.com/abc.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00").
					AddRow(2, "title2", "description2", 20, "https://www.abc.com/abc.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes ORDER BY rating DESC, title ASC LIMIT 10 OFFSET 0`)).WillReturnRows(rows)
			},
		},
		{
			name: "Query error",
			args: args{
				ctx:    ctx,
				limit:  10,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes ORDER BY rating DESC, title ASC LIMIT 10 OFFSET 0`)).WillReturnError(errors.New("query error"))
			},
		},
		{
			name: "Scan error",
			args: args{
				ctx:    ctx,
				limit:  10,
				offset: 0,
			},
			want:    nil,
			wantErr: true,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow("not number", "title", "description", 10, "https://www.abc.com/abc.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes ORDER BY rating DESC, title ASC LIMIT 10 OFFSET 0`)).WillReturnRows(rows)
			},
		},
		{
			name: "Success",
			args: args{
				ctx:    ctx,
				limit:  10,
				offset: 0,
			},
			want:    []m.Cake{},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"})
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes ORDER BY rating DESC, title ASC LIMIT 10 OFFSET 0`)).WillReturnRows(rows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &repository{
				db: db,
			}
			got, err := r.GetListOfCakes(tt.args.ctx, tt.args.limit, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetListOfCakes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetListOfCakes() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_repository_GetDetailsOfCake(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    m.Cake
		wantErr bool
		mock    func()
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.Cake{Id: 1, Title: "title", Description: "description", Rating: 10, Image: "https://www.abc.com/abc.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: "2022-12-01 20:29:00"},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow(1, "title", "description", 10, "https://www.abc.com/abc.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes WHERE id='1'`)).WillReturnRows(rows)
			},
		},
		{
			name: "err while scanning",
			args: args{
				ctx: ctx,
				id:  1,
			},
			want:    m.Cake{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes WHERE id='1'`)).WillReturnError(errors.New("query error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &repository{
				db: db,
			}
			got, err := r.GetDetailsOfCake(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetDetailsOfCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetDetailsOfCake() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_repository_InsertCake(t *testing.T) {
	ctx := context.Background()
	ct := time.Now()
	currentTime := fmt.Sprintf("%d-%d-%d %d:%d:%d", ct.Year(), ct.Month(), ct.Day(), ct.Hour(), ct.Minute(), ct.Second())

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx  context.Context
		cake m.Cake
	}
	tests := []struct {
		name    string
		args    args
		want    m.Cake
		wantErr bool
		mock    func()
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "title",
					Description: "desc",
					Rating:      10,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   currentTime,
					UpdatedAt:   currentTime,
				},
			},
			want:    m.Cake{Id: 1, Title: "title", Description: "desc", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", CreatedAt: currentTime, UpdatedAt: currentTime},
			wantErr: false,
			mock: func() {
				query := `INSERT INTO privy_cakes`
				sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))
			},
		},
		{
			name: "Query Error",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "title",
					Description: "desc",
					Rating:      10,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   currentTime,
					UpdatedAt:   currentTime,
				},
			},
			want:    m.Cake{},
			wantErr: true,
			mock: func() {
				query := `UPDATE INTO privy_cakes`
				sqlMock.ExpectExec(query).WillReturnError(errors.New("Query Error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &repository{
				db: db,
			}
			got, err := r.InsertCake(tt.args.ctx, tt.args.cake)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.InsertCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.InsertCake() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_repository_UpdateCake(t *testing.T) {
	ctx := context.Background()
	ct := time.Now()
	currentTime := fmt.Sprintf("%d-%d-%d %d:%d:%d", ct.Year(), ct.Month(), ct.Day(), ct.Hour(), ct.Minute(), ct.Second())

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx  context.Context
		cake m.Cake
	}
	tests := []struct {
		name    string
		args    args
		want    m.Cake
		wantErr bool
		mock    func()
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "newtitle",
					Description: "newdesc",
					Rating:      10,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   "2022-12-01 20:29:00",
					UpdatedAt:   "2022-12-01 20:29:00",
				},
			},
			want:    m.Cake{Id: 1, Title: "newtitle", Description: "newdesc", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: currentTime},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow(1, "title", "description", 10, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes`)).WillReturnRows(rows)
				sqlMock.ExpectExec("UPDATE privy_cakes").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Success With Empty Title",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "",
					Description: "newdesc",
					Rating:      10,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   "2022-12-01 20:29:00",
					UpdatedAt:   "2022-12-01 20:29:00",
				},
			},
			want:    m.Cake{Id: 1, Title: "title", Description: "newdesc", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: currentTime},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow(1, "title", "description", 10, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes`)).WillReturnRows(rows)
				sqlMock.ExpectExec("UPDATE privy_cakes").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Success With Empty Desc",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "newtitle",
					Description: "",
					Rating:      10,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   "2022-12-01 20:29:00",
					UpdatedAt:   "2022-12-01 20:29:00",
				},
			},
			want:    m.Cake{Id: 1, Title: "newtitle", Description: "description", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: currentTime},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow(1, "title", "description", 10, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes`)).WillReturnRows(rows)
				sqlMock.ExpectExec("UPDATE privy_cakes").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Success With Empty Rating",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "newtitle",
					Description: "newdesc",
					Rating:      0,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   "2022-12-01 20:29:00",
					UpdatedAt:   "2022-12-01 20:29:00",
				},
			},
			want:    m.Cake{Id: 1, Title: "newtitle", Description: "newdesc", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: currentTime},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow(1, "title", "description", 10, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes`)).WillReturnRows(rows)
				sqlMock.ExpectExec("UPDATE privy_cakes").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Success With Empty Image",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "newtitle",
					Description: "newdesc",
					Rating:      10,
					Image:       "",
					CreatedAt:   "2022-12-01 20:29:00",
					UpdatedAt:   "2022-12-01 20:29:00",
				},
			},
			want:    m.Cake{Id: 1, Title: "newtitle", Description: "newdesc", Rating: 10, Image: "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", CreatedAt: "2022-12-01 20:29:00", UpdatedAt: currentTime},
			wantErr: false,
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
					AddRow(1, "title", "description", 10, "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg", "2022-12-01 20:29:00", "2022-12-01 20:29:00")
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM privy_cakes`)).WillReturnRows(rows)
				sqlMock.ExpectExec("UPDATE privy_cakes").WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "wrong select query",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "newtitle",
					Description: "newdesc",
					Rating:      10,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   "2022-12-01 20:29:00",
					UpdatedAt:   "2022-12-01 20:29:00",
				},
			},
			want:    m.Cake{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECTTT *  FROM privy_cakes`)).WillReturnError(errors.New("query error"))
			},
		},
		{
			name: "wrong update query",
			args: args{
				ctx: ctx,
				cake: m.Cake{
					Id:          1,
					Title:       "newtitle",
					Description: "newdesc",
					Rating:      10,
					Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
					CreatedAt:   "2022-12-01 20:29:00",
					UpdatedAt:   "2022-12-01 20:29:00",
				},
			},
			want:    m.Cake{},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectQuery(regexp.QuoteMeta(`SELECTTT *  FROM privy_cakes`)).WillReturnError(errors.New("query error"))
				sqlMock.ExpectExec("UPDATE XYZ privy_cakes").WillReturnError(errors.New("query error"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				tt.mock()

				r := &repository{
					db: db,
				}
				got, err := r.UpdateCake(tt.args.ctx, tt.args.cake)
				if (err != nil) != tt.wantErr {
					t.Errorf("repository.UpdateCake() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("repository.UpdateCake() = %v, want %v", got, tt.want)
				}
			})
		})
	}
}
func Test_repository_DeleteCake(t *testing.T) {
	ctx := context.Background()

	db, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: false,
			mock: func() {
				sqlMock.ExpectExec(`DELETE FROM privy_cakes`).
					WillDelayFor(time.Second).
					WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))
			},
		},
		{
			name: "Query Error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(`SELECT FROM privy_cakes`).
					WillDelayFor(time.Second).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("Query Error")))
			},
		},
		{
			name: "No Rows Affected",
			args: args{
				ctx: ctx,
				id:  1,
			},
			wantErr: true,
			mock: func() {
				sqlMock.ExpectExec(`DELETE FROM privy_cakes`).
					WillDelayFor(time.Second).
					WillReturnResult(sqlmock.NewResult(int64(1), int64(0)))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			r := &repository{
				db: db,
			}
			err := r.DeleteCake(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteCake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
