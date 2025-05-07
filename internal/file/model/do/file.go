// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TFile is the golang structure of table t_file for DAO operations like Where/Data.
type TFile struct {
	g.Meta      		`orm:"table:t_file, do:true"`
	Id          		interface{} // ID
	MRoomNumber  	    interface{} // ID
	FileName       		interface{} //
	SaveName        	interface{} //
	FileSize    		interface{} //
	FileType       		interface{} //
	StoragePath      	interface{} //
	Status     			interface{} //
	UploaderId    		interface{} //
	UploaderName   		interface{} // ID
	UploadTime 			*gtime.Time //
	DeletorId 			interface{}
	DeletorName 		interface{}
	DeleteTime   		*gtime.Time //
}