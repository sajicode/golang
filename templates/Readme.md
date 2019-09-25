# Templates

. If you're going to apply the same function set to numerous templates, you can use a Go function to create your templates and add your template functions each time.

`
func parseTemplateString(name, tpl string) *template.Template {
  t := template.New(name)
  t.Funcs(funcMap)
  t = template.Must(t.parse(tpl))
  return t
}
`

