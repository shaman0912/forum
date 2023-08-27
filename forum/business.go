package forum

import (
	"01.alem.school/git/atastemi/forum/forum/domain"
	"github.com/gofrs/uuid"
)

type Business interface {
	Login(username, password string) (uuid.UUID, error)
	Registration(username, password, email string) error
	Session(sessionID string) (*domain.Session, error)
	Post(domain.Posts) error
	GetAllPosts() ([]domain.Posts, error)
	GetMyPosts(userId int) ([]domain.Posts, error)
	DeletePost(postId int) error
	GetPostByID(postId int) (domain.Posts, error)
	AddComment(comment domain.Comments) error
	GetComments(postId int) ([]domain.Comments, error)
	DeleteComment(comment_id int) error
	GetUserById(userId int) ([]domain.User, error)
	LikePost(postID, userID int) error
	DislikePost(postID, userID int) error
	GetLikedPosts(userID int) ([]domain.Posts, error)
	GetPostsByCategories(categories []string) ([]domain.Posts, error)
	LikeComment(commentID int, userID int) error
	DislikeComment(commentID int, userID int) error
}
