package repository
import(
	"database/sql"
	"exacta/backend/model/domain"
	"errors"
	"time"	
)

type UserRepositoryImpl struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db,
	}
}

func (u *UserRepositoryImpl) FetchUserByID(id uint) (domain.UserDomain, error) {
	user := domain.UserDomain{}
	sqlStatement := `SELECT user_id, username, email, nama_sekolah FROM users WHERE user_id = ?`

	row := u.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.NamaSekolah,
		
	)
	if err != nil {
		return user, err
	}
	return user, nil
	// defer row.Close()
}

func (u *UserRepositoryImpl) FetchUsers() ([]domain.UserDomain, error) {
	var users []domain.UserDomain

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.UserDomain

		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.NamaSekolah,
			&user.Password,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)


	}
	return users, nil
}

func (u *UserRepositoryImpl) InsertUser(username string, email string, password string, nama_sekolah string) (*string, error) {
	users, err := u.FetchUsers()
	//fmt.Println("isi users",users)

	for _, value := range users {
		if value.Email == email {
			return nil, err
		}
	}
	_, err = u.db.Exec("INSERT INTO users (username, email, nama_sekolah,password) VALUES (?,?,?,?)", username, email, nama_sekolah,password)

	if err != nil {
		return nil, err
	}

	return &email, err
}

func (u *UserRepositoryImpl) LoginUser(email string, password string) (*uint, error) {
	users, err := u.FetchUsers()
	if err != nil {
		return nil, errors.New("login failed")
	}

	for _, user := range users {
		if user.Email == email && user.Password == password {
			return &user.Id, nil
		}

	}
	return nil, errors.New("login failed")
}

func (u *UserRepositoryImpl) GetPasswordCompare(email string) (*string, error) {
	var pass string
	sqlStatement := `SELECT password FROM users WHERE email = ?`

	row := u.db.QueryRow(sqlStatement, email)
	err := row.Scan(&pass)
	if err != nil {
		return nil, err
	}
	return &pass, err
}
func (u *UserRepositoryImpl) FetchUserIdByEmail(email string) (*int, error) {
	var user_id int

	rows, err := u.db.Query("SELECT id FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user_id)
		if err != nil {
			return nil, err
		}
	}
	return &user_id, nil
}

func (u *UserRepositoryImpl) PushToken(user_id uint, token string, expired_at time.Time) (*string, error) {
	_, err := u.db.Exec("INSERT INTO auth (user_id, token, expired_at) VALUES (?, ?,?)",
		user_id, token, expired_at)

	if err != nil {
		return nil, err
	}
	return &token, err
}

func (u *UserRepositoryImpl) DeleteToken(id uint) (bool, error) {
	_, err := u.db.Exec("DELETE FROM auth WHERE user_id = ?", id)

	if err != nil {
		return false, err
	}
	return true, err
}

func (u *UserRepositoryImpl) GetUserIDByToken(token string) (uint, error) {
	var user_id uint

	sqlStatement := `SELECT user_id FROM auth WHERE token = ?`

	row := u.db.QueryRow(sqlStatement, token)
	err := row.Scan(&user_id)
	if err != nil {
		return user_id, err
	}
	return user_id, nil

}