package login

import (
	"context"
	"database/sql"
	"encoding/json"
	server "project_yd/server"
	request "project_yd/server/server_packet/request_packet"
	response "project_yd/server/server_packet/response_packet"
	"project_yd/util"

	"github.com/google/uuid"
)

func RegistLoginRpc() {
	server.RegistRpc("login", LoginRpc)
}

func LoginRpc(payload string) string {
	requestPacket := request.Login{}
	err := json.Unmarshal([]byte(payload), &requestPacket)
	if err != nil {
		return util.ResponseErrorMessage(util.BadRequest, err.Error())
	}

	var userId string
	ctx := context.Background()
	//-- 로그인 정보 가져오기
	db := server.DBManager.Database[util.LOGIN]
	loginSql := `SELECT uid FROM account WHERE user_id = ?`
	err = db.QueryRowContext(ctx, loginSql, requestPacket.Id).Scan(&userId)

	//-- 등록된 정보가 없으므로 신규 유저로 처리
	responsePacket := response.Login{}
	if err == sql.ErrNoRows {
		//-- uuid 생성
		uid := CreateNameUUID(requestPacket.Id)
		createUUIDSql := `INSERT INTO account (uid, user_id) VALUES(?,?)`
		_, err := db.ExecContext(ctx, createUUIDSql, uid, requestPacket.Id)
		if err != nil {
			return util.ResponseErrorMessage(util.ServerError, err.Error())
		}

		responsePacket.UUID = uid
		responsePacket.Message = "Success"
		responsePacket.MessageCode = util.Success

		return util.ResponseMessage(responsePacket)
	}

	responsePacket.UUID = userId
	responsePacket.MessageCode = 200
	responsePacket.Message = "Sueccess"

	return util.ResponseMessage(responsePacket)
}

// -- 랜덤 uuid 생성
func CreateRandomUUID() string {
	uid := uuid.New().String()
	println("CreatRandomUUID:: uuid:", uid)
	return uid
}

// -- 이름기반 uuid 생성
func CreateNameUUID(name string) string {
	baseUUID := uuid.New()
	nameByte := []byte(name)
	nameUUID := uuid.NewSHA1(baseUUID, nameByte)
	println("CreateNameUUID:: name:", name, "/uuid:", nameUUID.String())
	return nameUUID.String()
}
