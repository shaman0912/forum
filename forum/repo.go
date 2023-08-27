package forum

import "01.alem.school/git/atastemi/forum/forum/domain"

type Repo interface {
	GetUser(username string) (domain.User, error)
	SaveSession(session domain.Session) error
	GetSession(sessionID string) (domain.Session, error)
	SaveUser(domain.User) error
	SavePosts(domain.Posts) error
	GetPosts() ([]domain.Posts, error)
	GetUserPosts(userId int) ([]domain.Posts, error)
	DeletePost(postId int) error
	DeleteComment(comment_id int) error
	GetPostByID(postID int) (domain.Posts, error)
	AddComment(domain.Comments) error
	GetComments(postId int) ([]domain.Comments, error)
	GetUserById(userId int) ([]domain.User, error)
	LikePost(postID, userID int) error
	DislikePost(postID, userID int) error
	GetLikedPostIDs(userID int) ([]int, error)
	GetPostsByCategories(categories []string) ([]domain.Posts, error)
	GetUserByEmail(email string) (domain.User, error)
	LikeComment(commentID int, userID int) error
	DislikeComment(commentID int, userID int) error
	InvalidateSessions(userID int) error
}
