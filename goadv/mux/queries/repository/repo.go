
package repository
//Repository represents generic interface for interacting with DB
import (uuid"github.com/satori/go.uuid"
"github.com/jinzhu/gorm")
type Repository interface {
    Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string,primaryKeyName string) error
    GetFirst(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error
    GetAll(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor)error
    GetAllTenant(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor)error
    GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error
	Count(uow *UnitOfWork, entity interface{},count interface{}) error
    Add(uow *UnitOfWork, out interface{}) error
    Update(uow *UnitOfWork, out interface{}) error
    Delete(uow *UnitOfWork, out interface{}) error
	First(uow *UnitOfWork, out interface{}, id uuid.UUID) error 
	//Count(uow *UnitOfWork, entity interface{},count interface{}) error
}

// UnitOfWork represents a connection
type UnitOfWork struct {
    DB        *gorm.DB
    committed bool
    readOnly  bool
}

// NewUnitOfWork creates new UnitOfWork
func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {
    if readOnly {
        return &UnitOfWork{DB: db.New(), committed: false, readOnly: true}
    }
    return &UnitOfWork{DB: db.New().Begin(), committed: false, readOnly: false}
}

// Complete marks end of unit of work
func (uow *UnitOfWork) Complete() {
    if !uow.committed && !uow.readOnly {
        uow.DB.Rollback()
    }
}

// Commit the transaction
func (uow *UnitOfWork) Commit() {
    if !uow.readOnly {
        uow.DB.Commit()
    }
    uow.committed = true
}

// GormRepository implements Repository
type GormRepository struct {
}

// NewRepository returns a new repository object
func NewRepository() Repository {
    return &GormRepository{}
}
func Paginate(limit int, offset int, count *int) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if out != nil && count != nil {
			if err := db.Debug().Model(out).Count(count).Error; err != nil {
				return db, err
			}
		}
		if limit != -1 {
			db = db.Limit(limit)
		}
		if offset > 0 {
			db = db.Offset(offset)
		}
		return db, nil
	}
}
// GetAll retrieves all the records for a specified entity and returns it  for pagination (same as get all)
func (repository *GormRepository) GetAllTenant(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB
	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}
	return db.Debug().Find(out).Error
}

// QueryProcessor allows to modify the query before it is executed
type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB,error)

//Filter will filter the results
func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where(condition, args...)
		return db, nil
	}
}
func (repository *GormRepository) First(uow *UnitOfWork, out interface{}, id uuid.UUID) error {
	db := uow.DB
	if err := db.First(out,id).Error; err != nil {
		return (err)
	}
	return nil
}
func (repository *GormRepository) GetFirst(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return (err)
			}
		}
	}
	if err := db.Debug().First(out).Error; err != nil {
		return (err)
	}
	return nil
}
// PreloadAssociations specified associations to be preloaded
func PreloadAssociations(preloadAssociations []string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		if preloadAssociations != nil {
			for _, association := range preloadAssociations {
				db = db.Preload(association)
			}
		}
		return db, nil
	}
}

// Get a record for specified entity with specific id
func (repository *GormRepository) Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string,primaryKeyName string) error {
    db := uow.DB
    for _, association := range preloadAssociations {
        db = db.Preload(association)
    }
    return db.Debug().First(out, primaryKeyName+" = ?", id).Error
}
//count specified entity
func (repository *GormRepository) Count(uow *UnitOfWork, entity interface{},count interface{}) error {
    return uow.DB.Debug().Model(entity).Count(count).Error
}
func Select(tablename string,value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Table(tablename).Select(value)
		return db, nil
	}
}
func ShowIndex(value string,tablename interface{} )QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Raw(value).Scan(&tablename)
		return db, nil
	}
}
// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, queryProcessors []QueryProcessor) error {
	db := uow.DB

	if queryProcessors != nil {
		var err error
		for _, queryProcessor := range queryProcessors {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}
	if err := db.Debug().Find(out).Error; err != nil {
		return err
	}
	return nil
}

// GetAllForTenant returns all objects of specified tenantID
func (repository *GormRepository) GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error {
    db := uow.DB
    for _, association := range preloadAssociations {
        db = db.Preload(association)
    }
    return db.Where("tenantID = ?", tenantID).Find(out).Error
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {
    return uow.DB.Debug().Create(entity).Error
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
    return uow.DB.Debug().Model(entity).Update(entity).Error
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}) error {
    return uow.DB.Delete(entity).Error
}
func Group(field string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Group(field)
		return db, nil
	}
}
func Having(value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Having(value)
		return db, nil
	}
}
func Order(value interface{}, reorder bool) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Order(value, reorder)
		return db, nil
	}
}
func Joins(value string) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Joins(value)
		return db, nil
	}
}