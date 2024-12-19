package timeutil

import "time"

func GetFirstAndLastDayOfMonth(date, layout string) (time.Time, time.Time, error) {
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	year, month, _ := t.Date()
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
	lastDay := firstDay.AddDate(0, 1, -1)
	return firstDay, lastDay, nil
}
