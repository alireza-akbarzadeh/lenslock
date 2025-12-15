package views

import (
	"html/template"
	"io"

	"github.com/techhubies/lenslocked/internal/models"
)

var userListTmpl = template.Must(template.New("userlist").Parse(`
<h1>User List</h1>
<ul>
{{range .}}
	<li>{{.Name}} ({{.Email}})</li>
{{else}}
	<li>No users found.</li>
{{end}}
</ul>
`))

func RenderUserList(w io.Writer, users []*models.User) error {
	return userListTmpl.Execute(w, users)
}
