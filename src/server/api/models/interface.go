package models

type Resource interface {
  Find(string)
}

type Resources interface {
  All()
}
