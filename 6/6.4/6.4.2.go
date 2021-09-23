package mongodao

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/jaeyeom/gogo/task"
)

type MongoAccessor struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func New(path, db, c string) *MongoAccessor {
	session, err := mgo.Dial(path)
	if err != nil {
		return nil
	}
	collection := session.DB(db).C(c)
	return &MongoAccessor{
		session:    session,
		collection: collection,
	}
}

func (m *MongoAccessor) Close() error {
	m.session.Close()
	return nil
}

func idToObjectId(id task.ID) bson.ObjectId {
	return bson.ObjectIdHex(string(id))
}

func objectIdToID(objID bson.ObjectId) task.ID {
	return task.ID(objID.Hex())
}

func (m *MongoAccessor) Get(id task.ID) (task.Task, error) {
	t := task.Task{}
	err := m.collection.FindId(idToObjectId(id)).One(&t)
	return t, err
}

func (m *MongoAccessor) Put(id task.ID, t task.Task) error {
	return m.collection.UpdateId(idToObjectId(id), t)
}

func (m *MongoAccessor) Post(t task.Task) (task.ID, error) {
	objID := bson.NewObjectId()
	_, err := m.collection.UpsertId(objID, &t)
	return objectIdToID(objID), err
}

func (m *MongoAccessor) Delete(id task.ID) error {
	return m.collection.RemoveId(idToObjectId(id))
}
