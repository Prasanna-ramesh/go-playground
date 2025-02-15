package user

type CreateUserDto struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (createUserDto CreateUserDto) isAgeValid() bool {
	return createUserDto.Age > 0 || createUserDto.Age < 120
}

type UpdateUserDto struct {
	Name *string `json:"name"`
	Age  *uint8  `json:"age,omitempty"`
}

type ResponseDto struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func toResponseDto(user User) ResponseDto {
	return ResponseDto{Id: user.id, Age: user.age, Name: user.name}
}

type ErrorResponse struct {
	Reason     string `json:"reason"`
	StatusText string `json:"statusText"`
}
