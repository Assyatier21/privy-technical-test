package repository

import (
	"context"
	"errors"
	m "privy/models"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

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

