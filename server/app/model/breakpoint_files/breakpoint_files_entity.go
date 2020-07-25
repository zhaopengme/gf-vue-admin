// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package breakpoint_files

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table breakpoint_files.
type Entity struct {
    Id         uint        `orm:"id,primary"  json:"id"`          // 自增ID    
    CreateAt   *gtime.Time `orm:"create_at"   json:"create_at"`   // 更新时间  
    UpdateAt   *gtime.Time `orm:"update_at"   json:"update_at"`   // 更新时间  
    DeleteAt   *gtime.Time `orm:"delete_at"   json:"delete_at"`   // 删除时间  
    FileName   string      `orm:"file_name"   json:"file_name"`   // 文件名    
    FileMd5    string      `orm:"file_md5"    json:"file_md_5"`   // 文件md5   
    FilePath   string      `orm:"file_path"   json:"file_path"`   // 文件路径  
    ChunkId    int         `orm:"chunk_id"    json:"chunk_id"`    // 关联id    
    ChunkTotal int         `orm:"chunk_total" json:"chunk_total"` // 切片总数  
    IsFinish   int         `orm:"is_finish"   json:"is_finish"`   // 是否完整  
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}