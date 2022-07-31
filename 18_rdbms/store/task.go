package store

import (
	"context"

	"github.com/budougumi0617/go_todo_app/entity"
)

//DBに対する操作

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	//時間設定
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `INSERT INTO task
		(title, status, created, modified)
	VALUES (?, ?, ?, ?)`
	//保存したidを取得できるresult
	result, err := db.ExecContext(
		ctx, sql, t.Title, t.Status,
		t.Created, t.Modified,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	//呼び出し元にIDを返す
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
			id, title,
			status, created, modified
		FROM task;`
	//複数行取得
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}
