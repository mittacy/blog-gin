package service

import (
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/log"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
)

type IArticleService interface {
	CreateArticle(model.Article) error
	DeleteArticle(int) error
	UpdateArticle(model.Article) error
	GetArticleByID(int) (*model.Article, error)
	GetArticlesByPage(int, int) ([]model.Article, int, error)
	GetRecent() ([]model.Article, error)
}

func NewArticleService(repo repository.IArticleRepository) IArticleService {
	return &ArticleService{repo}
}

type ArticleService struct {
	ArticleRepository repository.IArticleRepository
}

func (as *ArticleService) CreateArticle(article model.Article) error {
	// 1. 更新到数据库
	if err := as.ArticleRepository.Add(&article); err != nil {
		return err
	}
	// 2. 更新到缓存区
	cache.AddArticleCache(article)
	// 3. 更新分类缓存
	cache.UpdateCategoryCacheIncr(article.CategoryID)
	return nil
}

func (as *ArticleService) DeleteArticle(id int) error {
	// 1. 更新到数据库
	if err := as.ArticleRepository.Delete(id); err != nil {
		return err
	}
	// 2. 更新到缓存区
	article, isExist := cache.GetArticleCacheByID(uint32(id))	// 缓存区删除文章前获取文章信息
	cache.DeleteArticleCache(uint32(id))
	// 3. 更新分类缓存
	go func() {
		if isExist {
			cache.UpdateCategoryCacheDecr(article.CategoryID)
		} else {
			if err := cache.InitCategoryCache(); err != nil {
				log.RecordErr(err)
			}
		}
	}()
	return nil
}

func (as *ArticleService) UpdateArticle(article model.Article) error {
	// 1. 更新到数据库
	if err := as.ArticleRepository.Update(&article); err != nil {
		return err
	}
	// 2. 更新到缓存区
	articleOld, isExist := cache.GetArticleCacheByID(uint32(article.ID))	// 缓存区更新文章前获取文章信息
	cache.UpdateArticleCache(article)
	// 3. 更新缓存器分类文章数量
	go func() {
		if isExist {
			if articleOld.CategoryID != article.CategoryID {
				cache.UpdateCategoryCacheDecr(articleOld.CategoryID)
				cache.UpdateCategoryCacheIncr(article.CategoryID)
			}
		} else {
			if err := cache.InitCategoryCache(); err != nil {
				log.RecordErr(err)
			}
		}
	}()
	return nil
}

func (as *ArticleService) GetArticleByID(id int) (*model.Article, error) {
	// 1. 数据库取数据
	return as.ArticleRepository.SelectByID(id)
}

func (as *ArticleService) GetArticlesByPage(page, onePageArticleNum int) (articles []model.Article, articleCount int, err error) {
	// 1. 缓存器取数据
	isExist := false
	articles, articleCount, isExist = cache.GetArticleCacheByPage(page, onePageArticleNum)
	if isExist {
		return
	}
	// 2. 数据库取数据
	articles, articleCount, err = as.ArticleRepository.SelectByPage(page, onePageArticleNum)
	if err != nil {
		return
	}
	return
}

func (as *ArticleService) GetRecent() (articles []model.Article, err error) {
	// 1. 缓存器取数据
	isExist := false
	articles, isExist = cache.GetRecentArticleCache()
	if isExist {
		return
	}
	// 2. 从数据库取数据
	articles, err = as.ArticleRepository.SelectRecent()
	return
}


