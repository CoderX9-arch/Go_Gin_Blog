package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`         // 子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model"` // 父评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                                      // 父评论ID
	Content            string          `gorm:"size:256" json:"content"`                                // 评论内容
	DiggCount          int             `gorm:"size:8;default:0" json:"digg_count"`                     // 点赞数
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`                  // 评论数
	Article            ArticleModel    `gorm:"foreignKey:article_id" json:"-"`                         // 所属文章 关联的文章 注意这里，如果数据库一开始创建了文章，这里会报错，因为文章表还没有创建，所以这里需要先创建文章表，再创建评论表
	//Article   ArticleModel `json:"-"`          // 所属文章 关联的文章 注意这里，如果数据库一开始创建了文章，这里会报错，因为文章表还没有创建，所以这里需要先创建文章表，再创建评论表
	ArticleID uint      `json:"article_id"` // 所属文章ID
	User      UserModel `json:"user_model"` // 评论用户 关联的用户
	UserID    uint      `json:"user_id"`    // 评论用户ID
}
