package ms

import (
	"context"
	"fmt"
	"log"

	pbu "github.com/cipta-ageung/simasgo/protobuf/user"
	usrepo "github.com/cipta-ageung/simasgo/repository/user"

	"golang.org/x/crypto/bcrypt"
)

// UserService : struct interface
type UserService struct {
	repo         usrepo.Repository
	tokenService Authable
}

// Create : method
func (uService *UserService) Create(ctx context.Context, in *pbu.User, out *pbu.Response) error {
	log.Println("Creating user: ", in)

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("error hashing password: %v", err))
	}

	in.Password = string(hashedPass)
	if err := uService.repo.Create(in); err != nil {
		return fmt.Errorf(fmt.Sprintf("error creating user: %v", err))
	}

	token, err := uService.tokenService.Encode(in)
	if err != nil {
		return err
	}

	out.User = in
	out.Token = &pbu.Token{Token: token}

	return nil
}

// Get : method
func (uService *UserService) Get(ctx context.Context, in *pbu.User, out *pbu.Response) error {
	user, err := uService.repo.Get(in.Id)
	if err != nil {
		return err
	}
	out.User = user
	return nil
}

// GetAll : method
func (uService *UserService) GetAll(ctx context.Context, in *pbu.Request, out *pbu.Response) error {
	users, err := uService.repo.GetAll()
	if err != nil {
		return err
	}
	out.Users = users
	return nil
}

// Auth : method
func (uService *UserService) Auth(ctx context.Context, in *pbu.User, out *pbu.Token) error {
	log.Println("Logging in with:", in.Email, in.Password)
	user, err := uService.repo.GetByEmail(in.Email)
	log.Println(user, err)
	if err != nil {
		return err
	}

	// hashed password and stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return err
	}

	token, err := uService.tokenService.Encode(user)
	if err != nil {
		return err
	}
	out.Token = token
	return nil
}

// ValidateToken : method
func (uService *UserService) ValidateToken(ctx context.Context, in *pbu.Token, out *pbu.Token) error {
	// Decode token
	claims, err := uService.tokenService.Decode(in.Token)

	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return fmt.Errorf("invalid user")
	}

	out.Valid = true

	return nil
}
