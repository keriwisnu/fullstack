package seed

import (
		"log"

		"github.com/jinzhu/gorm"
		"github.com/keriwisnu/fullstack/api/models"
)

var users = []models.User {
	models.User{
		Nickname: "Keri Wisnu",
		Email: "keriwisnu@gmail.com",
		Password: "password",
	},
	models.User {
		Nickname: "Finite",
		Email: "finite@gmail.com",
		Password: "finite",
	},
}

var posts = []models.Post {
	models.Post{
		Title:     "How to sleep",
		Content:   "Daily life",
	},
	models.Post{
		Title:     "How to sneeze",
		Content:   "Daily life",
	},
}

func Load (db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table : %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table : %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error : %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table : %v", err)
		}

		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table :%v", err)
		}
	}
}







