package templates

import (
	"fmt"
	"os"
	"sqltostruct/internal/model"
	"sqltostruct/libs/words"
	"text/template"

	"github.com/gogf/gf/v2/text/gstr"
)

/**
{{- 忽略前面的空
gt 大于
lt 小于
eq 等于
ne 不等于
*/
/*const structTpl = `
package {{.PackageName}}

type {{.DB.TableName | ToBigCamelCase }}Base struct{
{{- range .DB.Columns }}
{{- $isShowBaseLength := len .IsShowBase }}
{{- if eq $isShowBaseLength 0}}
{{$length := len .Comment }}
{{- if gt $length 0}}	//{{- .Comment }} {{- else}}	//{{- .Name}} {{end}}
{{- $typeLen := len .Type }}
{{if gt $typeLen 0}}	{{.Name | ToBigCamelCase}}	{{.Type}} {{.Tag}} {{else}}	{{.Name}}{{end}}
{{- end}}
{{- end}}
}

type {{.DB.TableName | ToBigCamelCase }} struct{
{{ range .DB.Columns }} {{$length := len .Comment }}{{if gt $length 0}}	//{{ .Comment }} {{else}} //{{.Name}} {{end}}{{$typeLen := len .Type }} {{if gt $typeLen 0}}
	{{.Name | ToBigCamelCase}}	{{.Type}} {{.Tag}} {{else}}{{.Name}}{{end}}
{{end}}
}

`*/

// 2022-11-14 备份
const structTpl = `
package {{.PackageName}}


type {{.DB.TableName | ToBigCamelCase }} struct{
{{ range .DB.Columns }} {{$typeLen := len .Type }}
	{{- if gt $typeLen 0}} {{ .Name | ToBigCamelCase}}	{{.Type}} {{ .Tag}} {{- else}}{{- .Name}}{{end}} {{- $length := len .Comment }} {{- if gt $length 0}} // {{ .Comment }} {{else}} // {{ .Name}} {{end}}
{{end}}
}
`

/*const structTpl = `
type {{.DB.TableName | ToBigCamelCase }}BaseTest struct{ {{- range .DB.Columns }} {{- $isShowBaseLength := len .IsShowBase }} {{- if eq $isShowBaseLength 0}} {{- $typeLen := len .Type }} {{if gt $typeLen 0}} {{.Name | ToBigCamelCase}} {{.Type}} {{.Tag}} {{else}} {{.Name}}{{end}} {{- $length := len .Comment }} {{if gt $length 0}} //{{- .Comment }} {{- else}} //{{- .Name}} {{end}} {{- end}} {{- end}} }
`*/

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name       string
	Type       string
	Tag        string
	Comment    string
	IsShowBase string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

type StructTemplateDBPackage struct {
	PackageName string
	DB          StructTemplateDB
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*model.TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		// `gorm:"create_user" form:"createUser" json:"createUser"`
		// tag := fmt.Sprintf("`"+"gorm:"+"\"%s\""+" form:"+"\"%s\""+" json:"+"\"%s\""+"`", column.ColumnName, words.ToCamelCase(column.ColumnName), words.ToCamelCase(column.ColumnName))
		tag := fmt.Sprintf("`"+"gorm:"+"\"%s\""+"`", column.ColumnName)
		if gstr.Contains(column.ColumnComment, "\r\n") {
			column.ColumnComment = gstr.ReplaceI(column.ColumnComment, "\r\n", " ")
		}
		tplColumns = append(tplColumns, &StructColumn{
			Type: model.DBTypeToStructType[column.DataType],
			Tag:  tag,
			// 过滤掉字段备注信息里的/r/n
			Comment:    column.ColumnComment,
			Name:       column.ColumnName,
			IsShowBase: model.DBSkipColName[column.ColumnName],
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(packageName string, tableName string, tplColumns []*StructColumn, filePwd string, fileName string) error {
	// 解析模板文件
	// t1, err :=template.ParseFiles("./test.tpl")
	// template.Must 检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写
	tpl := template.Must(template.New(packageName).Funcs(template.FuncMap{
		"ToBigCamelCase": words.ToBigCamelCase,
		"ToCamelCase":    words.ToBigCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDBPackage{
		PackageName: packageName,
		DB: StructTemplateDB{
			TableName: tableName,
			Columns:   tplColumns,
		},
	}

	println(fmt.Sprintf("%s%s.go", filePwd, fileName))

	// 创建文件句柄
	f, er := os.Create(fmt.Sprintf("%s%s.go", filePwd, fileName))
	if er != nil {
		panic(er)
	}

	/*
		func (t Template) ExecuteTemplate(wrio.Writer, namestring, data interface{})error
		渲染模板,第一个参数是io类型,可以选择写入文件或者是控制台os.Stdout
	*/
	err := tpl.Execute(f, tplDB)
	if err != nil {
		return err
	}

	return nil
}
