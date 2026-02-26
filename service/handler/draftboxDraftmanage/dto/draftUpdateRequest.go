package dto

import "encoding/json"

type ArticleType string

const (
	ArticleTypeNews    ArticleType = "news"    //图文消息
	ArticleTypeNewsPic ArticleType = "newspic" //图片消息
)

type ImageItem struct {
	ImageMediaId string `json:"image_media_id" comment:"图片消息里的图片素材id（必须是永久MediaID）"`
}
type ImageInfo struct {
	ImageList []ImageItem `json:"image_list" comment:"图片列表"`
}

type CoverItem struct {
	Radio string `json:"radio" comment:"裁剪比例，支持 “1_1”，“16_9”,“2.35_1”"`
	X1    string `json:"x1" comment:"以图片左上角（0,0），右下角（1,1）建立平面坐标系，经过裁剪后的图片，其左上角所在的坐标填入x1，y1参数"`
	Y1    string `json:"y1" comment:"以图片左上角（0,0），右下角（1,1）建立平面坐标系，经过裁剪后的图片，其左上角所在的坐标填入x1，y1参数"`
	X2    string `json:"x2" comment:"以图片左上角（0,0），右下角（1,1）建立平面坐标系，经过裁剪后的图片，其右下角所在的坐标填入x2，y2参数"`
	Y2    string `json:"y2" comment:"以图片左上角（0,0），右下角（1,1）建立平面坐标系，经过裁剪后的图片，其右下角所在的坐标填入x2，y2参数"`
}
type CoverInfo struct {
	CropPercentList []CoverItem `json:"crop_percent_list" comment:"封面裁剪信息，。以图片左上角（0,0），右下角（1,1）建立平面坐标系，经过裁剪后的图片，其左上角所在的坐标填入x1，y1参数，右下角所在的坐标填入x2，y2参数"`
}
type ProductItem struct {
	ProductKey string `json:"product_key" comment:"商品key"`
}
type ProductInfo struct {
	FooterProductInfo ProductItem `json:"footer_product_info" comment:"文末插入商品相关信息"`
}
type ArticlesInfo struct {
	ArticleType        ArticleType  `json:"article_type" comment:"文章类型，分别有图文消息（news）、图片消息（newspic），不填默认为图文消息（news）"`
	Title              string       `json:"title" comment:"标题"`
	Author             string       `json:"author" comment:"作者"`
	Digest             string       `json:"digest" comment:"图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前54个字。"`
	Content            string       `json:"content" comment:"图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源"`
	ContentSourceUrl   string       `json:"content_source_url" comment:"图文消息的原文地址，即点击“阅读原文”后的URL"`
	ThumbMediaId       string       `json:"thumb_media_id" comment:"图文消息的封面图片素材id（必须是永久MediaID）"`
	NeedOpenComment    int          `json:"need_open_comment" comment:"是否打开评论，0不打开(默认)，1打开"`
	OnlyFansCanComment int          `json:"only_fans_can_comment" comment:"是否粉丝才可评论，0所有人可评论(默认)，1粉丝才可评论"`
	PicCrop2351        string       `json:"pic_crop_235_1" comment:"图文消息封面裁剪为2.35:1规格的坐标字段"`
	PicCrop11          string       `json:"pic_crop_1_1" comment:"图文消息封面裁剪为1:1规格的坐标字段，裁剪原理同pic_crop_235_1，裁剪后的图片必须符合规格要求。"`
	ImageInfo          *ImageInfo   `json:"image_info" comment:"图片消息里的图片相关信息，图片数量最多为20张，首张图片即为封面图"`
	CoverInfo          *CoverInfo   `json:"cover_info" comment:"图片消息的封面信息"`
	ProductInfo        *ProductInfo `json:"product_info" comment:"商品信息"`
}
type DraftUpdateRequest struct {
	MediaId  string       `json:"media_id" comment:"要修改的图文消息的id"`
	Index    int          `json:"index" comment:"要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0"`
	Articles ArticlesInfo `json:"articles" comment:"图文信息"`
}

func (p *DraftUpdateRequest) ToByte() []byte {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil
	}
	return marshal
}
