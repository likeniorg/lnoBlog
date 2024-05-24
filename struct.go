package main

import (
	"time"
)

type Model struct {
	ID        int       `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

/**
CREATE TABLE users(
	id int(10) unsigned  auto_increment primary key comment "用户id",
	username varchar(100) not null comment "用户名",
	password char(30) not null comment "用户密码",
	email varchar(255) not null comment  "用户邮箱",
	identify char(15) not null  comment "用户身份",
	create_at timestamp default current_timestamp() not null comment "创建时间",
	last_login timestamp  comment "最后登录时间"
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
**/
// 用户，对应的数据表：users
type User struct {
	Model
	Username  string    `validate:"required" gorm:"column:username"`
	Password  string    `validate:"required" gorm:"column:password"`
	Email     string    `validate:"required,email" gorm:"column:email"`
	Identify  string    `gorm:"column:identify"`
	LastLogin time.Time `gorm:"column:last_login"`
}

/**
CREATE TABLE articles(
	id int(10) unsigned  auto_increment comment "文章id",
	title varchar(255) not null comment "文章标题",
	content_md text not null comment "markdown格式存储文章",
	content_html text not null comment "md转换为html",
	category_id int(10) unsigned not null comment "版块id",
	tags varchar(100),
	author_id int(10) unsigned not null,
	read_num int(10) unsigned default 0,
	create_at timestamp default current_timestamp(),
	update_at timestamp,
	primary key(id),
	FOREIGN KEY(category_id) REFERENCES article_categories(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/
// 文章，对应的数据表名：articles
type Article struct {
	Model
	// 文章读取次数
	ReadCount int `gorm:"column:read_count"`
	// 文章标题
	Title string `gorm:"column:title"`
	// 文章markdown格式
	ContentMD string `gorm:"column:content_md"`
	// 文章HTML格式
	ContentHTML string `gorm:"column:content_html"`
	// 版块id
	CategoryID int `gorm:"column:category_id"`
	// 文章版块外键
	ArticleCategory ArticleCategory `gorm:"foreignKey:category_id"`
	// 文章标签
	Tags string `gorm:"column:tags"`
	// 文章作者ID
	AuthorID int `gorm:"column:author_id"`
}

/*
CREATE TABLE article_categories(
	id int(100) unsigned auto_increment,
	category varchar(100) not null,
	primary key(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/
// 文章类型，对应表
type ArticleCategory struct {
	ID       int    `gomr:"column:id;primarykey"`
	Category string `gorm:"column:category"`
}

/*
CREATE TABLE comments(
	id int(10) unsigned AUTO_INCREMENT PRIMARY KEY,
	user_id int(10) unsigned,
	article_id int(10) unsigned,
	comment_text text,
	create_at TIMESTAMP DEFAULT current_timestamp(),
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (article_id) REFERENCES articles(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/
// 评论，对应表：comments
type Comment struct {
	Model
	CommentText string  `gorm:"column:comment_text"`
	UserID      int     `gorm:"column:user_id"`
	ArticleID   int     `gorm:"column:article_id"`
	User        User    `gorm:"foreignKey:user_id"`
	Article     Article `gorm:"foreignKey:article_id"`
}

// 身份认证信息
type AuthInfo struct {
	// 用户ID
	Uid string
	// 用户名
	Username string
	// 用户身份
	Identify string
}

// 配置
type Conf struct {
	Domain       string
	ServerIP     string
	Key          string
	HTTPS        string
	DatabaseIP   string
	DatabasePort string
	DBuser       string
	DBpass       string
}
