package routes

import (
	"net/http"
	"time"

	"github.com/code-golf/code-golf/pager"
	"github.com/code-golf/code-golf/session"
)

// RecentGolfers serves GET /recent/golfers
func RecentGolfers(w http.ResponseWriter, r *http.Request) {
	type golfer struct {
		Country, Name string
		Date          time.Time
	}

	golfers := make([]golfer, 0, pager.PerPage)

	rows, err := session.Database(r).Query(
		`WITH recent AS (
		    SELECT user_id, MIN(earned) date
		      FROM trophies
		  GROUP BY user_id
		  ORDER BY date DESC
		     LIMIT $1
		) SELECT COALESCE(CASE WHEN show_country THEN country END, ''),
		         date, login
		    FROM recent
		    JOIN users ON id = user_id
		ORDER BY date DESC`,
		pager.PerPage,
	)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var g golfer

		if err := rows.Scan(&g.Country, &g.Date, &g.Name); err != nil {
			panic(err)
		}

		golfers = append(golfers, g)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	render(w, r, "recent/golfers", golfers, "Recent Golfers")
}
