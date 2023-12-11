package selector

import "github.com/gin-gonic/gin"

func SelectLanguageFormat(ctx *gin.Context, nameHtml string) string {
	if value, _ := ctx.Get("lang"); value == "ru" {
		return "ru-" + nameHtml
	}

	return "en-" + nameHtml
}
