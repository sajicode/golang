# Go for Web Development

Large scale web applications need to be
. Scalable
. Modular
. Maintainable
. High-performant.

## URL Struct

`
type URL struct {
  Scheme string
  Opaque string
  User *Userinfo
  Host string
  Path string
  RawQuery string
  Fragment string
}`

## Cookie Struct

`
type Cookie struct {
  Name string
  Value string
  Path string
  Domain string
  Expires time.Time
  RawExpires string
  MaxAge int
  Secure bool
  HttpOnly bool
  Raw string
  Unparsed []string
}`

. If the Expires field isn't set in a cookie struct, then the cookie is a session or temporary cookie.
. Session cookies are removed from the browser when the browser is closed.
. There are two ways of specifying the expiry time: the Expires field & the MaxAge field.
. Expires tells us exactly when the cookie will expire & MaxAge tells us how long the cookie should last (in seconds), starting from the time it's created in the browser.
