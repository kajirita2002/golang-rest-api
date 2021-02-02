package repository

import (
	"context"
	"log"

	"github/kaji2002/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// repoという構造体を作る
type repo struct{}

// 新しいレポジトリを作成する
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

// 定数
const (
	projectId      string = "pragmatic-reviews-4eb12"
	collectionName string = "posts"
)

// repoのポインタ型のメソッド
func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	// contextを作成
	ctx := context.Background()
	// Firestore client を取得
	client, err := firestore.NewClient(ctx, projectId)
	// err処理
	if err != nil {
		log.Fatalf("Fail to create a Firestore Client: %v", err)
		return nil, err
	}
	// 処理が終わったらclientを閉じる
	defer client.Close()
	// collectionNameのコレクションにpostの値を追加する
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	// エラー処理
	if err != nil {
		log.Fatalf("Fail adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	// firestore client 作成
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Fail to create a Firestore Client: %v", err)
		return nil, err
	}
	//処理が終わったらclientを閉じる
	defer client.Close()
	//
	var posts []entity.Post
	//
	it := client.Collection(collectionName).Documents(ctx)
	for {
		// indexをインクリメントしたものを代入
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		// error処理
		if err != nil {
			log.Fatalf("Fail to iterate the list of posts: %v", err)
			return nil, err
		}
		//
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
