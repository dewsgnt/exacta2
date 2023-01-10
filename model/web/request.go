package web

import (
	"errors"
	"net"
	"net/mail"
	"strings"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NamaSekolah string `json:"nama_sekolah"`
}

type AnswerAttemptRequest struct {
	Answers    []DataAttemptRequest `json:"answers"`
	CategoryId uint                 `json:"category_id"`
	Duration   string               `json:"duration"`
}

type DataAttemptRequest struct {
	QuizId uint   `json:"quiz_id"`
	Answer string `json:"answer"`
}

var (
	ErrRequiredUsername   = errors.New("required username")
	ErrRequiredEmail      = errors.New("required email")
	ErrRequiredPassword   = errors.New("required password")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrDomainNotFound     = errors.New("domain not found")
	ErrRequiredAnswer     = errors.New("required Answer")
	ErrRequiredQuizId     = errors.New("required Quiz Id")
	ErrRequiredCategoryId = errors.New("required Category Id")
	ErrRequiredDuration   = errors.New("required Duration")
)

func (l *LoginRequest) ValidateLogin() error {
	if err := validateEmail(l.Email); err != nil {
		return err
	}
	if l.Password == "" {
		return ErrRequiredPassword
	}

	return nil
}

func (r *RegisterRequest) ValidateRegister() error {
	if r.Username == "" {
		return ErrRequiredUsername
	}
	if err := validateEmail(r.Email); err != nil {
		return err
	}
	if r.Password == "" {
		return ErrRequiredPassword
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return ErrRequiredEmail
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return ErrInvalidEmail
	}

	parts := strings.Split(email, "@")

	_, err = net.LookupMX(parts[1])
	if err != nil {
		return ErrDomainNotFound
	}

	return nil
}

func (a *AnswerAttemptRequest) ValidateAnswerAttempt() error {
	if a.CategoryId == 0 {
		return ErrRequiredCategoryId
	}
	if a.Duration == "" {
		return ErrRequiredDuration
	}
	for _, v := range a.Answers {
		if err := v.ValidateDataAttempt(); err != nil {
			return err
		}
	}

	return nil
}

func (d *DataAttemptRequest) ValidateDataAttempt() error {
	if d.Answer == "" {
		return ErrRequiredAnswer
	}
	if d.QuizId == 0 {
		return ErrRequiredQuizId
	}

	return nil
}
