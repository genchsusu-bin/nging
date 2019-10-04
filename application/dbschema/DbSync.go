// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
	
	"time"
)

type Slice_DbSync []*DbSync

func (s Slice_DbSync) Range(fn func(m factory.Model) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_DbSync) RangeRaw(fn func(m *DbSync) error ) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// DbSync 数据表同步方案
type DbSync struct {
	base    factory.Base
	objects []*DbSync
	
	Id                    	uint    	`db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name                  	string  	`db:"name" bson:"name" comment:"方案名" json:"name" xml:"name"`
	SourceAccountId       	uint    	`db:"source_account_id" bson:"source_account_id" comment:"源数据库账号ID" json:"source_account_id" xml:"source_account_id"`
	DsnSource             	string  	`db:"dsn_source" bson:"dsn_source" comment:"同步源" json:"dsn_source" xml:"dsn_source"`
	DestinationAccountId  	uint    	`db:"destination_account_id" bson:"destination_account_id" comment:"目标数据库账号ID" json:"destination_account_id" xml:"destination_account_id"`
	DsnDestination        	string  	`db:"dsn_destination" bson:"dsn_destination" comment:"目标数据库" json:"dsn_destination" xml:"dsn_destination"`
	Tables                	string  	`db:"tables" bson:"tables" comment:"要同步的表" json:"tables" xml:"tables"`
	SkipTables            	string  	`db:"skip_tables" bson:"skip_tables" comment:"要跳过的表" json:"skip_tables" xml:"skip_tables"`
	AlterIgnore           	string  	`db:"alter_ignore" bson:"alter_ignore" comment:"要忽略的列、索引、外键" json:"alter_ignore" xml:"alter_ignore"`
	Drop                  	uint    	`db:"drop" bson:"drop" comment:"删除待同步数据库中多余的字段、索引、外键 " json:"drop" xml:"drop"`
	MailTo                	string  	`db:"mail_to" bson:"mail_to" comment:"发送邮件" json:"mail_to" xml:"mail_to"`
	Created               	uint    	`db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated               	int     	`db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

// - base function

func (this *DbSync) Trans() *factory.Transaction {
	return this.base.Trans()
}

func (this *DbSync) Use(trans *factory.Transaction) factory.Model {
	this.base.Use(trans)
	return this
}

func (this *DbSync) SetContext(ctx echo.Context) factory.Model {
	this.base.SetContext(ctx)
	return this
}

func (this *DbSync) Context() echo.Context {
	return this.base.Context()
}

func (this *DbSync) SetConnID(connID int) factory.Model {
	this.base.SetConnID(connID)
	return this
}

func (this *DbSync) SetNamer(namer func (string) string) factory.Model {
	this.base.SetNamer(namer)
	return this
}

func (this *DbSync) Namer() func(string) string {
	return this.base.Namer()
}

func (this *DbSync) SetParam(param *factory.Param) factory.Model {
	this.base.SetParam(param)
	return this
}

func (this *DbSync) Param() *factory.Param {
	if this.base.Param() == nil {
		return this.NewParam()
	}
	return this.base.Param()
}

// - current function

func (this *DbSync) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName,connID[0]).Use(this.base.Trans())
	}
	return factory.NewModel(structName,this.base.ConnID()).Use(this.base.Trans())
}

func (this *DbSync) Objects() []*DbSync {
	if this.objects == nil {
		return nil
	}
	return this.objects[:]
}

func (this *DbSync) NewObjects() factory.Ranger {
	return &Slice_DbSync{}
}

func (this *DbSync) InitObjects() *[]*DbSync {
	this.objects = []*DbSync{}
	return &this.objects
}

func (this *DbSync) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(this.base.ConnID()).SetTrans(this.base.Trans()).SetCollection(this.Name_()).SetModel(this)
}

func (this *DbSync) Short_() string {
	return "db_sync"
}

func (this *DbSync) Struct_() string {
	return "DbSync"
}

func (this *DbSync) Name_() string {
	if this.base.Namer() != nil {
		return WithPrefix(this.base.Namer()(this.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(this.Short_())(this))
}

func (this *DbSync) CPAFrom(source factory.Model) factory.Model {
	this.SetContext(source.Context())
	this.Use(source.Trans())
	this.SetNamer(source.Namer())
	return this
}

func (this *DbSync) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := this.base
	err := this.Param().SetArgs(args...).SetRecv(this).SetMiddleware(mw).One()
	this.base = base
	return err
}

func (this *DbSync) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *DbSync) GroupBy(keyField string, inputRows ...[]*DbSync) map[string][]*DbSync {
	var rows []*DbSync
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string][]*DbSync{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*DbSync{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (this *DbSync) KeyBy(keyField string, inputRows ...[]*DbSync) map[string]*DbSync {
	var rows []*DbSync
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]*DbSync{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (this *DbSync) AsKV(keyField string, valueField string, inputRows ...[]*DbSync) map[string]interface{} {
	var rows []*DbSync
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = this.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (this *DbSync) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = this.InitObjects()
	}
	return this.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (this *DbSync) Add() (pk interface{}, err error) {
	this.Created = uint(time.Now().Unix())
	this.Id = 0
	err = DBI.Fire("creating", this, nil)
	if err != nil {
		return
	}
	pk, err = this.Param().SetSend(this).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", this, nil)
	}
	return
}

func (this *DbSync) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	this.Updated = int(time.Now().Unix())
	if err = DBI.Fire("updating", this, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(this).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", this, mw, args...)
}

func (this *DbSync) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return this.Param().SetArgs(args...).SetMiddleware(mw)
}

func (this *DbSync) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return this.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (this *DbSync) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	kvset["updated"] = int(time.Now().Unix())
	m := *this
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = this.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (this *DbSync) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = this.Param().SetArgs(args...).SetSend(this).SetMiddleware(mw).Upsert(func() error { this.Updated = int(time.Now().Unix())
		return DBI.Fire("updating", this, mw, args...)
	}, func() error { this.Created = uint(time.Now().Unix())
	this.Id = 0
		return DBI.Fire("creating", this, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			this.Id = v
		} else if v, y := pk.(int64); y {
			this.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", this, mw, args...)
		} else {
			err = DBI.Fire("created", this, nil)
		}
	} 
	return 
}

func (this *DbSync) Delete(mw func(db.Result) db.Result, args ...interface{})  (err error) {
	
	if err = DBI.Fire("deleting", this, mw, args...); err != nil {
		return
	}
	if err = this.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", this, mw, args...)
}

func (this *DbSync) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return this.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (this *DbSync) Reset() *DbSync {
	this.Id = 0
	this.Name = ``
	this.SourceAccountId = 0
	this.DsnSource = ``
	this.DestinationAccountId = 0
	this.DsnDestination = ``
	this.Tables = ``
	this.SkipTables = ``
	this.AlterIgnore = ``
	this.Drop = 0
	this.MailTo = ``
	this.Created = 0
	this.Updated = 0
	return this
}

func (this *DbSync) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = this.Id
	r["Name"] = this.Name
	r["SourceAccountId"] = this.SourceAccountId
	r["DsnSource"] = this.DsnSource
	r["DestinationAccountId"] = this.DestinationAccountId
	r["DsnDestination"] = this.DsnDestination
	r["Tables"] = this.Tables
	r["SkipTables"] = this.SkipTables
	r["AlterIgnore"] = this.AlterIgnore
	r["Drop"] = this.Drop
	r["MailTo"] = this.MailTo
	r["Created"] = this.Created
	r["Updated"] = this.Updated
	return r
}

func (this *DbSync) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
			case "id": this.Id = param.AsUint(value)
			case "name": this.Name = param.AsString(value)
			case "source_account_id": this.SourceAccountId = param.AsUint(value)
			case "dsn_source": this.DsnSource = param.AsString(value)
			case "destination_account_id": this.DestinationAccountId = param.AsUint(value)
			case "dsn_destination": this.DsnDestination = param.AsString(value)
			case "tables": this.Tables = param.AsString(value)
			case "skip_tables": this.SkipTables = param.AsString(value)
			case "alter_ignore": this.AlterIgnore = param.AsString(value)
			case "drop": this.Drop = param.AsUint(value)
			case "mail_to": this.MailTo = param.AsString(value)
			case "created": this.Created = param.AsUint(value)
			case "updated": this.Updated = param.AsInt(value)
		}
	}
}

func (this *DbSync) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
		case map[string]interface{}:
			for kk, vv := range k {
				this.Set(kk, vv)
			}
		default:
			var (
				kk string
				vv interface{}
			)
			if k, y := key.(string); y {
				kk = k
			} else {
				kk = fmt.Sprint(key)
			}
			if len(value) > 0 {
				vv = value[0]
			}
			switch kk {
				case "Id": this.Id = param.AsUint(vv)
				case "Name": this.Name = param.AsString(vv)
				case "SourceAccountId": this.SourceAccountId = param.AsUint(vv)
				case "DsnSource": this.DsnSource = param.AsString(vv)
				case "DestinationAccountId": this.DestinationAccountId = param.AsUint(vv)
				case "DsnDestination": this.DsnDestination = param.AsString(vv)
				case "Tables": this.Tables = param.AsString(vv)
				case "SkipTables": this.SkipTables = param.AsString(vv)
				case "AlterIgnore": this.AlterIgnore = param.AsString(vv)
				case "Drop": this.Drop = param.AsUint(vv)
				case "MailTo": this.MailTo = param.AsString(vv)
				case "Created": this.Created = param.AsUint(vv)
				case "Updated": this.Updated = param.AsInt(vv)
			}
	}
}

func (this *DbSync) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = this.Id
	r["name"] = this.Name
	r["source_account_id"] = this.SourceAccountId
	r["dsn_source"] = this.DsnSource
	r["destination_account_id"] = this.DestinationAccountId
	r["dsn_destination"] = this.DsnDestination
	r["tables"] = this.Tables
	r["skip_tables"] = this.SkipTables
	r["alter_ignore"] = this.AlterIgnore
	r["drop"] = this.Drop
	r["mail_to"] = this.MailTo
	r["created"] = this.Created
	r["updated"] = this.Updated
	return r
}

func (this *DbSync) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = this.AsRow()
	}
	return factory.BatchValidate(this.Short_(), kvset)
}

func (this *DbSync) Validate(field string, value interface{}) error {
	return factory.Validate(this.Short_(), field, value)
}

