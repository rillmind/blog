package posts

import "gorm.io/gorm"

type Service struct {
	connection *gorm.DB
}

func NewService(connection *gorm.DB) Service {
	return Service{
		connection,
	}
}

func (ps *Service) CreatePost(post Model) (Model, error) {
	query := ps.connection.Create(&post)
	if query.Error != nil {
		return Model{}, query.Error
	}

	return post, nil
}

func (ps *Service) ReadPosts() ([]Model, error) {
	posts := []Model{}

	query := ps.connection.Find(&posts)
	if query.Error != nil {
		return []Model{}, query.Error
	}

	return posts, nil
}
