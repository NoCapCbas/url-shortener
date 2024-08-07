package urlshortener

type RedirectService interface {
  Find(code string) (*Redirect, error)
  Store(redirect *Redirect) error
}
