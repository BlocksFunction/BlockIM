package utils

import (
	"database/sql"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql" // 只保留MySQL驱动
	"gopkg.in/yaml.v2"
)

// Config 定义配置结构
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

// ColumnDefinition 定义表列结构
type ColumnDefinition struct {
	Name     string
	Type     string
	Nullable bool
	Primary  bool
	Unique   bool
	Default  string
}

// Database 封装MySQL数据库连接和操作
type Database struct {
	db *sql.DB
}

// 错误常量
const (
	MySQLDuplicateError = 1062
)

// UseDatabase 创建新的MySQL数据库实例
func UseDatabase(configPath string) (*Database, error) {
	// 读取配置文件
	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("无法读取配置文件: %w", err)
	}

	// 解析 YAML 配置
	var config Config
	if err = yaml.Unmarshal(configFile, &config); err != nil {
		return nil, fmt.Errorf("无法转译YAML: %w", err)
	}

	// 创建MySQL数据库连接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("无法连接数据库: %w", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// 测试数据库连接
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("无法连接数据库: %w", err)
	}

	return &Database{db: db}, nil
}

// Close 关闭数据库连接
func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

// IsDuplicateError 检查是否为唯一约束冲突错误
func (d *Database) IsDuplicateError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), fmt.Sprintf("Error %d", MySQLDuplicateError))
}

// Exec 执行SQL命令
func (d *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.db == nil {
		return nil, fmt.Errorf("数据库未初始化")
	}

	return d.db.Exec(query, args...)
}

// Query 执行查询操作
func (d *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if d.db == nil {
		return nil, fmt.Errorf("数据库未初始化")
	}

	return d.db.Query(query, args...)
}

// QueryRow 执行单行查询
func (d *Database) QueryRow(query string, args ...interface{}) *sql.Row {
	if d.db == nil {
		// 返回一个会报错的Row对象
		return &sql.Row{}
	}

	return d.db.QueryRow(query, args...)
}

// CreateTable 创建新表
func (d *Database) CreateTable(tableName string, columns []ColumnDefinition) error {
	if len(columns) == 0 {
		return fmt.Errorf("没有为创建表提供列")
	}

	// 构建列定义SQL
	columnDefs := make([]string, 0, len(columns))

	for _, col := range columns {
		def := fmt.Sprintf("%s %s", col.Name, col.Type)

		if col.Primary {
			def += " AUTO_INCREMENT PRIMARY KEY"
		} else {
			if !col.Nullable {
				def += " NOT NULL"
			}

			if col.Unique {
				def += " UNIQUE"
			}

			if col.Default != "" {
				def += " DEFAULT " + col.Default
			}
		}

		columnDefs = append(columnDefs, def)
	}

	// 构建完整的CREATE TABLE语句
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n  %s\n)", tableName, strings.Join(columnDefs, ",\n  "))
	_, err := d.Exec(query)
	return err
}

// DeleteTable 删除表
func (d *Database) DeleteTable(tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err := d.Exec(query)
	return err
}

// Insert 执行插入操作并返回插入的ID
func (d *Database) Insert(table string, data map[string]interface{}) (int64, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("没有为插入操作提供数据")
	}

	// 准备列名和值
	columns := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data))

	// 为了确保顺序一致，对键进行排序
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	placeholders := make([]string, 0, len(data))
	for _, k := range keys {
		columns = append(columns, k)
		values = append(values, data[k])
		placeholders = append(placeholders, "?")
	}

	// 构建SQL语句
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "))

	result, err := d.Exec(query, values...)
	if err != nil {
		return 0, err
	}

	// 获取最后插入的ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("获取插入ID失败: %w", err)
	}

	return id, nil
}

// Delete 执行删除操作
func (d *Database) Delete(table string, where map[string]interface{}) (sql.Result, error) {
	if len(where) == 0 {
		return nil, fmt.Errorf("为了安全起见，不允许删除不带 where 子句的删除操作")
	}

	// 准备条件和值
	conditions := make([]string, 0, len(where))
	values := make([]interface{}, 0, len(where))

	keys := make([]string, 0, len(where))
	for k := range where {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		conditions = append(conditions, fmt.Sprintf("%s = ?", k))
		values = append(values, where[k])
	}

	// 构建SQL语句
	query := fmt.Sprintf("DELETE FROM %s WHERE %s",
		table,
		strings.Join(conditions, " AND "))

	return d.Exec(query, values...)
}

// Update 执行更新操作
func (d *Database) Update(table string, data map[string]interface{}, where map[string]interface{}) (sql.Result, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("未提供用于更新的数据")
	}
	if len(where) == 0 {
		return nil, fmt.Errorf("为了安全起见，不允许不带 where 子句的更新操作")
	}

	// 准备更新字段
	setClauses := make([]string, 0, len(data))
	values := make([]interface{}, 0, len(data)+len(where))

	// 处理data部分
	dataKeys := make([]string, 0, len(data))
	for k := range data {
		dataKeys = append(dataKeys, k)
	}
	sort.Strings(dataKeys)

	for _, k := range dataKeys {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", k))
		values = append(values, data[k])
	}

	// 处理where条件
	whereClauses := make([]string, 0, len(where))
	whereKeys := make([]string, 0, len(where))
	for k := range where {
		whereKeys = append(whereKeys, k)
	}
	sort.Strings(whereKeys)

	for _, k := range whereKeys {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = ?", k))
		values = append(values, where[k])
	}

	// 构建SQL语句
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s",
		table,
		strings.Join(setClauses, ", "),
		strings.Join(whereClauses, " AND "))

	return d.Exec(query, values...)
}

// Select 执行查询操作
func (d *Database) Select(table string, columns []string, where map[string]interface{}) (*sql.Rows, error) {
	// 构建SELECT子句
	cols := "*"
	if len(columns) > 0 {
		cols = strings.Join(columns, ", ")
	}

	query := fmt.Sprintf("SELECT %s FROM %s", cols, table)

	// 准备WHERE条件和值
	var values []interface{}
	if len(where) > 0 {
		conditions := make([]string, 0, len(where))
		values = make([]interface{}, 0, len(where))

		keys := make([]string, 0, len(where))
		for k := range where {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			conditions = append(conditions, fmt.Sprintf("%s = ?", k))
			values = append(values, where[k])
		}

		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	return d.Query(query, values...)
}
