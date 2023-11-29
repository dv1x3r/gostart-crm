package sqlitedb

import (
	"fmt"

	"w2go/internal/app/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Todo struct {
	db *sqlx.DB
}

func NewTodo(db *sqlx.DB) *Todo {
	return &Todo{db: db}
}

func (ts *Todo) SelectList(p model.QueryListParams) ([]model.TodoDTO, int64, error) {
	sb := sqlbuilder.NewSelectBuilder().
		Select("*").
		From("todo").
		Limit(p.Limit).
		Offset(p.Offset)

	for _, w := range p.Where {
		switch w.Operator {
		case "=":
			sb = sb.Where(sb.EQ(w.Field, w.Value))
		case ">":
			sb = sb.Where(sb.GT(w.Field, w.Value))
		case "<":
			sb = sb.Where(sb.LT(w.Field, w.Value))
		case ">=":
			sb = sb.Where(sb.GTE(w.Field, w.Value))
		case "<=":
			sb = sb.Where(sb.LTE(w.Field, w.Value))
		case "begins":
			sb = sb.Where(sb.Like(w.Field, fmt.Sprintf("%v%%", w.Value)))
		case "contains":
			sb = sb.Where(sb.Like(w.Field, fmt.Sprintf("%%%v%%", w.Value)))
		case "ends":
			sb = sb.Where(sb.Like(w.Field, fmt.Sprintf("%%%v", w.Value)))
		case "between:":
			if values, ok := w.Value.([]any); ok && len(values) == 2 {
				sb = sb.Where(sb.Between(w.Field, values[0], values[1]))
			}
		}
	}

	// for _, s := range q.Sort {
	// 	sb = sb.OrderBy(fmt.Sprint(s.Field, s.Direction))
	// }
	// fmt.Printf("%+v \n", q)
	fmt.Println(sb.String())
	return nil, 0, nil
}

func (ts *Todo) SelectByID(id int64) {
	_ = sqlbuilder.NewSelectBuilder()
}

func (ts *Todo) DeleteByID(id []int64) {
	_ = sqlbuilder.NewDeleteBuilder()
}

func (ts *Todo) Insert() {
	_ = sqlbuilder.NewInsertBuilder()
}

func (ts *Todo) Update() {
	_ = sqlbuilder.NewUpdateBuilder()
}

func (ts *Todo) Patch() {
	_ = sqlbuilder.NewUpdateBuilder()
}
