package handler

import (
	"net/http"

	"github.com/budougumi0617/go_todo_app/entity"
	"github.com/budougumi0617/go_todo_app/store"
)

//タスクを一覧するエンドポイント
//addtask同様

type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID `json:"id"`
	Title  string        `json:"title"`
	Status string        `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	//データをstoreから取得
	tasks := lt.Store.All()
	rsp := []task{}
	//余分なデータを取り除く為、詰め直す
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
