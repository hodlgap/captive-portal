{{- $alias := .Aliases.Table .Table.Name -}}

{{- /* PK 생략을 위한 slice 실행 */ -}}
{{ $columnsWithoutPK := slice .Table.Columns 1 }}

func chunk{{$alias.UpPlural}}(o []*{{$alias.UpSingular}}, chuckSize int) [][]*{{$alias.UpSingular}} {
	chunks := make([][]*{{$alias.UpSingular}}, 0, (len(o)+chuckSize-1)/chuckSize)

	for chuckSize < len(o) {
		o, chunks = o[chuckSize:], append(chunks, o[0:chuckSize:chuckSize])
	}
	chunks = append(chunks, o)

	return chunks
}

// BulkInsert{{$alias.UpSingular}} Ref) https://stackoverflow.com/a/25192138/8979550
func BulkInsert{{$alias.UpSingular}}(unsavedRows []*{{$alias.UpSingular}}, exec boil.ContextExecutor) error {
	if len(unsavedRows) == 0 {
		return nil
	}

	for _, {{$alias.DownPlural}} := range chunk{{$alias.UpPlural}}(unsavedRows, {{- div 65535 ($columnsWithoutPK | len) -}}) {
		valueStrings := make([]string, 0, len(unsavedRows))
		valueArgs := make([]interface{}, 0, len(unsavedRows)*{{- $columnsWithoutPK | len -}})

		for _, {{$alias.DownSingular}} := range {{$alias.DownPlural}} {
            valueStrings = append(valueStrings, "(
		    {{- range $index, $column := $columnsWithoutPK -}}
			{{- if $index -}},{{- end -}}?
            {{- end -}}
			)")
		    {{ range $column := $columnsWithoutPK }}
		    {{ $colAlias := $alias.Column $column.Name -}}
			valueArgs = append(valueArgs, {{$alias.DownSingular}}.{{ $colAlias }})
            {{- end -}}
		}
		stmt := fmt.Sprintf(
			"INSERT INTO "+
				TableNames.{{ titleCase .Table.Name }}+
				"("+
                {{- range $index, $column := $columnsWithoutPK -}}
                {{- $colAlias := $alias.Column $column.Name -}}
                {{ if $index }} ", " + {{ end }}
                {{ $alias.UpSingular }}Columns.{{ $colAlias }} +

                {{- end }}
				") VALUES %s",
			strings.Join(valueStrings, ","))
		if _, err := exec.Exec(stmt, valueArgs...); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
