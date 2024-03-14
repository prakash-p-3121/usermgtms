package impl

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/mysqllib"
	"github.com/prakash-p-3121/usermodel"
	"golang.org/x/crypto/bcrypt"
	"sync"
	"time"
)

type UserRepositoryImpl struct {
	ShardConnectionsMap   *sync.Map
	SingleStoreConnection *sql.DB
}

func (repository *UserRepositoryImpl) UserCreate(shardID int64,
	idGenResp *idgenmodel.IDGenResp,
	req *usermodel.UserCreateReq) errorlib.AppError {
	db, err := mysqllib.RetrieveShardConnectionByShardID(repository.ShardConnectionsMap, shardID)
	if err != nil {
		return errorlib.NewBadReqError(err.Error())
	}

	tx, err := db.BeginTx(context.Background(), nil)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}
	appErr := repository.userCreate(tx, shardID, idGenResp, req)
	if appErr != nil {
		return errorlib.NewInternalServerError(mysqllib.RollbackTx(tx, err).Error())
	}
	err = tx.Commit()
	if err != nil {
		return errorlib.NewInternalServerError(mysqllib.RollbackTx(tx, err).Error())
	}
	return nil
}

func (repository *UserRepositoryImpl) userCreate(tx *sql.Tx, shardID int64,
	idGenResp *idgenmodel.IDGenResp,
	req *usermodel.UserCreateReq) errorlib.AppError {

	createdAt := time.Now().UTC()
	qry := `INSERT INTO users (
                   id, 
                   id_bit_count, 
                   first_name, 
                   last_name, 
                   email_id, 
                   coutry_code, 
                   phone_number, 
                   created_at, 
                   updated_at
                   ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?) ;`
	_, err := tx.Exec(qry, idGenResp.ID,
		idGenResp.BitCount,
		req.FirstName,
		req.LastName,
		req.EmailID,
		req.CountryCode,
		req.PhoneNumberStr,
		createdAt,
		createdAt,
	)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}

	hashedPassword, err := repository.hashPassword(*req.Password)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}

	qry = `INSERT INTO passwords (user_id, hashed_password, updated_at) VALUES (?, ?, ?);`
	_, err = tx.Exec(qry, idGenResp.ID, hashedPassword, createdAt)
	if err != nil {
		return errorlib.NewInternalServerError(err.Error())
	}

	return nil
}

func (repository *UserRepositoryImpl) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error generating hashed password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

func (repository *UserRepositoryImpl) validatePassword(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
