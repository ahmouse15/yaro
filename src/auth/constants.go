package auth

import (
	"net/url"
	"strconv"
	"yaro/utils"
)

const clientSecret = "4mRLtYMb6vq9qglomWEaT4ChxsXWcyqbQpuBNfMPOYOiDmYYQmjuaBsF2Zp0RyVeWkfqhE9TuGgAw7te"
const redirectUri = "qrc:///html/login_successful.html"
const pcSign = `eyJhdiI6InYxIiwiYnNuIjoiU3lzdGVtIFNlcmlhbCBOdW1iZXIiLCJnaWQiOjU3MTAsImhzbiI6IjAwMDEiLCJtYWMiOiIkMDEiLCJtaWQiOiIwIiwibXNuIjoiMCIsInN2IjoidjIiLCJ0cyI6IjIwMjQtNy0xMCAxOjIzOjQxOjU5MSJ9.wHTqyvV0amAPG_oaTEZEYSkksxzTU5KLu0dQGCqWgW4`

// var pcSignInput = `{"av":"v1","bsn":"System Serial Number","gid":5710,"hsn":"0001","mac":"$01","mid":"0","msn":"0","sv":"v2","ts":"2024-7-10 1:23:41:591"}`
// var pcSignEncoded = utils.ToBase64Url(pcSignInput)
// var pcSign = pcSignEncoded + "." + utils.GenerateHmac(pcSignSecret, pcSignEncoded)

var Verifier = utils.GenerateCodeVerifier()
var challenge = utils.GenerateCodeChallenge(Verifier)

var nonce = utils.GenerateNonce()

var paramValues = url.Values{"token_format": {"JWS"},
	"code_challenge_method": {"S256"},
	"sbiod_enabled":         {"true"},
	"client_id":             {"JUNO_PC_CLIENT"},
	"code_challenge":        {challenge},
	"pc_sign":               {pcSign},
	"nonce":                 {strconv.Itoa(nonce)},
	"redirect_uri":          {redirectUri},
	"response_type":         {"code id_token"}}

var params = "?" + paramValues.Encode()

var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Origin/10.6.0.00000 EAApp/13.229.0.5742 Chrome/109.0.5414.120 Safari/537.36"
