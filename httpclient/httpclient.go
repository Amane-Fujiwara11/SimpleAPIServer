package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// URLを取得するための関数
func FetchURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// エラーハンドリング: リクエストエラーの表示
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close() // 関数が終了する際にレスポンスボディを閉じる

	// ステータスコードの確認
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received response code %d\n", resp.StatusCode)
		return
	}

	// レスポンスの読み取り
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// エラーハンドリング: レスポンス読み取りエラー
		fmt.Println("Error reading response body:", err)
		return
	}

	// レスポンスボディの内容を表示
	fmt.Println("Response Body:", string(body))
}
