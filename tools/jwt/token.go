package jwt

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/9299381/wego/configs"
	"github.com/9299381/wego/tools/errors"
	"strconv"
	"strings"
	"time"
)

type Token struct {
	Claims *Claims
	config *configs.TokenConfig
}

func NewToken() *Token {
	token:= &Token{
		Claims:&Claims{},
	}
	token.config = (&configs.TokenConfig{}).Load().(*configs.TokenConfig)
	return token
}

func (it *Token) SetId(id string) *Token {
	it.Claims.Id = id
	return it
}

func (it *Token) SetName(name string) *Token {
	it.Claims.Name =name
	return it
}

func (it *Token) SetRole(role string) *Token {
	it.Claims.Role = role
	return it
}

func (it *Token) GetToken() string {
	it.Claims.Iat = time.Now().Unix()
	it.Claims.Exp = it.getExpTime()
	jsonClaim, err := json.Marshal(it.Claims)
	if err != nil{
		panic(err)
	}
	payload := base64.StdEncoding.EncodeToString(jsonClaim)
	sign := it.getSign(it.Claims)
	ret := string(payload) + "." + sign
	return ret
}


func (it *Token) VerifyToken(sign string) (*Claims,error) {
	m:=strings.Split(sign,".")
	if len(m) <1 {
		return nil , errors.New("6100","格式错误,请重新登陆")
	}
	jsonClaim, decodeErr := base64.StdEncoding.DecodeString(m[0])
	if decodeErr !=nil{
		return nil,decodeErr
	}
	claims := &Claims{}
	jsonErr := json.Unmarshal(jsonClaim, claims)
	if jsonErr!=nil{
		return nil,jsonErr
	}

	if claims.Exp < time.Now().Unix() {
		return nil,errors.New("6200","登陆过期,请重新登陆")
	}

	if m[1] != it.getSign(claims) {
		return nil,errors.New("6300","签名错误,请重新登陆")
	}
	return claims,nil
}
func (it *Token) getExpTime() int64 {
	period := it.config.Exp
	return it.Claims.Iat + period
}
func (it *Token)getSign(claims *Claims) string {
	key := it.config.Key
	keyPlain := claims.Id + strconv.Itoa(int(claims.Iat)) + key
	h:=md5.New()
	h.Write([]byte(keyPlain))
	return hex.EncodeToString(h.Sum(nil))

}