module github.com/aduryagin/cfop-backend

go 1.12

require (
	github.com/99designs/gqlgen v0.7.2
	github.com/agnivade/levenshtein v1.0.1
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/golang-lru v0.5.1
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a
	github.com/lib/pq v1.0.0
	github.com/pkg/errors v0.8.1
	github.com/rs/cors v1.6.0
	github.com/urfave/cli v1.20.0
	github.com/vektah/gqlparser v1.1.2
	golang.org/x/tools v0.0.0-20190321232350-e250d351ecad
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/aduryagin/cfop-backend/db v0.0.0 => ./db

replace github.com/aduryagin/cfop-backend v0.0.0 => ./
