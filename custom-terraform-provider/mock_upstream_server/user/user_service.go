package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

func sendJsonResponse(writer http.ResponseWriter, code int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)

	if err := json.NewEncoder(writer).Encode(data); err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
	}
}

func getUserHandler(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseUint(chi.URLParam(request, "userId"), 10, 0)
	if err != nil {
		sendJsonResponse(
			writer,
			http.StatusBadRequest,
			ErrorResponse{
				Reason:     "invalid user id",
				StatusText: http.StatusText(http.StatusBadRequest),
			},
		)
		return
	}

	user := findUserById(uint(userId))
	if user == nil {
		sendJsonResponse(writer, http.StatusNotFound,
			ErrorResponse{
				Reason:     fmt.Sprintf("invalid user id %d", userId),
				StatusText: http.StatusText(http.StatusNotFound),
			},
		)
		return
	}

	sendJsonResponse(writer, http.StatusOK, toResponseDto(*user))
}

func getUsersHandler(writer http.ResponseWriter, request *http.Request) {
	users := getAllUsers()

	usersResponse := make([]ResponseDto, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, toResponseDto(user))
	}

	sendJsonResponse(writer, http.StatusOK, usersResponse)
}

func createUserHandler(writer http.ResponseWriter, request *http.Request) {
	var user CreateUserDto

	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		sendJsonResponse(
			writer,
			http.StatusBadRequest,
			ErrorResponse{
				Reason:     "invalid body",
				StatusText: http.StatusText(http.StatusBadRequest),
			},
		)
		return
	}

	for _, existingUser := range getAllUsers() {
		if existingUser.name == user.Name && existingUser.age == user.Age {
			sendJsonResponse(
				writer,
				http.StatusConflict,
				ErrorResponse{
					Reason:     "user already exists",
					StatusText: http.StatusText(http.StatusConflict),
				},
			)
			return
		}
	}

	newUser := toResponseDto(save(User{id: uint(time.Now().UnixMicro()), name: user.Name, age: user.Age}))

	sendJsonResponse(writer, http.StatusCreated, newUser)
}

func modifyUserHandler(writer http.ResponseWriter, request *http.Request) {
	var updateUserDetails UpdateUserDto

	defer request.Body.Close()
	if err := json.NewDecoder(request.Body).Decode(&updateUserDetails); err != nil {
		sendJsonResponse(
			writer,
			http.StatusBadRequest,
			ErrorResponse{
				Reason:     "invalid body",
				StatusText: http.StatusText(http.StatusBadRequest),
			},
		)
		return
	}

	userId, err := strconv.ParseUint(chi.URLParam(request, "userId"), 10, 0)
	if err != nil {
		sendJsonResponse(
			writer,
			http.StatusBadRequest,
			ErrorResponse{
				Reason:     "invalid user id",
				StatusText: http.StatusText(http.StatusBadRequest),
			},
		)
		return
	}

	user := findUserById(uint(userId))
	if user == nil {
		sendJsonResponse(writer, http.StatusNotFound,
			ErrorResponse{
				Reason:     fmt.Sprintf("invalid user id %d", userId),
				StatusText: http.StatusText(http.StatusNotFound),
			},
		)
		return
	}

	if updateUserDetails.Age != nil {
		user.age = *updateUserDetails.Age
	}

	if updateUserDetails.Name != nil {
		user.name = *updateUserDetails.Name
	}

	updatedUser := updateUser(*user)
	if updatedUser == nil {
		sendJsonResponse(writer, http.StatusInternalServerError,
			ErrorResponse{
				Reason:     fmt.Sprintf("Update failed"),
				StatusText: http.StatusText(http.StatusInternalServerError),
			},
		)
		return
	}

	sendJsonResponse(writer, http.StatusOK, toResponseDto(*updatedUser))
}

func deleteUserHandler(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.ParseUint(chi.URLParam(request, "userId"), 10, 0)
	if err != nil {
		sendJsonResponse(
			writer,
			http.StatusBadRequest,
			ErrorResponse{
				Reason:     "invalid user id type",
				StatusText: http.StatusText(http.StatusBadRequest),
			},
		)
		return
	}

	err = deleteUserById(uint(userId))
	if err != nil {
		sendJsonResponse(
			writer,
			http.StatusNotFound,
			ErrorResponse{
				Reason:     fmt.Sprintf("invalid user id %d", userId),
				StatusText: http.StatusText(http.StatusBadRequest),
			},
		)
		return
	}

	sendJsonResponse(
		writer,
		http.StatusOK,
		struct{}{},
	)
}
