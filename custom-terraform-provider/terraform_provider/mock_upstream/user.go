package mock_upstream

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	User struct {
		Id   int64  `json:"id,omitempty"`
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}
)

func (mockUpstreamClient *MockUpstreamClient) CreateUser(ctx context.Context, user *User) error {
	createUserPath := "/users"

	response, err := mockUpstreamClient.post(ctx, createUserPath, user)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(user); err != nil {
		return err
	}

	return nil
}

func (mockUpstreamClient *MockUpstreamClient) UpdateUser(ctx context.Context, user *User) error {
	createUserPath := fmt.Sprintf("/users/%d", user.Id)

	response, err := mockUpstreamClient.put(ctx, createUserPath, user)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(user); err != nil {
		return err
	}

	return nil
}

func (mockUpstreamClient *MockUpstreamClient) GetUser(ctx context.Context, user *User) error {
	createUserPath := fmt.Sprintf("/users/%d", user.Id)

	response, err := mockUpstreamClient.get(ctx, createUserPath)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(user); err != nil {
		return err
	}

	return nil
}

func (mockUpstreamClient *MockUpstreamClient) DeleteUser(ctx context.Context, user *User) error {
	createUserPath := fmt.Sprintf("/users/%d", user.Id)

	_, err := mockUpstreamClient.delete(ctx, createUserPath)
	if err != nil {
		return err
	}

	return nil
}
