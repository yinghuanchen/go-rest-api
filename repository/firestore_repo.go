package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"go-rest-api/entity"
	"google.golang.org/api/iterator"
	"log"
)

const (
	projectId = "golang-project"
	collectionName = "posts"
)

type repo struct{}

func NewFirestorePostRepository() PostRepository {
	return &repo{}
}

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err :=  firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("fail to create firestore client %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID": post.ID, 
		"Title": post.Title, 
		"Text": post.Text,
	})
	if err != nil {
		log.Fatalf("fail to add a new post %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err :=  firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("fail to create firestore client %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)
	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("fail to iterate through collection %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}