package rest

import (
	"MuchUp/backend/internal/domain/entity"
	"MuchUp/backend/internal/domain/usecase"
	"MuchUp/backend/pkg/auth"
	"MuchUp/backend/pkg/logger"
	"MuchUp/backend/pkg/middleware"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)
type Handler struct {
	userUsecase    usecase.UserUsecase
	messageUsecase usecase.MessageUsecase
	logger         logger.Logger
}
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}
type CreateUserRequest struct {
	Name               string                 `json:"name" validate:"required,min=2,max=50"`
	Email              string                 `json:"email" validate:"required,email"`
	PasswordHash       string                 `json:"password" validate:"required,min=8"`
	UsagePurpose       string                 `json:"usage_purpose"`
	PersonalityProfile map[string]interface{} `json:"personality_profile"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type CreateMessageRequest struct {
	Content string `json:"content" validate:"required,min=1,max=1000"`
	GroupID string `json:"group_id" validate:"required"`
}
type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" validate:"omitempty,min=2,max=50"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
}
func NewHandler(
	userUsecase usecase.UserUsecase,
	messageUsecase usecase.MessageUsecase,
	logger logger.Logger,
) *Handler {
	return &Handler{
		userUsecase:    userUsecase,
		messageUsecase: messageUsecase,
		logger:         logger,
	}
}

func (h *Handler) SetupRoutes(validator auth.TokenValidator) *mux.Router{
	r := mux.NewRouter()

    api := r.PathPrefix("/api/v1").Subrouter()

    
    api.HandleFunc("/auth/login", h.Login).Methods("POST")
    api.HandleFunc("/users", h.CreateUser).Methods("POST") 
    api.HandleFunc("/health", h.HealthCheck).Methods("GET") 

    api.Handle("/users/{id}", middleware.JWTMiddleware(http.HandlerFunc(h.GetUser), validator)).Methods("GET")
    api.Handle("/users/{id}", middleware.JWTMiddleware(http.HandlerFunc(h.UpdateUser), validator)).Methods("PUT")
    api.Handle("/users/{id}", middleware.JWTMiddleware(http.HandlerFunc(h.DeleteUser), validator)).Methods("DELETE")
    api.Handle("/users", middleware.JWTMiddleware(http.HandlerFunc(h.GetUsers), validator)).Methods("GET")

    api.Handle("/auth/logout", middleware.JWTMiddleware(http.HandlerFunc(h.Logout), validator)).Methods("POST")
    api.Handle("/messages", middleware.JWTMiddleware(http.HandlerFunc(h.CreateMessage), validator)).Methods("POST")
    api.Handle("/messages/{id}", middleware.JWTMiddleware(http.HandlerFunc(h.GetMessage), validator)).Methods("GET")
    api.Handle("/messages/{id}", middleware.JWTMiddleware(http.HandlerFunc(h.UpdateMessage), validator)).Methods("PUT")
    api.Handle("/messages/{id}", middleware.JWTMiddleware(http.HandlerFunc(h.DeleteMessage), validator)).Methods("DELETE")

    
    api.Handle("/groups/{group_id}/messages", middleware.JWTMiddleware(http.HandlerFunc(h.GetMessagesByGroup), validator)).Methods("GET")
    api.Handle("/groups/{group_id}/users", middleware.JWTMiddleware(http.HandlerFunc(h.GetGroupUsers), validator)).Methods("GET")
    api.Handle("/groups/{group_id}/join", middleware.JWTMiddleware(http.HandlerFunc(h.JoinGroup), validator)).Methods("POST")
    api.Handle("/groups/{group_id}/leave", middleware.JWTMiddleware(http.HandlerFunc(h.LeaveGroup), validator)).Methods("POST")

    r.HandleFunc("/health", h.HealthCheck).Methods("GET")

    return r
}

// @Summary Creating User By entity
// @description Creating User by request
// @Tags User
// @Accept json
// @Produce json 
// @Success 200 {object} entity.User
// @Failure 500 {object} string
// @Router /users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	user := &entity.User{
		NickName:           req.Name,
		Email:              &req.Email,
		PasswordHash:       req.PasswordHash,
		UsagePurpose:       req.UsagePurpose,
		PersonalityProfile: req.PersonalityProfile,
	}
	createdUser, err := h.userUsecase.CreateUser(user)
	if err != nil {
		h.logger.Error("Failed to create user", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to create user")
		return
	}
	createdUser.PasswordHash = ""
	h.sendSuccessResponse(w, http.StatusCreated, createdUser, "User created successfully")
}

// @Summary UserIDからユーザーを取得
// @Tags User
// @Accept json 
// @Produce json
// @Param id path string true "User ID"
// @Param       Authorization header string  true  "認証トークン (Bearer)"
// @Success 200 {object} entity.User
// @Failure 404 {object}  string
// @Failure 401 {object} string "authorization error" 
// @Router /users/{id} [get]
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		h.logger.Error("Failed to get user", err)
		h.sendErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	user.PasswordHash = ""
	h.sendSuccessResponse(w, http.StatusOK, user, "")
}

// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param       Authorization header string  true  "認証トークン (Bearer)"
// @Success 200 {object} entity.User
// @Failure 404 {object}  string
// @Failure 401 {object} string "authorization error" 
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	user, err := h.userUsecase.GetUserByID(userID)
	if err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	if req.Name != "" {
		user.NickName = req.Name
	}
	if req.Email != "" {
		user.Email = &req.Email
	}
	updatedUser, err := h.userUsecase.UpdateUser(user)
	if err != nil {
		h.logger.Error("Failed to update user", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update user")
		return
	}
	updatedUser.PasswordHash = ""
	h.sendSuccessResponse(w, http.StatusOK, updatedUser, "User updated successfully")
}

// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param       Authorization header string  true  "認証トークン (Bearer)"
// @Success 200 {object} entity.User
// @Failure 404 {object}  string
// @Failure 401 {object} string "authorization error" 
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	err := h.userUsecase.DeleteUser(userID)
	if err != nil {
		h.logger.Error("Failed to delete user", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	h.sendSuccessResponse(w, http.StatusOK, nil, "User deleted successfully")
}
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit := 10
	offset := 0
	if l := query.Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	if o := query.Get("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}
	users, err := h.userUsecase.GetUsers(limit, offset)
	if err != nil {
		h.logger.Error("Failed to get users", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get users")
		return
	}
	for i := range users {
		users[i].PasswordHash = ""
	}
	h.sendSuccessResponse(w, http.StatusOK, users, "")
}

// @Summary emailとpasswordでログイン
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} string
// @Failure 401 {object} string 
// @Roter /auth/login" [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	token, err := h.userUsecase.Login(req.Email, req.Password)
	if err != nil {
		h.logger.Error("Login failed", err)
		h.sendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}
	response := map[string]interface{}{
		"token":      token,
		"expires_at": time.Now().Add(24 * time.Hour).Unix(),
	}
	h.sendSuccessResponse(w, http.StatusOK, response, "Login successful")
}
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	h.sendSuccessResponse(w, http.StatusOK, nil, "Logout successful")
}

// @Tags Message
// @Accept json 
// @Produce json
// @Param   Authorization header string  true  "認証トークン (Bearer)"
// @Success 200 {object} entity.Message
// @Failure 404 {object}  string
// @Failure 401 {object} string "authorization error" 
// @Router /message [post]
func (h *Handler) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var req CreateMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	senderID := "some-user-id"
	message, err := entity.NewMessage(senderID, req.GroupID, req.Content)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	createdMessage, err := h.messageUsecase.CreateMessage(message)
	if err != nil {
		h.logger.Error("Failed to create message", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to create message")
		return
	}
	h.sendSuccessResponse(w, http.StatusCreated, createdMessage, "Message created successfully")
}
func (h *Handler) GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID := vars["id"]
	message, err := h.messageUsecase.GetMessageByID(messageID)
	if err != nil {
		h.logger.Error("Failed to get message", err)
		h.sendErrorResponse(w, http.StatusNotFound, "Message not found")
		return
	}
	h.sendSuccessResponse(w, http.StatusOK, message, "")
}
func (h *Handler) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID := vars["id"]
	var req CreateMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	message, err := h.messageUsecase.GetMessageByID(messageID)
	if err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "Message not found")
		return
	}
	if req.Content != "" {
		message.Text = &req.Content
	}
	updatedMessage, err := h.messageUsecase.UpdateMessage(message)
	if err != nil {
		h.logger.Error("Failed to update message", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update message")
		return
	}
	h.sendSuccessResponse(w, http.StatusOK, updatedMessage, "Message updated successfully")
}
func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID := vars["id"]
	err := h.messageUsecase.DeleteMessage(messageID)
	if err != nil {
		h.logger.Error("Failed to delete message", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete message")
		return
	}
	h.sendSuccessResponse(w, http.StatusOK, nil, "Message deleted successfully")
}
func (h *Handler) GetMessagesByGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["group_id"]
	query := r.URL.Query()
	limit := 50
	offset := 0
	if l := query.Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}
	if o := query.Get("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}
	messages, err := h.messageUsecase.GetMessagesByGroup(groupID, limit, offset)
	if err != nil {
		h.logger.Error("Failed to get messages", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get messages")
		return
	}
	h.sendSuccessResponse(w, http.StatusOK, messages, "")
}
func (h *Handler) GetGroupUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["group_id"]
	users, err := h.userUsecase.GetUsersByGroup(groupID)
	if err != nil {
		h.logger.Error("Failed to get group users", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get group users")
		return
	}
	for i := range users {
		users[i].PasswordHash = ""
	}
	h.sendSuccessResponse(w, http.StatusOK, users, "")
}
func (h *Handler) JoinGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["group_id"]
	userID := r.Header.Get("X-User-ID")
	err := h.userUsecase.JoinGroup(userID, groupID)
	if err != nil {
		h.logger.Error("Failed to join group", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to join group")
		return
	}
	h.sendSuccessResponse(w, http.StatusOK, nil, "Successfully joined group")
}
func (h *Handler) LeaveGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID := vars["group_id"]
	userID := r.Header.Get("X-User-ID")
	err := h.userUsecase.LeaveGroup(userID, groupID)
	if err != nil {
		h.logger.Error("Failed to leave group", err)
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to leave group")
		return
	}
	h.sendSuccessResponse(w, http.StatusOK, nil, "Successfully left group")
}
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.sendSuccessResponse(w, http.StatusOK, map[string]string{
		"status": "healthy",
		"time":   time.Now().UTC().Format(time.RFC3339),
	}, "Service is healthy")
}
func (h *Handler) sendSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		Success: true,
		Data:    data,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}
func (h *Handler) sendErrorResponse(w http.ResponseWriter, statusCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := Response{
		Success: false,
		Error:   errorMsg,
	}
	json.NewEncoder(w).Encode(response)
}
