namespace java com.uber.cadence.user

struct User {
    10: optional string username
    20: optional bool admin
    30: optional list<string> groups
}

struct AddUserRequest {
    10: optional User user
    20: optional string password
}

struct AddUserResponse {
      10: optional  User user
}

struct DeleteUserRequest {
      10: optional string username
}

struct GetUserRequest {
    10: optional string username
}

struct GetUserResponse {
      10: optional User user
}

struct ListUsersResponse {
    10: optional list<User> users
}

struct UpdateUserRequest {
 10: optional string username
 20: optional User user
}

struct UpdateUserResponse {
  10: optional User user
}

struct ChangeUserPasswordRequest {
  10: optional string username
  20: optional string password
}
