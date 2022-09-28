package blog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/yuiexample/database"
)

func List(w http.ResponseWriter, r *http.Request) {
	// Set the header of the response as application/json [応答のヘッダーを application/json として設定します]
	w.Header().Set("Content-Type", "application/json")

	// initialize response variable with ResponseModel data structure
	// ResponseModel データ構造で応答変数を初期化する
	var response database.ResponseModel

	// initialize blogs variable with array BlogModel data structure
	// 配列 BlogModel データ構造でブログ変数を初期化する
	var blogs []database.BlogModel

	// execute the query [クエリを実行する]
	result, err := database.DB.Query("SELECT id, name, post_id FROM blogs")
	if err != nil {
		response.Message = "unexpected database error"
		json.NewEncoder(w).Encode(response)
		return
	}
	// close the result body [結果の本文を閉じます]
	defer result.Close()

	// looping the result data [結果データのループ]
	for result.Next() {
		// initialize blog variable with BlogModel data structure
		// BlogModel データ構造でブログ変数を初期化する
		var blog database.BlogModel

		// parsing the data [データの解析]
		err := result.Scan(&blog.ID, &blog.Name, &blog.PostID)
		if err != nil {
			response.Message = "failed to read data"
			json.NewEncoder(w).Encode(response)
			return
		}

		// insert blog data to array blogs data [ブログ データを配列ブログ データに挿入する]
		blogs = append(blogs, blog)
	}

	// looping the blogs data [ブログデータのループ]
	for i, _ := range blogs {
		// Call the POST API to get the detail of post data [POST API を呼び出して投稿データの詳細を取得する]
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(blogs[i].PostID))
		if err != nil {
			response.Message = "failed to get post from API"
			json.NewEncoder(w).Encode(response)
			return
		}
		// close the response body [応答本文を閉じます]
		defer resp.Body.Close()

		// read the API body [API 本体を読む]
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			response.Message = "failed to get API body"
			json.NewEncoder(w).Encode(response)
			return
		}

		// initialize post variable with PostModel data structure
		// PostModel データ構造でポスト変数を初期化する
		var post database.PostModel
		// parsing the data [データの解析]
		json.Unmarshal(body, &post)

		// insert post data to the blog post data [投稿データをブログ投稿データに挿入する]
		blogs[i].Post = post
	}

	// Return API response as JSON [API レスポンスを JSON として返す]
	json.NewEncoder(w).Encode(blogs)
}
