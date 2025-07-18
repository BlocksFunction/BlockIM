package database

import (
	"Backed/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type Article struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	Excerpt  string    `json:"excerpt"`
	Author   string    `json:"author"`
	Date     time.Time `json:"date"`
	ReadTime int       `json:"read_time"`
	Likes    int       `json:"likes"`
	Comments int       `json:"comments"`
	Views    int       `json:"views"`
	Category string    `json:"category"`
	Tags     []string  `json:"tags"`
	Featured bool      `json:"featured"`
	Content  string    `json:"content"`
}

type ArticleData struct {
	db *utils.Database
}

func UseArticleData() (*ArticleData, error) {
	data, err := utils.UseDatabase("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("无法打开数据库: %v", err)
	}

	articleTableColumns := []utils.ColumnDefinition{
		{Name: "id", Type: "INT", Primary: true},
		{Name: "title", Type: "VARCHAR(255)", Nullable: false, Unique: true},
		{Name: "excerpt", Type: "TEXT"},
		{Name: "author", Type: "VARCHAR(255)", Nullable: false},
		{Name: "date", Type: "TIMESTAMP", Default: "CURRENT_TIMESTAMP"},
		{Name: "read_time", Type: "INT", Default: "0"},
		{Name: "likes", Type: "INT", Default: "0"},
		{Name: "comments", Type: "INT", Default: "0"},
		{Name: "views", Type: "INT", Default: "0"},
		{Name: "category", Type: "VARCHAR(100)"},
		{Name: "tags", Type: "JSON"},
		{Name: "featured", Type: "BOOLEAN", Default: "false"},
	}

	err = data.CreateTable("articles", articleTableColumns)
	if err != nil {
		return nil, fmt.Errorf("无法创建文章数据表: %v", err)
	}

	if err := createFullTextIndex(data); err != nil {
		fmt.Errorf("无法创建全文索引: %v", err)
	}

	return &ArticleData{db: data}, nil
}

// createFullTextIndex 创建完整索引
func createFullTextIndex(db *utils.Database) error {
	// 检查索引是否已存在
	checkQuery := `
        SELECT COUNT(*) FROM information_schema.statistics 
        WHERE table_schema = DATABASE() 
        AND table_name = 'articles' 
        AND index_name = 'idx_fulltext_search'
    `

	var count int
	row := db.QueryRow(checkQuery)
	if err := row.Scan(&count); err != nil {
		return fmt.Errorf("检查索引失败: %w", err)
	}

	// 如果索引不存在，则创建
	if count == 0 {
		createQuery := `
            ALTER TABLE articles 
            ADD FULLTEXT INDEX idx_fulltext_search (title, excerpt, content)
            WITH PARSER ngram
        `

		if _, err := db.Exec(createQuery); err != nil {
			return fmt.Errorf("创建全文索引失败: %w", err)
		}
	}

	return nil
}

func (a *ArticleData) Close() error {
	if a.db != nil {
		return a.db.Close()
	}
	return nil
}

func (a *ArticleData) IsDuplicateError(err error) bool {
	return a.db.IsDuplicateError(err)
}

func (a *ArticleData) GetArticleByID(id int64) (*Article, error) {
	where := map[string]interface{}{
		"id": id,
	}

	rows, err := a.db.Select("articles", nil, where)
	if err != nil {
		return nil, fmt.Errorf("查询失败: %w", err)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	var article Article
	var tagsJSON string

	err = rows.Scan(
		&article.ID,
		&article.Title,
		&article.Excerpt,
		&article.Author,
		&article.Date,
		&article.ReadTime,
		&article.Likes,
		&article.Comments,
		&article.Views,
		&article.Category,
		&tagsJSON,
		&article.Featured,
	)

	if err != nil {
		return nil, fmt.Errorf("解析数据失败: %w", err)
	}

	// 解析JSON
	if err := json.Unmarshal([]byte(tagsJSON), &article.Tags); err != nil {
		return nil, fmt.Errorf("解析标签失败: %w", err)
	}

	return &article, nil
}

func (a *ArticleData) InsertArticle(article *Article) (*Article, error) {
	if article.Title == "" {
		return nil, fmt.Errorf("标题不能为空")
	}
	// 使用当前时间
	if article.Date.IsZero() {
		article.Date = time.Now()
	}

	// 序列化标签为JSON
	tagsJSON, err := json.Marshal(article.Tags)
	if err != nil {
		return nil, fmt.Errorf("序列化标签失败: %w", err)
	}

	data := map[string]interface{}{
		"title":     article.Title,
		"excerpt":   article.Excerpt,
		"author":    article.Author,
		"date":      article.Date,
		"read_time": article.ReadTime,
		"likes":     article.Likes,
		"comments":  article.Comments,
		"views":     article.Views,
		"category":  article.Category,
		"tags":      string(tagsJSON),
		"featured":  article.Featured,
	}

	var id int64
	id, err = a.db.Insert("articles", data)
	if err != nil {
		return nil, fmt.Errorf("插入文章失败: %w", err)
	}

	article.ID = id
	return article, nil
}

func (a *ArticleData) UpdateArticle(article *Article) error {
	if article.ID == 0 {
		return fmt.Errorf("无效的文章ID")
	}

	// 序列化标签为JSON
	tagsJSON, err := json.Marshal(article.Tags)
	if err != nil {
		return fmt.Errorf("序列化标签失败: %w", err)
	}

	data := map[string]interface{}{
		"title":     article.Title,
		"excerpt":   article.Excerpt,
		"author":    article.Author,
		"date":      article.Date,
		"read_time": article.ReadTime,
		"likes":     article.Likes,
		"comments":  article.Comments,
		"views":     article.Views,
		"category":  article.Category,
		"tags":      string(tagsJSON),
		"featured":  article.Featured,
	}

	where := map[string]interface{}{
		"id": article.ID,
	}

	_, err = a.db.Update("articles", data, where)
	return err
}

func (a *ArticleData) DeleteArticle(id int64) (bool, error) {
	if id == 0 {
		return false, fmt.Errorf("无效的文章ID")
	}

	where := map[string]interface{}{
		"id": id,
	}

	result, err := a.db.Delete("articles", where)
	if err != nil {
		return false, fmt.Errorf("删除文章失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("获取受影响行数失败: %w", err)
	}

	return rowsAffected > 0, nil
}

func (a *ArticleData) IncrementViews(id int64) error {
	// 确保原子操作
	query := "UPDATE articles SET views = views + 1 WHERE id = ?"
	_, err := a.db.Exec(query, id)
	return err
}

func (a *ArticleData) ListArticles(limit, offset int, category, tag string, featuredOnly bool) ([]*Article, error) {
	// 构建查询查询
	where := make(map[string]interface{})

	if category != "" {
		where["category"] = category
	}

	if featuredOnly {
		where["featured"] = true
	}

	// 处理标签过滤
	tagFilter := ""
	if tag != "" {
		tagFilter = " AND JSON_CONTAINS(tags, JSON_ARRAY(?))"
	}

	// 构建基础查询
	query := fmt.Sprintf("SELECT * FROM articles WHERE 1=1%s ORDER BY date DESC LIMIT ? OFFSET ?", tagFilter)

	// 准备参数
	params := []interface{}{}
	if tag != "" {
		params = append(params, tag)
	}
	params = append(params, limit, offset)

	// 执行!
	rows, err := a.db.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("查询文章列表失败: %w", err)
	}
	defer rows.Close()

	return a.scanArticles(rows)
}

func (a *ArticleData) SearchArticles(query string, limit int) ([]*Article, error) {
	searchQuery := `
			SELECT * FROM articles 
			WHERE MATCH(title, excerpt, content) AGAINST(? IN NATURAL LANGUAGE MODE)
			ORDER BY MATCH(title, excerpt, content) AGAINST(? IN NATURAL LANGUAGE MODE) DESC
			LIMIT ?
		`

	// 准备参数
	params := []interface{}{query, query, limit}

	rows, err := a.db.Query(searchQuery, params...)
	if err != nil {
		return nil, fmt.Errorf("搜索文章失败: %w", err)
	}
	defer rows.Close()

	return a.scanArticles(rows)
}

// scanArticles 从查询结果中扫描文章列表
func (a *ArticleData) scanArticles(rows *sql.Rows) ([]*Article, error) {
	var articles []*Article

	for rows.Next() {
		var article Article
		var tagsJSON string

		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Excerpt,
			&article.Author,
			&article.Date,
			&article.ReadTime,
			&article.Likes,
			&article.Comments,
			&article.Views,
			&article.Category,
			&tagsJSON,
			&article.Featured,
		)

		if err != nil {
			return nil, fmt.Errorf("解析文章数据失败: %w", err)
		}

		if err := json.Unmarshal([]byte(tagsJSON), &article.Tags); err != nil {
			return nil, fmt.Errorf("解析标签失败: %w", err)
		}

		articles = append(articles, &article)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历结果集失败: %w", err)
	}

	return articles, nil
}

func (a *ArticleData) GetRecentArticles(limit int) ([]*Article, error) {
	query := "SELECT * FROM articles ORDER BY date DESC LIMIT ?"
	rows, err := a.db.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("获取最新文章失败: %w", err)
	}
	defer rows.Close()

	return a.scanArticles(rows)
}

func (a *ArticleData) GetFeaturedArticles(limit int) ([]*Article, error) {
	where := map[string]interface{}{
		"featured": true,
	}

	rows, err := a.db.Select("articles", nil, where)
	if err != nil {
		return nil, fmt.Errorf("获取精选文章失败: %w", err)
	}
	defer rows.Close()

	return a.scanArticles(rows)
}

func (a *ArticleData) GetArticlesByCategory(category string, limit int) ([]*Article, error) {
	where := map[string]interface{}{
		"category": category,
	}

	rows, err := a.db.Select("articles", nil, where)
	if err != nil {
		return nil, fmt.Errorf("获取分类文章失败: %w", err)
	}
	defer rows.Close()

	return a.scanArticles(rows)
}

func (a *ArticleData) IncrementLikes(id int64) error {
	query := "UPDATE articles SET likes = likes + 1 WHERE id = ?"
	_, err := a.db.Exec(query, id)
	return err
}

func (a *ArticleData) IncrementComments(id int64) error {
	query := "UPDATE articles SET comments = comments + 1 WHERE id = ?"
	_, err := a.db.Exec(query, id)
	return err
}
