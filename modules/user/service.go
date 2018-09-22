package user

var (
	// Collection default name.
	Collection = "deployments"
)

// Service contains all biz logic functioneleties to reduce coupling from http layer
type Service struct {
	// DB               *db.DB                    `inject:""`
}

/* func (r *Service) FindOne(query *bson.M) (*helpers.Model, error) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(Collection)
	dep := helpers.Model{}
	err := collection.Find(query).One(&dep)
	return &dep, err
} */
