// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package http

import (
	"html/template"
	"path"

	"github.com/anthony19114/commerciumx/swarm/api"
)

type htmlListData struct {
	URI  *api.URI
	List *api.ManifestList
}

var htmlListTemplate = template.Must(template.New("html-list").Funcs(template.FuncMap{"basename": path.Base}).Parse(`
<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="shortcut icon" type="image/x-icon" href="data:image/png;base64,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"/>
	<title>Swarm index of {{ .URI }}</title>
</head>

<body>
  <h1>Swarm index of {{ .URI }}</h1>
  <hr>
  <table>
    <thead>
      <tr>
	<th>Path</th>
	<th>Type</th>
	<th>Size</th>
      </tr>
    </thead>

    <tbody>
      {{ range .List.CommonPrefixes }}
	<tr>
	  <td><a href="{{ basename . }}/">{{ basename . }}/</a></td>
	  <td>DIR</td>
	  <td>-</td>
	</tr>
      {{ end }}

      {{ range .List.Entries }}
	<tr>
	  <td><a href="{{ basename .Path }}">{{ basename .Path }}</a></td>
	  <td>{{ .ContentType }}</td>
	  <td>{{ .Size }}</td>
	</tr>
      {{ end }}
  </table>
  <hr>
</body>
`[1:]))
