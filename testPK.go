package tests

import (
	"errors"
	"testing"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type IntId struct {
	Id   int `xorm:"pk autoincr"`
	Name string
}

type Int32Id struct {
	Id   int32 `xorm:"pk autoincr"`
	Name string
}

type UintId struct {
	Id   uint `xorm:"pk autoincr"`
	Name string
}

type Uint32Id struct {
	Id   uint32 `xorm:"pk autoincr"`
	Name string
}

type Uint64Id struct {
	Id   uint64 `xorm:"pk autoincr"`
	Name string
}

type StringPK struct {
	Id   string `xorm:"pk notnull"`
	Name string
}

func testIntId(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&IntId{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&IntId{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	cnt, err := engine.Insert(&IntId{Name: "test"})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}

	bean := new(IntId)
	has, err := engine.Get(bean)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if !has {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans := make([]IntId, 0)
	err = engine.Find(&beans)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans2 := make(map[int]IntId)
	err = engine.Find(&beans2)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans2) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	cnt, err = engine.Id(bean.Id).Delete(&IntId{})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}
}

func testInt32Id(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&Int32Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&Int32Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	cnt, err := engine.Insert(&Int32Id{Name: "test"})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}

	bean := new(Int32Id)
	has, err := engine.Get(bean)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if !has {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans := make([]Int32Id, 0)
	err = engine.Find(&beans)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans2 := make(map[int32]Int32Id, 0)
	err = engine.Find(&beans2)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans2) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	cnt, err = engine.Id(bean.Id).Delete(&Int32Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}
}

func testUintId(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&UintId{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&UintId{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	cnt, err := engine.Insert(&UintId{Name: "test"})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}

	bean := new(UintId)
	has, err := engine.Get(bean)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if !has {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans := make([]UintId, 0)
	err = engine.Find(&beans)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans2 := make(map[uint]UintId, 0)
	err = engine.Find(&beans2)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans2) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	cnt, err = engine.Id(bean.Id).Delete(&UintId{})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}
}

func testUint32Id(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&Uint32Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&Uint32Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	cnt, err := engine.Insert(&Uint32Id{Name: "test"})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}

	bean := new(Uint32Id)
	has, err := engine.Get(bean)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if !has {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans := make([]Uint32Id, 0)
	err = engine.Find(&beans)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans2 := make(map[uint32]Uint32Id, 0)
	err = engine.Find(&beans2)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans2) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	cnt, err = engine.Id(bean.Id).Delete(&Uint32Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}
}

func testUint64Id(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&Uint64Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&Uint64Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	idbean := &Uint64Id{Name: "test"}
	cnt, err := engine.Insert(idbean)
	if err != nil {
		t.Error(err)
		panic(err)
	}

	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}

	bean := new(Uint64Id)
	has, err := engine.Get(bean)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if !has {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	if bean.Id != idbean.Id {
		panic(errors.New("should be equal"))
	}

	beans := make([]Uint64Id, 0)
	err = engine.Find(&beans)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	if *bean != beans[0] {
		panic(errors.New("should be equal"))
	}

	beans2 := make(map[uint64]Uint64Id, 0)
	err = engine.Find(&beans2)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans2) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	if *bean != beans2[bean.Id] {
		panic(errors.New("should be equal"))
	}

	cnt, err = engine.Id(bean.Id).Delete(&Uint64Id{})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}
}

func testStringPK(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&StringPK{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&StringPK{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	cnt, err := engine.Insert(&StringPK{Id: "1-1-2", Name: "test"})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}

	bean := new(StringPK)
	has, err := engine.Get(bean)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if !has {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans := make([]StringPK, 0)
	err = engine.Find(&beans)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	beans2 := make(map[string]StringPK)
	err = engine.Find(&beans2)
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if len(beans2) != 1 {
		err = errors.New("get count should be one")
		t.Error(err)
		panic(err)
	}

	cnt, err = engine.Id(bean.Id).Delete(&StringPK{})
	if err != nil {
		t.Error(err)
		panic(err)
	}
	if cnt != 1 {
		err = errors.New("insert count should be one")
		t.Error(err)
		panic(err)
	}
}

type CompositeKey struct {
	Id1       int64 `xorm:"id1 pk"`
	Id2       int64 `xorm:"id2 pk"`
	UpdateStr string
}

func testCompositeKey(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&CompositeKey{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&CompositeKey{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	cnt, err := engine.Insert(&CompositeKey{11, 22, ""})
	if err != nil {
		t.Error(err)
	} else if cnt != 1 {
		t.Error(errors.New("failed to insert CompositeKey{11, 22}"))
	}

	cnt, err = engine.Insert(&CompositeKey{11, 22, ""})
	if err == nil || cnt == 1 {
		t.Error(errors.New("inserted CompositeKey{11, 22}"))
	}

	var compositeKeyVal CompositeKey
	has, err := engine.Id(core.PK{11, 22}).Get(&compositeKeyVal)
	if err != nil {
		t.Error(err)
	} else if !has {
		t.Error(errors.New("can't get CompositeKey{11, 22}"))
	}

	var compositeKeyVal2 CompositeKey
	// test passing PK ptr, this test seem failed withCache
	has, err = engine.Id(&core.PK{11, 22}).Get(&compositeKeyVal2)
	if err != nil {
		t.Error(err)
	} else if !has {
		t.Error(errors.New("can't get CompositeKey{11, 22}"))
	}

	if compositeKeyVal != compositeKeyVal2 {
		t.Error(errors.New("should be equal"))
	}

	var cps = make([]CompositeKey, 0)
	err = engine.Find(&cps)
	if err != nil {
		t.Error(err)
	}
	if len(cps) != 1 {
		t.Error(errors.New("should has one record"))
	}
	if cps[0] != compositeKeyVal {
		t.Error(errors.New("should be equal"))
	}

	cnt, err = engine.Insert(&CompositeKey{22, 22, ""})
	if err != nil {
		t.Error(err)
	} else if cnt != 1 {
		t.Error(errors.New("failed to insert CompositeKey{22, 22}"))
	}

	if engine.Cacher != nil {
		engine.Cacher.ClearBeans(engine.TableInfo(compositeKeyVal).Name)
	}

	cps = make([]CompositeKey, 0)
	err = engine.Find(&cps)
	if err != nil {
		t.Error(err)
	}
	if len(cps) != 2 {
		t.Error(errors.New("should has two record"))
	}
	if cps[0] != compositeKeyVal {
		t.Error(errors.New("should be equeal"))
	}

	compositeKeyVal = CompositeKey{UpdateStr: "test1"}
	cnt, err = engine.Id(core.PK{11, 22}).Update(&compositeKeyVal)
	if err != nil {
		t.Error(err)
	} else if cnt != 1 {
		t.Error(errors.New("can't update CompositeKey{11, 22}"))
	}

	cnt, err = engine.Id(core.PK{11, 22}).Delete(&CompositeKey{})
	if err != nil {
		t.Error(err)
	} else if cnt != 1 {
		t.Error(errors.New("can't delete CompositeKey{11, 22}"))
	}
}

type User struct {
	UserId   string `xorm:"varchar(19) not null pk"`
	NickName string `xorm:"varchar(19) not null"`
	GameId   uint32 `xorm:"integer pk"`
	Score    int32  `xorm:"integer"`
}

func testCompositeKey2(engine *xorm.Engine, t *testing.T) {
	err := engine.DropTables(&User{})

	if err != nil {
		t.Error(err)
		panic(err)
	}

	err = engine.CreateTables(&User{})
	if err != nil {
		t.Error(err)
		panic(err)
	}

	cnt, err := engine.Insert(&User{"11", "nick", 22, 5})
	if err != nil {
		t.Error(err)
	} else if cnt != 1 {
		t.Error(errors.New("failed to insert User{11, 22}"))
	}

	cnt, err = engine.Insert(&User{"11", "nick", 22, 6})
	if err == nil || cnt == 1 {
		t.Error(errors.New("inserted User{11, 22}"))
	}

	var user User
	has, err := engine.Id(core.PK{"11", 22}).Get(&user)
	if err != nil {
		t.Error(err)
	} else if !has {
		t.Error(errors.New("can't get User{11, 22}"))
	}

	// test passing PK ptr, this test seem failed withCache
	has, err = engine.Id(&core.PK{"11", 22}).Get(&user)
	if err != nil {
		t.Error(err)
	} else if !has {
		t.Error(errors.New("can't get User{11, 22}"))
	}

	user = User{NickName: "test1"}
	cnt, err = engine.Id(core.PK{"11", 22}).Update(&user)
	if err != nil {
		t.Error(err)
	} else if cnt != 1 {
		t.Error(errors.New("can't update User{11, 22}"))
	}

	cnt, err = engine.Id(core.PK{"11", 22}).Delete(&User{})
	if err != nil {
		t.Error(err)
	} else if cnt != 1 {
		t.Error(errors.New("can't delete CompositeKey{11, 22}"))
	}
}
