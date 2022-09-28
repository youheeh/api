package blog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yuiexample/database"
)

func Create(w http.ResponseWriter, r *http.Request) {
	// Set the header of the response as application/json [応答のヘッダーを application/json として設定します]
	w.Header().Set("Content-Type", "application/json")

	// initialize response variable with ResponseModel data structure
	// ResponseModel データ構造で応答変数を初期化する
	var response database.ResponseModel

	// read request JSON body sent by user [ユーザーが送信した読み取りリクエストの JSON 本文]
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Message = "failed read body"
		json.NewEncoder(w).Encode(response)
		return
	}
	//fmt.Println(1)
	// parsing the data [データの解析]
	keyVal := make(map[string]interface{})
	json.Unmarshal(body, &keyVal)

	// get name field value [nameフィールドの値を取得]
	name := keyVal["name"]
	// get post_id field value [post_id フィールドの値を取得する]
	postID := keyVal["post_id"]

	// preparing insert query [挿入クエリの準備]
	stmt, err := database.DB.Prepare("INSERT INTO blogs(name, post_id) VALUES(?, ?)")
	if err != nil {
		response.Message = "unexpected database error"
		json.NewEncoder(w).Encode(response)
		return
	}

	// execute the query [クエリを実行する]
	_, err = stmt.Exec(name, postID)
	if err != nil {
		fmt.Print(err.Error())
		response.Message = "failed insert data"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "blog has been created"
	// Return API response as JSON [API レスポンスを JSON として返す]
	json.NewEncoder(w).Encode(response)
}
