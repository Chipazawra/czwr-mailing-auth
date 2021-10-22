package auth

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Chipazawra/czwr-mailing-auth/pkg/jwtmng"
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
	Users:      map[string]string{"admin": "21232f297a57a5a743894a0e4a801fc3"}, // md5 sum
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

	authorized := g.Group("/auth")
	authorized.Use(a.BasicAuthWrapper())
	authorized.Use(gin.BasicAuth(a.config.Users))
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
func (a *Auth) BasicAuthWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.Request.Header.Get("Authorization")

		abort := func() {
			realm := "Basic realm=" + strconv.Quote("Authorization Required")
			c.Header("WWW-Authenticate", realm)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if h != "" {
			dh, err := decodeBasic(h)
			if err != nil {
				abort()
				return
			}
			c.Request.Header.Set("Authorization", dh)
		}
		c.Next()
	}
}

func decodeBasic(header string) (string, error) {
	dh, err := base64.StdEncoding.DecodeString(
		strings.ReplaceAll(header, "Basic ", ""),
	)
	if err != nil {
		return "", err
	}
	usrpass := strings.SplitN(string(dh), ":", 2)
	if len(usrpass) != 2 {
		return "", fmt.Errorf("invalid string")
	}
	base := usrpass[0] + ":" + fmt.Sprintf("%x", md5.Sum([]byte(usrpass[1])))
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(base)), nil

}
