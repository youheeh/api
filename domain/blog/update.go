package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/yuiexample/database"
)

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	// Set the header of the response as application/json [応答のヘッダーを application/json として設定します]
	w.Header().Set("Content-Type", "application/json")

	// initialize response variable with ResponseModel data structure
	// ResponseModel データ構造で応答変数を初期化する
	var response database.ResponseModel

	// read mux parameter [mux パラメータの読み取り]
	params := mux.Vars(r)
	// get the id parameter [id パラメータを取得する]
	id := params["id"]

	// read request JSON body sent by user [ユーザーが送信した読み取りリクエストの JSON 本文]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Message = "failed ready body"
		json.NewEncoder(w).Encode(response)
		return
	}

	// parsing the data [データの解析]
	keyVal := make(map[string]interface{})
	json.Unmarshal(body, &keyVal)

	// get name field value [nameフィールドの値を取得]
	name := keyVal["name"]
	// get post_id field value [post_id フィールドの値を取得する]
	postID := keyVal["post_id"]

	// prepare update query [更新クエリを準備する]
	stmt, err := database.DB.Prepare("UPDATE blogs SET name = ?, post_id = ? WHERE id = ?")
	if err != nil {
		response.Message = "unexpected database error"
		json.NewEncoder(w).Encode(response)
		return
	}

	// execute the query [クエリを実行する]
	_, err = stmt.Exec(name, postID, id)
	if err != nil {
		response.Message = "failed update data"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = fmt.Sprintf("blog with ID: %s has been updated", id)
	// Return API response as JSON [API レスポンスを JSON として返す]
	json.NewEncoder(w).Encode(response)
}
