package gktopic

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"io/ioutil"
)

var (
	HtmlWeb = ""
	HtmlH5  = ""
	regSeo  = regexp.MustCompile(`baiduspider|twitterbot|facebookexternalhit|rogerbot|linkedinbot|embedly|quora link preview|showyoubot|outbrain|pinterest|slackbot|vkShare|W3C_Validator`)
	regH5   = regexp.MustCompile(`(MIDP)|(WAP)|(UP.Browser)|(Smartphone)|(Obigo)|(Mobile)|(AU.Browser)|(wxd.Mms)|(WxdB.Browser)|(CLDC)|(UP.Link)|(KM.Browser)|(UCWEB)|(SEMC\-Browser)|(Mini)|(Symbian)|(Palm)|(Nokia)|(Panasonic)|(MOT\-)|(SonyEricsson)|(NEC\-)|(Alcatel)|(Ericsson)|(BENQ)|(BenQ)|(Amoisonic)|(Amoi\-)|(Capitel)|(PHILIPS)|(SAMSUNG)|(Lenovo)|(Mitsu)|(Motorola)|(SHARP)|(WAPPER)|(LG\-)|(LG/)|(EG900)|(CECT)|(Compal)|(kejian)|(Bird)|(BIRD)|(G900/V1.0)|(Arima)|(CTL)|(TDG)|(Daxian)|(DAXIAN)|(DBTEL)|(Eastcom)|(EASTCOM)|(PANTECH)|(Dopod)|(Haier)|(HAIER)|(KONKA)|(KEJIAN)|(LENOVO)|(Soutec)|(SOUTEC)|(SAGEM)|(SEC\-)|(SED\-)|(EMOL\-)|(INNO55)|(ZTE)|(iPhone)|(Android)|(Windows CE)|(Wget)|(Java)|(curl)|(Opera)/`)
)

func init() {
	w, e := ioutil.ReadFile("web/index.html")
	if e == nil {
		HtmlWeb = string(w)
	}
	w2, e2 := ioutil.ReadFile("m/index.html")
	if e2 == nil {
		HtmlH5 = string(w2)
	}
}
func WebHome(ctx *gin.Context) {
	ua := GetUa(ctx)
	if ua == "web" {
		ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlWeb))
	} else if ua == "h5" {
		ctx.Data(200, "text/html;charset=utf-8", [] byte(HtmlH5))
	} else {
		WebSeoHome(ctx)
	}
}

func GetUa(ctx *gin.Context) string {
	ua := ctx.Request.UserAgent()
	if len(ua) == 0 {
		return "web"
	}
	if regH5.MatchString(ua) {
		return "h5"
	}
	if regSeo.MatchString(ua) {
		return "seo"
	}
	return "web"
}
