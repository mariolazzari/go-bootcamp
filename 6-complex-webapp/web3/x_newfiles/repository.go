package repository

import "web3/models"

type DatabaseRepo interface {
	// 26 Define function for entering new
	// posts
	InsertPost(newPost models.Post) error

	// 27 Gets user data with ID
	GetUserByID(id int) (models.User, error)

	// 27 Used to authenticate login
	AuthenticateUser(email, testPassword string) (int, string, error)

	// 29 Returns one article from the DB
	GetAnArticle() (int, int, string, string, error)

	// 29 Get 3 posts
	Get3Articles() (models.ArticleList, error)
}
