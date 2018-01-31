package tmpl

var Pojo = `
package {{.SqlData.PackageName}}.domain;

import java.util.Date;

/**
* @table  {{.SqlData.ClassNameLower}}
*/
public class {{.SqlData.ClassName}}  {
private static final long serialVersionUID = 1L;
    {{range $index, $column := .SqlData.ColumnList}}
    private {{$column.Type}} {{$column.Name}}; {{if $column.CnName}} // $column.CnName {{end}}
    {{end}}
    {{range $index, $column := .SqlData.ColumnList}}
    public {{$column.Type}} get{{$column.NameUpper}}() {
        return this.{{$column.Name}};
    }

    public void set{{$column.NameUpper}}({{$column.Type}} {{$column.Name}}) {
        this.{{$column.Name}} = {{$column.Name}};
    }{{end}}

}
`
