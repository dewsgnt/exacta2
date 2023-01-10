package repository

import (
	"database/sql"
	"exacta/backend/model/domain"
)

type QuizRepositoryImpl struct{
	db *sql.DB
}

func NewQuizRepository (db *sql.DB) *QuizRepositoryImpl {
	return &QuizRepositoryImpl{
		db : db,
	}
}

func (q *QuizRepositoryImpl) FindCategories() ([]domain.CategoryDomain, error) {
	var categories []domain.CategoryDomain

	query := `SELECT id, name, description FROM categories;`
	rows, err := q.db.Query(query)
	if err != nil {
		return categories, err
	}
	defer rows.Close()

	for rows.Next() {
		var category domain.CategoryDomain

		err := rows.Scan(
			&category.Id, 
			&category.Name, 
			&category.Description,
		)
		if err != nil {
			return categories, err
		}

		categories = append(categories, category)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (q *QuizRepositoryImpl) FindCategoryById(categoryId uint) (domain.CategoryDomain, error) {
	var category domain.CategoryDomain

	query := `SELECT id, name, description FROM categories WHERE id = ?;`
	row := q.db.QueryRow(query, categoryId)
	err := row.Scan(
		&category.Id, 
		&category.Name, 
		&category.Description)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (q *QuizRepositoryImpl) FindQuizByCategoryIdWithPagination(categoryId, page, limit uint) ([]domain.QuizDomain, error) {
	query := `
	SELECT id, category_id, question, correct_answer FROM quizzes 
	WHERE category_id = ? ORDER BY id LIMIT ? OFFSET ?;`

	rows, err := q.db.Query(query, categoryId, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []domain.QuizDomain
	for rows.Next() {
		var quiz domain.QuizDomain
		err := rows.Scan(
			&quiz.Id, 
			&quiz.CategoryId, 
			&quiz.Question, 
			&quiz.CorrectAnswer)
			if err != nil {
				return quizzes, err
			}

		quizzes = append(quizzes, quiz)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (q *QuizRepositoryImpl) FindIncorrectAnswersByQuizId(quizId uint) (domain.IncorrectAnswerDomain, error) {
	query := `
	SELECT id, quiz_id, option_one, option_two FROM incorrect_answers 
	WHERE quiz_id = ?;`

	var incorrectAnswerDomain domain.IncorrectAnswerDomain
	row := q.db.QueryRow(query, quizId)
	err := row.Scan(
		&incorrectAnswerDomain.Id, 
		&incorrectAnswerDomain.QuizId,
		&incorrectAnswerDomain.OptionOne, 
		&incorrectAnswerDomain.OptionTwo,
	)
	if err != nil {
		return incorrectAnswerDomain, err
	}

	return incorrectAnswerDomain, nil
}
func (q *QuizRepositoryImpl) SaveAnswerAttempt(userId uint, answersAttempt []domain.AnswerAttemptDomain, namaSekolah string) (bool, error) {
	query := `DELETE FROM answer_attempts;`
	_, err := q.db.Exec(query)
	if err != nil {
		return false, err
	}

	query = `INSERT INTO answer_attempts (answer, quiz_id, user_id, nama_sekolah) VALUES (?, ?, ?,?);`

	for _, answerAttempt := range answersAttempt {
		_, err := q.db.Exec(query, 
			answerAttempt.Answer, 
			answerAttempt.QuizId, 
			userId,
			namaSekolah)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (q *QuizRepositoryImpl) SaveResult(duration string, userId, categoryId uint, namaSekolah string) (bool, error) {

	query := `
	INSERT INTO results (correct, wrong, duration, user_id, category_id, nama_sekolah)
	SELECT 
	(SELECT COUNT(DISTINCT aa.answer) FROM answer_attempts AS aa 
		INNER JOIN quizzes AS q on aa.quiz_id = q.id 
			WHERE aa.answer = q.correct_answer) AS correct,
	(SELECT COUNT(DISTINCT aa.answer) FROM answer_attempts AS aa 
		INNER JOIN incorrect_answers AS ia on aa.quiz_id = ia.quiz_id 
			WHERE aa.answer = ia.option_one OR aa.answer = ia.option_two) AS wrong,
	?, ?,?, ?;`

	_, err := q.db.Exec(query, duration, userId, categoryId, namaSekolah)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (q *QuizRepositoryImpl) FindResultByCategoryId(categoryId uint, userId uint, namaSekolah string) (domain.ResultDomain, error) {
	query := `SELECT * FROM results WHERE category_id = ? ORDER BY id DESC;`

	var result domain.ResultDomain
	row := q.db.QueryRow(query, categoryId)
	
	err := row.Scan(
		&result.Id, 
		&result.Correct, 
		&result.Wrong, 
		&result.Duration,
		&result.NamaSekolah,
		&result.UserId, 
		&result.CategoryId,
		&result.CreatedAt, 
		&result.UpdatedAt,
	)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (q *QuizRepositoryImpl) FindScoresBoardByCategoryId(categoryId uint) ([]domain.ResultDomain, error) {
	query := `SELECT * FROM results WHERE category_id = ?;`

	rows, err := q.db.Query(query, categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.ResultDomain
	for rows.Next() {
		var result domain.ResultDomain
		err := rows.Scan(
			&result.Id, 
			&result.Correct, 
			&result.Wrong, 
			&result.Duration, 
			&result.NamaSekolah,
			&result.UserId,
			&result.CategoryId, 
			&result.CreatedAt, 
			&result.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}