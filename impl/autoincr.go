package impl

import (
	"errors"
	"regexp"
	"shorturl/util"
)
const (
	shortToLong string = "shorttolong"
	longToShort string = "longtoshort"
)
type URLConventor struct {
	Map *util.RedisMap
	Count *util.IntBase62
}
func NewURLConventor(redisMap *util.RedisMap) *URLConventor {
	return &URLConventor{
		Map:redisMap,
		Count:util.NewIntBase62("0"),
	}
}
func (this *URLConventor) Generator(url string)(string,error) {
	this.Count.Acc()
	shortUrl := this.Count.String()
	err := this.Map.Set(longToShort,url,shortUrl)
	if err != nil {
		return "",err
	}else {
		err = this.Map.Set(shortToLong,shortUrl,url)
		if err != nil {
			return "",err
		}else {
			return shortUrl,nil
		}
	}
}
func (this *URLConventor) Revert(url string)(string,error) {
	isBase62Num,err:= regexp.MatchString(`\w{1,6}`,url)
	if !isBase62Num {
		return "",err
	}
	now := util.NewIntBase62(url)
	if now.GreaterThan(this.Count) {
		return "",errors.New("not exists url")
	}
	longURL,err := this.Map.Get(shortToLong,url)
	if err != nil {
		return "",err
	} else {
		return longURL,nil
	}
}
