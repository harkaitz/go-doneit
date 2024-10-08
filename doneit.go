package doneit

import (
	"gorm.io/gorm"
	"time"
	"fmt"
)

type Task struct {
	ID		string `gorm:"primaryKey"`
	Done		bool
	DoneDate	time.Time
}

func InitDatabase(gdb *gorm.DB) (err error) {
	return gdb.AutoMigrate(
		&Task{},
	)
}

func OnlyOnce(gdb *gorm.DB, prog func() (error), format string, a ...any) (err error) {
	var tasks	[]Task
	var shouldDo	bool
	var id		string = fmt.Sprintf(format, a...)

	err = gdb.Where("id = ?", id).Find(&tasks).Error
	if err != nil { return }

	switch {
	case len(tasks) == 0 || !tasks[0].Done:
		shouldDo = true
	case tasks[0].Done:
		shouldDo = false
	}

	if shouldDo {
		err = prog()
		if err != nil { return }
		err = gdb.Create(
			&Task{
				ID: id,
				Done: true,
				DoneDate: time.Now(),
			},
		).Error
		if err != nil { return }
	}

	return
}

