package entity

type WorkInfoEntity struct {
	UserID        string `bson:"user_id"`
	Team          string `bson:"team"`
	Position      string `bson:"position"`
	DefaultShift  string `bson:"default_shift"`
	WeekdayOff    string `bson:"weekday_off"`
	WeekendDayOff string `bson:"weekend_day_off"`
	SuperiorID    string `bson:"superior_id"`
}
