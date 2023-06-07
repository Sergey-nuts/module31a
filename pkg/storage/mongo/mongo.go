package mongo

import (
	"GoNews/pkg/storage"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const (
// 	dbName         = "GoNews"
// 	collectionName = "posts"
// )

type Mongodb struct {
	db         *mongo.Client
	dbName     string
	collection string
}

func New(conf, dbName, coll string) (*Mongodb, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return &Mongodb{client, dbName, coll}, nil
}

// получение всех публикаций
func (m *Mongodb) Posts() ([]storage.Post, error) {
	coll := m.db.Database(m.dbName).Collection(m.collection)
	filter := bson.D{}
	cur, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var posts []storage.Post
	for cur.Next(context.Background()) {
		var p storage.Post
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, cur.Err()
}

// создание новой публикации
func (m *Mongodb) AddPost(p storage.Post) error {
	coll := m.db.Database(m.dbName).Collection(m.collection)
	_, err := coll.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	return nil
}

// обновление публикации
func (m *Mongodb) UpdatePost(p storage.Post) error {
	coll := m.db.Database(m.dbName).Collection(m.collection)
	filter := bson.D{{"id", p.ID}}
	update := bson.D{{"$set", bson.D{
		{"title", p.Title},
		{"content", p.Content},
		{"author_id", p.AuthorID},
		{"authorname", p.AuthorName},
		{"createdat", p.CreatedAt},
	}}}

	_, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// удаление публикации по ID
func (m *Mongodb) DeletePost(p storage.Post) error {
	coll := m.db.Database(m.dbName).Collection(m.collection)
	filter := bson.D{{"id", p.ID}}
	_, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
