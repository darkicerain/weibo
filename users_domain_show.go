// https://open.weibo.com/wiki/2/users/domain_show
// 请求参数
// access_token	true	string	采用OAuth授权方式为必填参数，OAuth授权后获得。
// domain	true	string	需要查询的个性化域名。
//
// 注意事项
// 接口升级后，对未授权本应用的uid，将无法获取其个人简介、认证原因、粉丝数、关注数、微博数及最近一条微博内容。
//
// 返回字段说明
//   id	int64	用户UID
//   idstr	string	字符串型的用户UID
//   screen_name	string	用户昵称
//   name	string	友好显示名称
//   province	int	用户所在省级ID
//   city	int	用户所在城市ID
//   location	string	用户所在地
//   description	string	用户个人描述
//   url	string	用户博客地址
//   profile_image_url	string	用户头像地址（中图），50×50像素
//   profile_url	string	用户的微博统一URL地址
//   domain	string	用户的个性化域名
//   weihao	string	用户的微号
//   gender	string	性别，m：男、f：女、n：未知
//   followers_count	int	粉丝数
//   friends_count	int	关注数
//   statuses_count	int	微博数
//   favourites_count	int	收藏数
//   created_at	string	用户创建（注册）时间
//   following	boolean	暂未支持
//   allow_all_act_msg	boolean	是否允许所有人给我发私信，true：是，false：否
//   geo_enabled	boolean	是否允许标识用户的地理位置，true：是，false：否
//   verified	boolean	是否是微博认证用户，即加V用户，true：是，false：否
//   verified_type	int	暂未支持
//   remark	string	用户备注信息，只有在查询用户关系时才返回此字段
//   status	object	用户的最近一条微博信息字段 详细
//   allow_all_comment	boolean	是否允许所有人对我的微博进行评论，true：是，false：否
//   avatar_large	string	用户头像地址（大图），180×180像素
//   avatar_hd	string	用户头像地址（高清），高清头像原图
//   verified_reason	string	认证原因
//   follow_me	boolean	该用户是否关注当前登录用户，true：是，false：否
//   online_status	int	用户的在线状态，0：不在线、1：在线
//   bi_followers_count	int	用户的互粉数
//   lang	string	用户当前的语言版本，zh-cn：简体中文，zh-tw：繁体中文，en：英语

package weibo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// UsersDomainShow 通过个性化域名获取用户资料以及用户最新的一条微博
func (w *Weibo) UsersDomainShow(token, domain string) (*UsersDomainShowResp, error) {
	apiURL := "https://api.weibo.com/2/users/domain_show.json"
	data := url.Values{
		"access_token": {token},
		"domain":       {domain},
	}
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "weibo UsersDomainShow NewRequest error")
	}
	// 这个接口不能设置这个Header，会报错
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = data.Encode()
	resp, err := w.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "weibo UsersDomainShow Do error")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "weibo UsersDomainShow ReadAll error")
	}
	r := &UsersDomainShowResp{}
	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.Wrap(err, "weibo UsersDomainShow Unmarshal error:"+string(body))
	}
	if r.Error != "" && r.ErrorCode != 0 {
		return nil, errors.New("weibo UsersDomainShow resp error:" + r.Error)
	}
	return r, nil
}
