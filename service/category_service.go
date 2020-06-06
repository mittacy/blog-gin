package service

import (
	"github.com/crazychat/blog-gin/cache"
	"github.com/crazychat/blog-gin/log"
	"github.com/crazychat/blog-gin/model"
	"github.com/crazychat/blog-gin/repository"
)

type ICategoryService interface {
	CreateCategory(model.Category) error
	DeleteCategory(int) error
	UpdateCategory(model.Category) error
	GetCategoryByID(int) (*model.Category, error)
	GetCategories() ([]model.Category, error)
	GetCategoriesByPage(int, int) ([]model.Category, int, error)
}

func NewCategoryService(repository repository.ICategoryRepository) ICategoryService {
	return &CategoryService{repository}
}

type CategoryService struct {
	CategoryRepository repository.ICategoryRepository
}
// CreateCategory 添加分类
func (cs *CategoryService) CreateCategory(cate model.Category) error {
	// 1. 更新到数据库
	if err := cs.CategoryRepository.Add(cate); err != nil {
		return err
	}
	// 2. 更新到缓存区
	cache.AddCategoryCache(cate)
	return nil
}
// DeleteCategory 删除分类
func (cs *CategoryService) DeleteCategory(id int) error {
	// 1. 更新到数据库
	if err := cs.CategoryRepository.Delete(id); err != nil {
		return err
	}
	// 2. 更新到缓存区
	cache.DeleteCategoryCache(uint32(id))
	return nil
}
// UpdateCategory 更新分类信息
func (cs *CategoryService) UpdateCategory(cate model.Category) error {
	// 1. 更新到数据库
	if err := cs.CategoryRepository.Update(cate); err != nil {
		return err
	}
	// 2. 更新到缓存区
	cache.UpdateCategoryCache(cate)
	return nil
}
// GetCategoryByID 通过id获取分类信息
func (cs *CategoryService) GetCategoryByID(id int) (category *model.Category, err error) {
	isExist := false
	// 1. 缓存器取数据
	category, isExist = cache.GetCategoryCacheByID(uint32(id))
	if !isExist {
		// 2. 数据库取数据
		category, err = cs.CategoryRepository.SelectByID(id)
		if err != nil {
			return
		}
		// 清空cate缓存，重新缓存分类
		go func() {
			categories, err := cs.CategoryRepository.Select()
			if err != nil {
				log.RecordErr(err)
			}
			cache.InitCategoryCache(categories)
		}()
	}
	return
}
// GetCategories 获取全部分类信息
func (cs *CategoryService) GetCategories() (categories []model.Category, err error) {
	isExist := false
	// 1. 缓存取数据
	categories, isExist = cache.GetCategoriesCache()
	if !isExist {
		// 2. 数据库取数据
		categories, err = cs.CategoryRepository.Select()
		if err != nil {
			return
		}
		// 缓存分类
		cache.InitCategoryCache(categories)
	}
	return
}
// GetCategoriesByPage 分页获取分类
func (cs *CategoryService) GetCategoriesByPage(page, onePageCategoryNum int) (categories []model.Category, categoryCount int, err error) {
	isExist := false
	// 1. 缓存取数据
	categories, categoryCount, isExist = cache.GetCategoriesCacheByPage(page)
	if !isExist {
		// 2. 数据库取数据
		categories, categoryCount, err = cs.CategoryRepository.SelectByPage(page, onePageCategoryNum)
		if err != nil {
			return
		}
		// cate缓存器数据有问题，清空，重新缓存分类
		go func() {
			categories, err := cs.CategoryRepository.Select()
			if err != nil {
				log.RecordErr(err)
			}
			cache.InitCategoryCache(categories)
		}()
	}
	return
}

