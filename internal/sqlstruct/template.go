package sqlstruct

import (
	"fmt"
	"os"
	"text/template"

	"test/internal/word"
)
//模板渲染 基本结构由一个go结构体 和其所属的tablename方法组成
/*
生成结果大致如下
type 大写驼峰的表的名称struct {
	//注释
	字段名 字段类型
	...	...
}
func (model 大写驼峰的表名称) TableName() string {
	return "表名称"
}

 */
const strcutTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	strcutTpl string
}
//用来存储转换后的go结构体中的所有字段信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}
//用来存储最终用于渲染的模板对象信息
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}
// 通过查询COLUMNS表所组装得到的tbColums进一步转换
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}
//对模板自定义函数和模块对象进行处理
func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	//声明sqlstruct模板对象  定义tocamelcase函数与word.UndersoreToUpperCamelCase 做绑定 等同于我们在模板中使用ToCamelCase函数就等于使用word.UndersoreToUpperCamelCase
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UndersoreToUpperCamelCase,
	}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	//最后调用excute进行渲染
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}

