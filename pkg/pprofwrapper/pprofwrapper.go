package pprofwrapper

import (
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

type PprofWrapper struct {
	PprofEnabled bool
}

func New() *PprofWrapper {
	return &PprofWrapper{PprofEnabled: false}
}

func (pw *PprofWrapper) Register(g *gin.RouterGroup) {
	g.POST("/pprof_enable", pw.enableHandler)
	g.POST("/pprof_disable", pw.disableHandler)
	gr := g.Group("/debug")
	gr.Use(pw.Switch())
	pprof.RouteRegister(gr, "/pprof")
}

// pprofwrapper godoc
// @Summary enable pprof API
// @Tags pprofwrapper
// @Description enable pprof API on service
// @Accept  json
// @Produce  json
// @Success 200
// @Router /debug/pprof_enable [post]
func (pw *PprofWrapper) enableHandler(c *gin.Context) {
	pw.PprofEnabled = true
	c.JSON(
		http.StatusOK,
		gin.H{"status": "pprof enabled."},
	)
}

// pprofwrapper godoc
// @Summary disable pprof API
// @Tags pprofwrapper
// @Description disable pprof API on service
// @Accept  json
// @Produce  json
// @Success 200
// @Router /debug/pprof_disable [post]
func (pw *PprofWrapper) disableHandler(c *gin.Context) {
	pw.PprofEnabled = false
	c.JSON(
		http.StatusOK,
		gin.H{"status": "pprof disabled."},
	)
}

func (pw *PprofWrapper) Switch() gin.HandlerFunc {
	return func(c *gin.Context) {
		if pw.PprofEnabled {
			c.Next()
		} else {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{"status": "pprof disabled."},
			)
		}

	}
}
