package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ndenisj/go_mem/account/model/apperrors"
)

type tokenReq struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	// bind data to req of type tokenReq
	var req tokenReq

	if ok := bindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	// verify refresh jwt
	refreshToken, err := h.TokenService.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// get up to date user
	u, err := h.UserService.Get(ctx, refreshToken.UID)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create fresh pair of tokens
	tokens, err := h.TokenService.NewPairFromUser(ctx, u, refreshToken.ID.String())
	if err != nil {
		log.Printf("Failed to create tokens for user: %+v. Error: %v\n", u, err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}
