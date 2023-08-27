package business

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"01.alem.school/git/atastemi/forum/forum"
	"01.alem.school/git/atastemi/forum/forum/domain"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Business struct {
	repo forum.Repo
}

func NewBusiness(repo forum.Repo) (*Business, error) {
	return &Business{
		repo: repo,
	}, nil
}

func (b *Business) Login(username, password string) (uuid.UUID, error) {
	user, err := b.repo.GetUser(username)
	if err != nil {
		return uuid.Nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return uuid.Nil, err
	}

	// Invalidate existing sessions for the user
	err = b.repo.InvalidateSessions(user.UserId)
	if err != nil {
		return uuid.Nil, err
	}

	sessionID, err := uuid.NewV4()
	if err != nil {
		return uuid.Nil, err
	}

	session := domain.Session{
		UserId:         user.UserId,
		Username:       user.Username,
		SessionId:      sessionID.String(),
		CreationDate:   time.Now(),
		ExpiritionDate: time.Now().Add(time.Hour * 24),
	}
	err = b.repo.SaveSession(session)
	return sessionID, err
}

func (b *Business) Registration(username, password, email string) error {
	if len(username) < 3 || len(username) > 15 {
		return domain.ErrInvalidDataonRegistartion
	}
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatch, err := regexp.MatchString(emailPattern, email)
	if err != nil {
		return domain.ErrInvalidDataonRegistartion
	}
	if !emailMatch {
		return domain.ErrInvalidDataonRegistartion
	}
	if len(password) < 6 || len(password) > 30 {
		return domain.ErrInvalidDataonRegistartion
	}
	_, err = b.repo.GetUserByEmail(email)
	if err == nil {
		return domain.ErrUserAlreadyExist
	} else if !errors.Is(err, domain.ErrInvalidUser) {
		return err
	}
	_, err = b.repo.GetUser(username)
	if err == nil {
		return domain.ErrUserAlreadyExist
	} else if !errors.Is(err, domain.ErrInvalidUser) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	newUser := domain.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = b.repo.SaveUser(newUser)
	return err
}

func (b *Business) Session(sessionID string) (*domain.Session, error) {
	session, err := b.repo.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (b *Business) Post(posts domain.Posts) error {
	err := b.repo.SavePosts(posts)
	return err
}

func (b *Business) GetAllPosts() ([]domain.Posts, error) {
	posts, err := b.repo.GetPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (b *Business) GetMyPosts(userId int) ([]domain.Posts, error) {
	posts, err := b.repo.GetUserPosts(userId)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (b *Business) DeletePost(postId int) error {
	err := b.repo.DeletePost(postId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (b *Business) GetPostByID(postId int) (domain.Posts, error) {
	post, err := b.repo.GetPostByID(postId)
	if err != nil {
		fmt.Println(err)
		return post, err
	}

	return post, nil
}

func (b *Business) AddComment(comment domain.Comments) error {
	err := b.repo.AddComment(comment)
	return err
}

func (b *Business) GetComments(postId int) ([]domain.Comments, error) {
	comments, err := b.repo.GetComments(postId)
	if err != nil {
		fmt.Println(err)
		return comments, err
	}

	return comments, nil
}

func (b *Business) DeleteComment(comment_id int) error {
	err := b.repo.DeleteComment(comment_id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (b *Business) GetUserById(userId int) ([]domain.User, error) {
	users, err := b.repo.GetUserById(userId)
	if err != nil {
		fmt.Println(err)
		return users, err
	}
	return users, nil
}

func (b *Business) LikePost(postID, userID int) error {
	err := b.repo.LikePost(postID, userID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (b *Business) DislikePost(postID, userID int) error {
	err := b.repo.DislikePost(postID, userID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (b *Business) GetLikedPosts(userID int) ([]domain.Posts, error) {
	likedPostIDs, err := b.repo.GetLikedPostIDs(userID)
	if err != nil {
		return nil, err
	}

	var likedPosts []domain.Posts
	for _, postID := range likedPostIDs {
		fmt.Println(postID)
		post, err := b.repo.GetPostByID(postID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		likedPosts = append(likedPosts, post)
	}

	return likedPosts, nil
}

func (b *Business) GetPostsByCategories(categories []string) ([]domain.Posts, error) {
	posts, err := b.repo.GetPostsByCategories(categories)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (b *Business) LikeComment(commentID int, userID int) error {
	return b.repo.LikeComment(commentID, userID)
}

func (b *Business) DislikeComment(commentID int, userID int) error {
	return b.repo.DislikeComment(commentID, userID)
}
