package managers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/EleMint/CONTd/service/contracts"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// UserManager struct
type UserManager struct {
	sess   *session.Session
	client *cognito.CognitoIdentityProvider
}

// NewUserManager creates a new user manager instance
func NewUserManager() *UserManager {
	log.Printf(
		"Initializing connection to pool: '%s' in region: '%s'\n",
		os.Getenv("POOL_ID"),
		os.Getenv("AWS_REGION"),
	)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigStateFromEnv,
	}))

	client := cognito.New(sess)

	return &UserManager{
		sess:   sess,
		client: client,
	}
}

// GetUserByID gets a user by id
func (manager *UserManager) GetUserByID(c *contracts.GetUserByID) chan *contracts.GetUserResponse {
	rch := make(chan *contracts.GetUserResponse)

	go func() {
		poolID := os.Getenv("POOL_ID")
		limit := int64(1)
		sub := fmt.Sprintf("sub = \"%s\"", c.ID)

		results, err := manager.client.ListUsers(
			&cognito.ListUsersInput{
				UserPoolId: &poolID,
				Limit:      &limit,
				Filter:     &sub,
			},
		)
		if err != nil {
			rch <- &contracts.GetUserResponse{
				Response: contracts.NewResponse(
					http.StatusInternalServerError,
					map[string]string{
						"ListUsers": fmt.Sprintf("error occurred while listing users `%v`", err),
					},
					map[string]string{},
				),
				FoundUser: nil,
			}
			return
		}

		resp := &contracts.GetUserResponse{}
		for _, u := range results.Users {
			subAttr := *u.Username

			if subAttr == c.ID {
				resp.FoundUser = u
				break
			}
		}

		if resp.FoundUser == nil {
			resp.Response = contracts.NewResponse(
				http.StatusNotFound,
				map[string]string{
					"UserNotFound": "user not found",
				},
				map[string]string{},
			)
			rch <- resp
		} else {
			resp.Response = contracts.NewResponse(
				http.StatusOK,
				map[string]string{},
				map[string]string{},
			)
			rch <- resp
		}
	}()

	return rch
}

// GetUserByEmail gets a user by email
func (manager *UserManager) GetUserByEmail(c *contracts.GetUserByEmail) chan *contracts.GetUserResponse {
	rch := make(chan *contracts.GetUserResponse)

	go func() {
		poolID := os.Getenv("POOL_ID")
		limit := int64(1)
		email := fmt.Sprintf("email = \"%s\"", c.Email)

		results, err := manager.client.ListUsers(
			&cognito.ListUsersInput{
				UserPoolId: &poolID,
				Limit:      &limit,
				Filter:     &email,
			},
		)
		if err != nil {
			rch <- &contracts.GetUserResponse{
				Response: contracts.NewResponse(
					http.StatusInternalServerError,
					map[string]string{
						"ListUsers": fmt.Sprintf("error occurred while listing users `%v`", err),
					},
					map[string]string{},
				),
				FoundUser: nil,
			}
			return
		}

		resp := &contracts.GetUserResponse{}
		for _, u := range results.Users {
			emailAttr := ""
			for _, attr := range u.Attributes {
				if *attr.Name == "email" {
					emailAttr = *attr.Value
					break
				}
			}
			log.Println(emailAttr)

			if emailAttr == c.Email {
				resp.FoundUser = u
				break
			}
		}

		log.Println(&resp.FoundUser)

		if resp.FoundUser == nil {
			resp.Response = contracts.NewResponse(
				http.StatusNotFound,
				map[string]string{
					"UserNotFound": "user not found",
				},
				map[string]string{},
			)
			rch <- resp
		} else {
			resp.Response = contracts.NewResponse(
				http.StatusOK,
				map[string]string{},
				map[string]string{},
			)
			rch <- resp
		}
	}()

	return rch
}

// GetUserByDisplayName gets a user by display name
func (manager *UserManager) GetUserByDisplayName(c *contracts.GetUserByDisplayName) chan *contracts.GetUserResponse {
	rch := make(chan *contracts.GetUserResponse)

	go func() {
		poolID := os.Getenv("POOL_ID")
		limit := int64(1)
		dname := fmt.Sprintf("preferred_username = \"%s\"", c.Name)

		results, err := manager.client.ListUsers(
			&cognito.ListUsersInput{
				UserPoolId: &poolID,
				Limit:      &limit,
				Filter:     &dname,
			},
		)
		if err != nil {
			rch <- &contracts.GetUserResponse{
				Response: contracts.NewResponse(
					http.StatusInternalServerError,
					map[string]string{
						"ListUsers": fmt.Sprintf("error occurred while listing users `%v`", err),
					},
					map[string]string{},
				),
				FoundUser: nil,
			}
			return
		}

		resp := &contracts.GetUserResponse{}
		for _, u := range results.Users {
			dnameAttr := ""
			for _, attr := range u.Attributes {
				if *attr.Name == "preferred_username" {
					dnameAttr = *attr.Value
					break
				}
			}

			if dnameAttr == c.Name {
				resp.FoundUser = u
				break
			}
		}

		if resp.FoundUser == nil {
			resp.Response = contracts.NewResponse(
				http.StatusNotFound,
				map[string]string{
					"UserNotFound": "user not found",
				},
				map[string]string{},
			)
			rch <- resp
		} else {
			resp.Response = contracts.NewResponse(
				http.StatusOK,
				map[string]string{},
				map[string]string{},
			)
			rch <- resp
		}
	}()

	return rch
}

// GetUsersByStartDisplayName gets a user by a starting string of a user's display name
func (manager *UserManager) GetUsersByStartDisplayName(c *contracts.GetUsersByStartDisplayName) chan *contracts.GetUsersResponse {
	rch := make(chan *contracts.GetUsersResponse)

	go func() {
		poolID := os.Getenv("POOL_ID")
		substr := fmt.Sprintf("preferred_username ^= \"%s\"", c.Substring)
		limit := int64(c.Limit)

		results, err := manager.client.ListUsers(
			&cognito.ListUsersInput{
				UserPoolId: &poolID,
				Limit:      &limit,
				Filter:     &substr,
			},
		)
		if err != nil {
			rch <- &contracts.GetUsersResponse{
				Response: contracts.NewResponse(
					http.StatusInternalServerError,
					map[string]string{
						"ListUsers": fmt.Sprintf("error occurred while listing users `%v`", err),
					},
					map[string]string{},
				),
				FoundUsers: nil,
			}
			return
		}

		rch <- &contracts.GetUsersResponse{
			Response: contracts.NewResponse(
				http.StatusNotFound,
				map[string]string{
					"UserNotFound": "user not found",
				},
				map[string]string{},
			),
			FoundUsers: results.Users,
		}
	}()

	return rch
}

// GetUsersByStartEmail gets a user by a starting string of a user's email
func (manager *UserManager) GetUsersByStartEmail(c *contracts.GetUsersByStartEmail) chan *contracts.GetUsersResponse {
	rch := make(chan *contracts.GetUsersResponse)

	go func() {
		poolID := os.Getenv("POOL_ID")
		substr := fmt.Sprintf("email ^= \"%s\"", c.Substring)
		limit := c.Limit

		results, err := manager.client.ListUsers(
			&cognito.ListUsersInput{
				UserPoolId: &poolID,
				Limit:      &limit,
				Filter:     &substr,
			},
		)
		if err != nil {
			rch <- &contracts.GetUsersResponse{
				Response: contracts.NewResponse(
					http.StatusInternalServerError,
					map[string]string{
						"ListUsers": fmt.Sprintf("error occurred while listing users `%v`", err),
					},
					map[string]string{},
				),
				FoundUsers: nil,
			}
			return
		}

		rch <- &contracts.GetUsersResponse{
			Response: contracts.NewResponse(
				http.StatusNotFound,
				map[string]string{},
				map[string]string{},
			),
			FoundUsers: results.Users,
		}
	}()

	return rch
}

// CreateUser creates a new user by contract
func (manager *UserManager) CreateUser(c *contracts.CreateUser) chan *contracts.CreateUserResponse {
	rch := make(chan *contracts.CreateUserResponse)

	go func() {
		clientID := os.Getenv("CLIENT_ID")
		prefUname := c.Email[0:strings.LastIndex(c.Email, "@")]

		newUser := &cognito.SignUpInput{
			Username: aws.String(c.Email),
			Password: aws.String(c.Password),
			ClientId: aws.String(clientID),
			UserAttributes: []*cognito.AttributeType{
				{
					Name:  aws.String("email"),
					Value: aws.String(c.Email),
				},
				{
					Name:  aws.String("preferred_username"),
					Value: aws.String(prefUname),
				},
			},
		}

		createdUser, err := manager.client.SignUp(newUser)
		if err != nil {
			status := http.StatusInternalServerError
			_, isConflictErr := err.(*cognito.UsernameExistsException)
			if isConflictErr {
				status = http.StatusConflict
			}
			rch <- &contracts.CreateUserResponse{
				Response: contracts.NewResponse(
					status,
					map[string]string{
						"CreateUser": fmt.Sprintf("error occurred while creating user `%v`", err),
					},
					map[string]string{},
				),
				ID: "",
			}
		} else {
			rch <- &contracts.CreateUserResponse{
				Response: contracts.NewResponse(
					http.StatusOK,
					map[string]string{},
					map[string]string{},
				),
				ID: *createdUser.UserSub,
			}
		}
	}()

	return rch
}

// UpdateUserDisplayName updates a user's display name
func (manager *UserManager) UpdateUserDisplayName(c *contracts.UpdateUserDisplayName) chan *contracts.UpdateUserDisplayNameResponse {
	rch := make(chan *contracts.UpdateUserDisplayNameResponse)

	go func() {
		poolID := os.Getenv("POOL_ID")
		email := manager.getEmailFromToken(c.Token)

		update := &cognito.AdminUpdateUserAttributesInput{
			Username:   aws.String(email),
			UserPoolId: aws.String(poolID),
			UserAttributes: []*cognito.AttributeType{
				{
					Name:  aws.String("preferred_username"),
					Value: aws.String(c.Name),
				},
			},
		}
		_, err := manager.client.AdminUpdateUserAttributes(update)
		if err != nil {
			rch <- &contracts.UpdateUserDisplayNameResponse{
				Response: contracts.NewResponse(
					http.StatusInternalServerError,
					map[string]string{
						"UpdateUserAttributes": fmt.Sprintf("error occurred while updating user attributes `%v`", err),
					},
					map[string]string{},
				),
			}
		} else {
			rch <- &contracts.UpdateUserDisplayNameResponse{
				Response: contracts.NewResponse(
					http.StatusOK,
					map[string]string{},
					map[string]string{},
				),
			}
		}
	}()

	return rch
}

// generateTempPassword generates a temp password ex: '3i[g0|)z'
func (manager *UserManager) generateTempPassword() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	all := digits + specials + upper + lower

	length := 8
	buf := make([]byte, length)

	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	buf[2] = upper[rand.Intn(len(upper))]
	buf[3] = lower[rand.Intn(len(lower))]
	for i := 4; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}

	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}

func (manager *UserManager) getEmailFromToken(token string) string {
	return "test@test.com"
}
