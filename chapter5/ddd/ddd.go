package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	minTitleLength = 1
	maxTitleLength = 50
)

var (
	// 给出无效标题时使用的错误
	ErrInvalidTitle  = errors.New("tab: could not use invalid title")
	ErrTitleTooShort = fmt.Errorf("%s: min length allowed is %d", ErrInvalidTitle, minTitleLength)
	ErrTitleTooLong  = fmt.Errorf("%s: max length allowed is %d", ErrInvalidTitle, maxTitleLength)
)

// Title 代表标签标题
type Title string

// NewTitle 返回标题和错误
func NewTitle(d string) (Title, error) {
	switch l := len(strings.TrimSpace(d)); {
	case l < minTitleLength:
		return "", ErrTitleTooShort
	case l > maxTitleLength:
		return "", ErrTitleTooLong
	default:
		return Title(d), nil
	}
}

// String 返回标题的字符串表示
func (t Title) String() string {
	return string(t)
}

// Equals 如果标题相等，则返回 true
func (t Title) Equals(t2 Title) bool {
	return t.String() == t2.String()
}

// 添加标签请求
type addTabReq struct {
	Title Title `json:"tab_title"`
}

// 解码JSON请求正文
func (r *addTabReq) UnmarshalJSON(data []byte) error {
	var req addTabReq
	var err error
	if err = json.Unmarshal(data, &req); err != nil {
		return err
	}
	if r.Title, err = NewTitle(req.Title.String()); err != nil {
		return err
	}
	return nil
}

// Tab 代表一个标签
type Tab struct {
	ID          int
	Title       Title
	Description string
	Icon        string
	Link        string
	Created     time.Time
	Updated     time.Time
}

// New 返回第一次创建的标签
func NewTab(id int, title Title, description string, icon string, link string) *Tab {
	return &Tab{
		ID:          id,
		Title:       title,
		Description: description,
		Icon:        icon,
		Link:        link,
		Created:     time.Now(),
	}
}

// Update 使用新属性更新标签
func (t *Tab) Update(title Title, description string, icon string, link string) {
	t.Title = title
	t.Description = description
	t.Icon = icon
	t.Link = link
	t.Updated = time.Now()
}

var (
	//存储库返回的错误
	ErrRepoNextID = errors.New("tab: could not return next id")
	ErrRepoList   = errors.New("tab: could not list")
	ErrNotFound   = errors.New("tab: could not find")
	ErrRepoGet    = errors.New("tab: could not get")
	ErrRepoAdd    = errors.New("tab: could not add")
	ErrRepoRemove = errors.New("tab: could not remove")
)

type ID int

type Repo interface {
	//NextID 返回下一个空闲 ID 并在失败时返回错误
	NextID() (ID, error)
	// 列表返回一个选项卡切片并在失败的情况下返回一个错误
	List() ([]*Tab, error)
	// Find 返回一个 tab 或者 nil ，如果它没有找到并且在失败的情况下，则返回一个错误
	Find(ID) (*Tab, error)
	// 如果未找到或失败，则获取返回选项卡和错误
	Get(ID) (*Tab, error)
	// 添加持久化选项卡（已经存在或不存在）并在失败时返回错误
	Add(*Tab) error
	// 删除选项卡并返回并在未找到或失败的情况下出错
	Remove(ID) error
}

// Collection 代表一个集合
type Collection struct {
	ID      int
	Name    string
	Tabs    []*Tab
	Created time.Time
	Updated time.Time
}

// New 返回第一次创建的集合
func NewCollection(id int, name string) *Collection {
	return &Collection{
		ID:      id,
		Name:    name,
		Tabs:    make([]*Tab, 0),
		Created: time.Now(),
	}
}

// Rename重命名集合
func (c *Collection) Rename(name string) {
	c.Name = name
	c.Updated = time.Now()
}

// AddTabs 将Tab添加到集合
func (c *Collection) AddTabs(tabs ...*Tab) {
	c.Tabs = append(c.Tabs, tabs...)
	c.Updated = time.Now()
}

// RemoveTab 如果标签存在，则删除它
func (c *Collection) RemoveTab(id int) bool {
	for i, t := range c.Tabs {
		if t.ID == id {
			c.Tabs[i] = c.Tabs[len(c.Tabs)-1]
			c.Tabs[len(c.Tabs)-1] = nil
			c.Tabs = c.Tabs[:len(c.Tabs)-1]
			c.Updated = time.Now()
			return true
		}
	}
	return false
}

// FindTab 如果存在则返回一个标签
func (c *Collection) FindTab(id int) (*Tab, bool) {
	for _, t := range c.Tabs {
		if t.ID == id {
			return t, true
		}
	}
	return nil, false
}

// UpdateTab 更新标签（如果存在）
func (c *Collection) UpdateTab(t *Tab) bool {
	for i, tb := range c.Tabs {
		if tb.ID == t.ID {
			c.Tabs[i] = t
			c.Updated = time.Now()
			return true
		}
	}
	return false
}

// 工作区
type Workspace struct {
	ID          int
	Name        string
	CustomerID  int
	Collections []*Collection
	Created     time.Time
	Updated     time.Time
}

// 返回第一次创建的工作区
func NewWorkspace(id int, name string, customerID int) *Workspace {
	return &Workspace{
		ID:          id,
		Name:        name,
		CustomerID:  customerID,
		Collections: make([]*Collection, 0),
		Created:     time.Now(),
	}
}

// 更改工作区的名称
func (w *Workspace) Rename(name string) {
	w.Name = name
	w.Updated = time.Now()
}

// 添加一个集合
func (w *Workspace) AddCollections(collections ...*Collection) {
	w.Collections = append(w.Collections, collections...)
	w.Updated = time.Now()
}

// 如果集合存在，则删除它
func (w *Workspace) RemoveCollection(id int) bool {
	for i, coll := range w.Collections {
		if coll.ID == id {
			w.Collections[i] = w.Collections[len(w.Collections)-1]
			w.Collections[len(w.Collections)-1] = nil
			w.Collections = w.Collections[:len(w.Collections)-1]
			w.Updated = time.Now()
			return true
		}
	}
	return false
}

// 重命名集合（如果存在）
func (w *Workspace) RenameCollection(id int, name string) bool {
	for _, coll := range w.Collections {
		if coll.ID == id {
			coll.Rename(name)
			w.Updated = time.Now()
			return true
		}
	}
	return false
}

// 这里是一个 MySQL 实现
type MysqlRepo struct {
}

func (r *MysqlRepo) Add(t *Tab) error {
	return fmt.Errorf("error: %s a more detailed reason here", ErrRepoAdd)
}

func main() {
}
