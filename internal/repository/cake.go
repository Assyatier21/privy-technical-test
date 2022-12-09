package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"privy/database"
	m "privy/models"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	GetListOfCakes(ctx context.Context, limit int, offset int) ([]m.Cake, error)
	GetDetailsOfCake(ctx context.Context, id int) (m.Cake, error)
	InsertCake(ctx context.Context, cake m.Cake) (m.Cake, error)
	UpdateCake(ctx context.Context, cake m.Cake) (m.Cake, error)
	DeleteCake(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetListOfCakes(ctx context.Context, limit int, offset int) ([]m.Cake, error) {
	var (
		err  error
		rows *sql.Rows
		data []m.Cake
	)

	query := fmt.Sprintf(database.GetListOfCakes, limit, offset)
	rows, err = r.db.Query(query)
	if err != nil {
		log.Println("[GetListOfCakes] can't get list of cakes, err:", err.Error())
		return nil, err
	}

	for rows.Next() {
		var temp = m.Cake{}
		if err := rows.Scan(&temp.Id, &temp.Title, &temp.Description, &temp.Rating, &temp.Image, &temp.CreatedAt, &temp.UpdatedAt); err != nil {
			log.Println("Error query :", err)
			return nil, err
		}
		data = append(data, temp)
	}

	if len(data) > 0 {
		return data, nil
	} else {
		return []m.Cake{}, nil
	}
}

func (r *repository) GetDetailsOfCake(ctx context.Context, id int) (m.Cake, error) {
	var (
		err  error
		cake m.Cake
	)

	query := fmt.Sprintf(database.GetDetailsOfCakeByID, id)
	err = r.db.QueryRow(query).Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err != nil {
		log.Println("[GetDetailsOfCake] can't get details of cake, err:", err.Error())
		return m.Cake{}, err
	}

	return cake, nil
}

func (r *repository) InsertCake(ctx context.Context, cake m.Cake) (m.Cake, error) {
	currentTime := time.Now().String()

	cake.CreatedAt = currentTime
	cake.UpdatedAt = currentTime

	query := fmt.Sprintf(database.InsertCake, cake.Id, cake.Title, cake.Description, cake.Rating, cake.Image, currentTime, currentTime)
	rows, err := r.db.Exec(query)
	if err != nil {
		log.Println("[InsertCake] can't insert cake, err:", err.Error())
		return m.Cake{}, nil
	}

	id, _ := rows.LastInsertId()
	cake.Id = int(id)
	cake.CreatedAt = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	cake.UpdatedAt = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	return cake, nil
}

func (r *repository) UpdateCake(ctx context.Context, cake m.Cake) (m.Cake, error) {
	var (
		err        error
		cakeTemp   m.Cake
		created_at string
		updated_at string
	)
	currentTime := time.Now().String()

	query := fmt.Sprintf(database.GetDetailsOfCakeByID, cake.Id)
	err = r.db.QueryRow(query).Scan(&cakeTemp.Id, &cakeTemp.Title, &cakeTemp.Description, &cakeTemp.Rating, &cakeTemp.Image, &cakeTemp.CreatedAt, &cakeTemp.UpdatedAt)
	if err == sql.ErrNoRows {
		log.Println("[UpdateCake] can't update cake, err:", err.Error())
		return m.Cake{}, ErrNotFound
	} else if err != nil {
		log.Println("[UpdateCake] can't update cake, err:", err.Error())
		return m.Cake{}, nil
	}

	if cake.Title == "" {
		cake.Title = cakeTemp.Title
	}
	if cake.Description == "" {
		cake.Description = cakeTemp.Description
	}
	if cake.Rating == -9999.9999 {
		cake.Rating = cakeTemp.Rating
	}
	if cake.Image == "" {
		cake.Image = cakeTemp.Image
	}

	created_at = cakeTemp.UpdatedAt
	updated_at = currentTime

	query = fmt.Sprintf(database.UpdateCakeByID, cake.Title, cake.Description, cake.Rating, cake.Image, created_at, updated_at, cake.Id)
	rows, err := r.db.Exec(query)
	if err != nil {
		log.Println("[UpdateCake] can't update cake, err:", err.Error())
		return m.Cake{}, nil
	}

	cake.CreatedAt = created_at
	cake.UpdatedAt = fmt.Sprintf("%d-%d-%d %d:%d:%d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		return cake, nil
	}

	return cake, nil
}

func (r *repository) DeleteCake(ctx context.Context, id int) (err error) {
	query := fmt.Sprintf(database.DeleteCakeByID, id)
	rows, err := r.db.Exec(query)
	if err != nil {
		log.Println("[DeleteCake] can't delete cake, err:", err.Error())
		return err
	}

	rowsAffected, _ := rows.RowsAffected()
	if rowsAffected > 0 {
		return nil
	}

	log.Println("[DeleteCake] can't delete cake, err:", ErrNotFound.Error())
	return ErrNotFound
}
