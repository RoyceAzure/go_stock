package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto      *paseto.V2
	symmerickey []byte
}

func NewPasetoMaker(symmerickey string) (Maker, error) {
	if len(symmerickey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid ket size : must be exactly %d charcters", chacha20poly1305.KeySize)
	}
	return &PasetoMaker{paseto.NewV2(), []byte(symmerickey)}, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	//相比jwt之所以只有這行是因為你不需要決定加密演算法
	//固定使用chacha演算法
	return maker.paseto.Encrypt(maker.symmerickey, payload, nil)
}

func (maker *PasetoMaker) VertifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	//之所以這麼簡單  是因為套件只會回傳ErrInvalidToken相關範圍得錯誤
	//你自己的paload valid要自己呼叫
	//也因為如此  不像jwt驗證是通通包再一起，你必須拆解jwt回傳的錯誤訊息
	//目前看到的InvalidToken錯誤包括key  資料不匹配
	err := maker.paseto.Decrypt(token, maker.symmerickey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
