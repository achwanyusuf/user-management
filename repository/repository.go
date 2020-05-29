package repository

import(
	"github.com/achwanyusuf/user-management/model"
	"database/sql"
	"strconv"
	_ "github.com/lib/pq"
	"github.com/achwanyusuf/user-management/proto"
)

type UserRepository interface {
	Create(*model.UserData) bool
	Update(*model.UserData) bool
	Update2(*model.UserData) bool
	Delete(string) bool
	ReadOneByEmail(string) (*model.UserData,error)
	ReadOneByUserId(string) (*model.UserData,error)
	ReadToken(string) bool
	ReadAll() (*proto.ReadAllResponse, error)
}

type repositoryPostgres struct{
	db *sql.DB
}

func NewRepositoryPostgres(db *sql.DB) *repositoryPostgres{
	return &repositoryPostgres{db}
}

func (rp *repositoryPostgres) Delete(userId string) (bool){
	statement, err := rp.db.Prepare(`DELETE FROM users WHERE UserId=$1`)
	if err != nil {
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(userId)
	if err != nil {
		return false
	}
	return true
}

func (rp *repositoryPostgres) Create(data *model.UserData) (bool){
	statement, err := rp.db.Prepare(`INSERT INTO users (userid, email, password, address) VALUES ($1, $2, $3, $4)`)
	if err != nil {
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(data.UserId, data.Email, data.Password, data.Address)
	if err != nil {
		return false
	}
	return true
}

func (rp *repositoryPostgres) ReadToken(token string) (bool){
	var row string
	statement, err := rp.db.Prepare(`SELECT
		COUNT(id) as rows 
		FROM users where token = $1 LIMIT 1`)
	if err != nil {
		return false
	}
	defer statement.Close()
	statement.QueryRow(token).Scan(&row)
	a, err := strconv.ParseUint(row, 10, 64)
	if err != nil {
		return false
	}
	if a < 1{
		return false
	}
	return true

}

func (rp *repositoryPostgres) Update(data *model.UserData) (bool){
	statement, err := rp.db.Prepare(`UPDATE users SET email = $1, password = $2, address = $3, latestlogin = $4, token = $5
	where userid = $6`)
	if err != nil {
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(data.Email, data.Password, data.Address, data.LatestLogin, data.Token, data.UserId)
	if err != nil {
		return false
	}
	return true
}

func (rp *repositoryPostgres) Update2(data *model.UserData) (bool){
	statement, err := rp.db.Prepare(`UPDATE users SET email = $1, password = $2, address = $3
	where userid = $4`)
	if err != nil {
		return false
	}
	defer statement.Close()
	_, err = statement.Exec(data.Email, data.Password, data.Address, data.UserId)
	if err != nil {
		return false
	}
	return true
}


func (rp *repositoryPostgres) ReadOneByEmail(input string) (*model.UserData,error){
	var userData model.UserData
	statement, err := rp.db.Prepare(`SELECT
		userid, email, password, address
		FROM users where email = $1 LIMIT 1`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows := statement.QueryRow(input).Scan(&userData.UserId, &userData.Email, &userData.Password, &userData.Address)
	if rows != nil {
		return nil, rows
	}
	return &userData, nil
}

func (rp *repositoryPostgres) ReadOneByUserId(input string) (*model.UserData,error){
	var userData model.UserData
	statement, err := rp.db.Prepare(`SELECT
		userid, email, password, address, token, latestlogin
		FROM users where userid = $1 LIMIT 1`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	statement.QueryRow(input).Scan(&userData.UserId, &userData.Email, &userData.Password, &userData.Address, &userData.Token, &userData.LatestLogin )
	return &userData, nil
}

func (rp *repositoryPostgres) ReadAll() (*proto.ReadAllResponse,error){
	rows, _ := rp.db.Query(`SELECT
		userid, email, address, latestlogin
		FROM users`)
	defer rows.Close()
	var users []*proto.ReadOneResponse
	for rows.Next(){
		var userData proto.ReadOneResponse
		rows.Scan(&userData.UserId, &userData.Email, &userData.Address, &userData.LatestLogin)
		users = append(users, &userData)
	}
	return &proto.ReadAllResponse{
		Users: users,
	}, nil
}