package blog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuiexample/database"
)

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	// Set the header of the response as application/json [応答のヘッダーを application/json として設定します]
	w.Header().Set("Content-Type", "application/json")

	// initialize response variable with ResponseModel data structure
	// ResponseModel データ構造で応答変数を初期化する
	var response database.ResponseModel

	// read mux parameter
	params := mux.Vars(r)
	// get the id parameter
	id := params["id"]

	// prepare the delete query
	stmt, err := database.DB.Prepare("DELETE FROM blogs WHERE id = ?")
	if err != nil {
		response.Message = "unexpected database error"
		json.NewEncoder(w).Encode(response)
		return
	}

	// execute the query [クエリを実行する]
	_, err = stmt.Exec(id)
	if err != nil {
		response.Message = "failed to delete data"
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = fmt.Sprintf("blog with ID: %s was deleted", id)
	// Return API response as JSON [API レスポンスを JSON として返す]
	json.NewEncoder(w).Encode(response)
}
