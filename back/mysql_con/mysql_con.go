package mysql_con

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

var DB *sql.DB

func initDB() bool {
	//DigitalOcean: 206.189.68.176
	//Tencent：193.112.15.230
	//CloudCone: 173.82.105.39
	path := "root:Kkndcc110!@tcp(193.112.15.230:3306)/Stoneren?charset=utf8"
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		return false
	}
	return true
}

func Exec(SQL string) string {
	var ins string=""
	if initDB() == true {
		ret, _ := DB.Exec(SQL)
		if ret != nil {
			if LastInsertId, err := ret.LastInsertId(); nil == err {
				ins=strconv.FormatInt(LastInsertId, 10)
			}
		}
	}
	return ins
}

func Query(SQL string) ([]map[string]string, bool) {

	if initDB() != true { //连接数据库
		return nil, false
	}
	rows, err := DB.Query(SQL) //执行SQL语句，比如select * from users
	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns()            //获取列的信息
	count := len(columns)                   //列的数量
	var values = make([]interface{}, count) //创建一个与列的数量相当的空接口
	for i, _ := range values {
		var ii interface{} //为空接口分配内存
		values[i] = &ii    //取得指针
	}
	ret := make([]map[string]string, 0) //创建返回值：不定长的map类型切片
	for rows.Next() {
		err := rows.Scan(values...)  //开始读行，Scan函数只接受指针变量
		m := make(map[string]string) //1列的键值对
		if err != nil {
			panic(err)
		}
		for i, colName := range columns {
			var raw_value = *(values[i].(*interface{})) //读出raw数据，类型为byte
			b, _ := raw_value.([]byte)
			v := string(b) //将raw数据转换成字符串
			m[colName] = v //colName是键，v是值
		}
		ret = append(ret, m) //将单行所有列的键值对附加在总的返回值上（以行为单位）
	}
	defer rows.Close()
	if len(ret) != 0 {
		return ret, true
	}
	return nil, false
}
