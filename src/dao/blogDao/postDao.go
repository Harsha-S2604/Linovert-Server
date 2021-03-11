package blogDao

import (
	"context"
	"db"
	"errors"
	"log"
	"model/blog"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertPost(post *blog.Posts) string {
	dbCol := db.Database.Collection("blog")
	p, e := GetPosts()
	if e != nil {
		log.Println("Sorry something went wrong")
		return "Sorry something went wrong"
	}
	var id int64 = int64(len(p) - 1)
	var title string = post.GetTitle()
	var description string = post.GetDescription()
	var contents []string = post.GetContents()
	var codes []string = post.GetCodes()
	var createdOn string = post.GetCreatedOn()
	var posts blog.Posts
	posts.SetCodes(codes)
	posts.SetContents(contents)
	posts.SetTitle(title)
	posts.SetId(id)
	posts.SetDescription(description)
	posts.SetCreatedOn(createdOn)
	_, err := dbCol.InsertOne(context.TODO(), posts)
	if err != nil {
		log.Println("Failed to insert data")
		return "Sorry something went wrong failed to insert data"
	}
	log.Printf("Title %s successfully inserted\n", title)
	return "Successfully posted"
}

func GetPosts() ([]blog.Posts, error) {
	var res []blog.Posts
	dbCol := db.Database.Collection("blog")
	cursor, err := dbCol.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, errors.New("Sorry! Error in fetching the data :(")
	}
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, errors.New("Sorry! something went wrong, We couldn't assign the value :(")
	}
	return res, nil
}
