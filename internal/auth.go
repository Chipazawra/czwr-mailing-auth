package auth

import (
	"net/http"
	"time"

	"github.com/Chipazawra/czwrMailing-auth/pkg/jwtmng"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	config *Config
}

type Config struct {
	Users      map[string]string `yaml:"users"`
	JwtTTL     int               `yaml:"jwtttl"`
	RefreshTTL int               `yaml:"refreshttl"`
}

var DefaultConfig = &Config{
	Users:      map[string]string{"admin": "admin"},
	JwtTTL:     60,
	RefreshTTL: 180,
}

func New(config *Config) *Auth {

	if config == nil {
		config = DefaultConfig
	}

	return &Auth{config: config}
}

func (a *Auth) Register(g *gin.Engine) *gin.RouterGroup {

	authorized := g.Group("/auth", gin.BasicAuth(a.config.Users))
	authorized.GET("/login", a.loginHandler)
	authorized.GET("/logout", a.logoutHandler)

	return authorized

}

// login godoc
// @Summary login in service
// @Tags auth
// @Description get auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /login [get]
func (a *Auth) loginHandler(c *gin.Context) {

	token, _ := jwtmng.NewJWT("usr", time.Duration(a.config.JwtTTL))
	refresh, _ := jwtmng.NewRefreshToken()

	c.SetCookie("access", token, a.config.JwtTTL, "/", "localhost", false, true)
	c.SetCookie("refresh", refresh, a.config.RefreshTTL, "/", "localhost", false, true)

	rurl, rdct := c.GetQuery("redirect_uri")

	if rdct {
		c.Redirect(http.StatusPermanentRedirect, rurl)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "login.",
		})
	}
}

// logout godoc
// @Summary logout from service
// @Tags auth
// @Description clear auth data
// @Accept  json
// @Produce  json
// @Success 200
// @Router /logout [get]
func (a *Auth) logoutHandler(c *gin.Context) {

	c.SetCookie("access", "", -1, "/", "localhost", false, true)
	c.SetCookie("refresh", "", -1, "/", "localhost", false, true)

	rurl, rdct := c.GetQuery("redirect_uri")

	if rdct {
		c.Redirect(http.StatusPermanentRedirect, rurl)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "logout.",
		})
	}
}
