package dataorm

import (
	"errors"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

var config = "****"


//数据库三个表中对象的实例化
//orm:"pk"代表主键

type User struct {
	ID int `orm:"pk;auto"`
	Name string `orm:"size(100)"`
	Firstname string `orm:"size(50)"`
	Lastname string `orm:"size(50)"`
	Phone string `orm:"size(50)"`
	Email string `orm:"size(50)"`
	Password string `orm:"size(50)"`
	Status string `orm:"default('')"`
}

type Room struct {
	ID int `orm:"pk;auto"`
	Name string `orm:"size(50)"`
	Creator string `orm:"size(100)"`
	Data string `orm:"size(50)"`
}

type Userinroom struct {
	ID int `orm:"pk;auto"`
	Userid int
	Roomid int
}


//init function
func init(){

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// register model
	orm.RegisterModel(new(User),new(Room),new(Userinroom))

	// set default database
	orm.RegisterDataBase("default","mysql",config,30)

	//新建表
	//orm.RunSyncdb("default", true, true)
}


//@function: 将输入对象插入到合适的数据库中
//@param: data:插入对象，主键缺省，数据库自动填充
//@return：如果插入失败，返回错误信息

func insert(data interface{}) error {
	orm.Debug=true
	o:=orm.NewOrm()
	o.Using("default")

	if value,ok:=data.(User);ok{
		 if _,err:=o.Insert(&value);err!=nil{
		 	return err
		 }
		 return nil
	}
	if value,ok:=data.(Room);ok{
		if _,err:=o.Insert(&value);err!=nil{
			return err
		}
		return nil
	}
	if value,ok:=data.(Userinroom);ok{
		if _,err:=o.Insert(&value);err!=nil{
			return err
		}
		return nil
	}
	err:=errors.New("Wrong type!")
	return err
}

//@function: 根据需求查询对象
//@param1: tablename:表名<User/Room/Userinroom>
//@param2: cols : 需要显示的列名，输入nil则代表*
//@param3: where : 匹配项，输入nil则代表无
//@param4: values : 匹配值，输入nil则代表无
//@return：如果查询失败，返回错误信息
//@example: tablename=User;cols=nil;where=["name"];values=["jack"] ==> select * from User where name = "jack"

func query(tablename string,cols []string,where []string, values []string)(interface{},error){

	orm.Debug = true;
	o := orm.NewOrm()
	o.Using("default")
	if tablename=="User"{
		var users [] User
		SQL:="SELECT "
		if len(cols)!=0{
			for i:=0;i<len(cols);i++{
				SQL=SQL+cols[i]+" ,"
			}
		}else {
			SQL=SQL+" * ,"
		}
		SQL=SQL[:len(SQL)-1]+" FROM user"
		if len(where)!=0{
			SQL+=" WHERE "
			for i:=0;i<len(cols);i++{
				if where[i]=="i_d"{
					SQL=SQL+where[i]+" = "+values[i]+" AND "
				}else {
					SQL=SQL+where[i]+" = "+"\""+values[i]+"\""+" AND "
				}

			}
			SQL=SQL[:len(SQL)-4]
		}
		_,err := o.Raw(SQL).QueryRows(&users)
		if err == nil {
			return users,err
		}
	}
	if tablename=="Room"{
		var rooms [] Room
		SQL:="SELECT "
		if len(cols)!=0{
			for i:=0;i<len(cols);i++{
				SQL=SQL+cols[i]+" ,"
			}
		}else {
			SQL=SQL+" * ,"
		}
		SQL=SQL[:len(SQL)-1]+" FROM room"
		if len(where)!=0{
			SQL+=" WHERE "
			for i:=0;i<len(cols);i++{
				if where[i]=="i_d"{
					SQL=SQL+where[i]+" = "+values[i]+" AND "
				}else {
					SQL=SQL+where[i]+" = "+"\""+values[i]+"\""+" AND "
				}

			}
			SQL=SQL[:len(SQL)-4]
		}
		_,err := o.Raw(SQL).QueryRows(&rooms)
		if err == nil {
			return rooms,err
		}
	}
	if tablename=="Userinroom"{
		var userinrooms [] Userinroom
		SQL:="SELECT "
		if len(cols)!=0{
			for i:=0;i<len(cols);i++{
				SQL=SQL+cols[i]+" ,"
			}
		}else {
			SQL=SQL+" * ,"
		}
		SQL=SQL[:len(SQL)-1]+" FROM userinroom"
		if len(where)!=0{
			SQL+=" WHERE "
			for i:=0;i<len(cols);i++{
				if where[i]=="i_d"{
					SQL=SQL+where[i]+" = "+values[i]+" AND "
				}else {
					SQL=SQL+where[i]+" = "+"\""+values[i]+"\""+" AND "
				}

			}
			SQL=SQL[:len(SQL)-4]
		}
		_,err := o.Raw(SQL).QueryRows(&userinrooms)
		if err == nil {
			return userinrooms,err
		}
	}
	return nil,errors.New("Wrong table name")

}

//@function: 删除目标项
//@param1: tablename:表名<User/Room/Userinroom>
//@param2: id : 待删除行的主键
//@return：如果删除失败，返回错误信息

func delete(tablename string,id int) error{
	orm.Debug = true;
	o := orm.NewOrm()
	o.Using("default")
	if tablename=="User"{
		user := User{ID: id}
		n, err := o.Delete(&user)
		if n > 0 && err == nil {
			return nil
		} else {
			return err
		}
	}
	if tablename=="Room"{
		room := Room{ID: id}
		n, err := o.Delete(&room)
		if n > 0 && err == nil {
			return nil
		} else {
			return err
		}
	}
	if tablename=="Userinroom"{
		uir := Userinroom{ID: id}
		n, err := o.Delete(&uir)
		if n > 0 && err == nil {
			return nil
		} else {
			return err
		}
	}
	return errors.New("Wrong table name")
}

//@function: 将输入对象替换数据表中的项(主键一样)
//@param: data:替换对象
//@return：如果替换失败，返回错误信息
//warning：是将输入的data替换整个原来的数据项，并不是修改data中不一样的项
func update(data interface{}) error{
	orm.Debug=true
	o:=orm.NewOrm()
	o.Using("default")

	if value,ok:=data.(User);ok{
		if _,err:=o.Update(&value);err!=nil{
			return err
		}
		return nil
	}
	if value,ok:=data.(Room);ok{
		if _,err:=o.Update(&value);err!=nil{
			return err
		}
		return nil
	}
	if value,ok:=data.(Userinroom);ok{
		if _,err:=o.Update(&value);err!=nil{
			return err
		}
		return nil
	}
	err:=errors.New("Wrong type!")
	return err
}
