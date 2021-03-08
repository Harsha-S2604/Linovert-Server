package blogService

import (
	"dao/blogDao"
	"log"
	"model/blog"
)

func InsertPost(post *blog.Posts) string {
	message := blogDao.InsertPost(post)
	return message
}

func GetPost() ([]blog.Posts, error) {
	val, err := blogDao.GetPosts()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for i, j := 0, len(val)-1; i < j; i, j = i+1, j-1 {
		val[i], val[j] = val[j], val[i]
	}

	if len(val) >= 5 {
		val = val[0:5]
	}

	return val, nil
}
