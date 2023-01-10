package repository

import (
	"exacta/backend/model/domain"
	"time"
	"database/sql"
)

type UserRepository interface{
	FetchUserByID(id uint) (domain.UserDomain, error)
	FetchUsers() ([]domain.UserDomain, error)
	InsertUser(username string, email string, password string, nama_sekolah string) (*string, error)
	LoginUser(email string, password string) (*uint, error)
	GetPasswordCompare(email string) (*string, error)
	FetchUserIdByEmail(email string) (*int, error)
	PushToken(user_id uint, token string, expired_at time.Time) (*string, error)
	DeleteToken(id uint) (bool, error)
}

type QuizRepository interface{
	FindCategories() ([]domain.CategoryDomain, error)
	FindCategoryById(categoryId uint) (domain.CategoryDomain, error)
	FindQuizByCategoryIdWithPagination(categoryId, page, limit uint) ([]domain.QuizDomain, error)
	FindIncorrectAnswersByQuizId(quizId uint) (domain.IncorrectAnswerDomain, error)
}

type Repository struct {
	UserRepository
	QuizRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		QuizRepository:   NewQuizRepository(db),
	}
}