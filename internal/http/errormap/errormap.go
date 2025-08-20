package errormap


import (
    "net/http"
    "github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/pkg/apperrors"
)

func HttpStatusFrom(code apperrors.Code) int {
    switch code {
    case apperrors.CodeNotFound:
        return http.StatusNotFound
    case apperrors.CodeEmailExists:
        return http.StatusConflict
    case apperrors.CodeInvalidArgument:
        return http.StatusBadRequest
    default:
        return http.StatusInternalServerError
    }
}
