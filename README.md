
![Logo](https://db5pap001files.storage.live.com/y4mCcDLbZr-ovo6BZJ60x6l1Vmu4ok2hzmvwwINdKs61iXeTe7aqOHj26T4I95NB6b0s37NsVFGTDZ_XcmCrRF0cLZ5LlgkpCq3dIzzlxe_8LRMhtVTNjuJJTdemSYjq3d66OSlJO6UHvIHPRjovVfPchXyEV3YlUImjnKhUOffOlr-0Oirdz3NSjZS721G_1hK?width=100&height=100&cropmode=none "Header")

# hazzard.uk/blog

This is the backend API for the https://hazzard.uk/blog website.

It takes the Environment Variable `MongoURI` for database access and uses REST for the routes.

To run you have two options:



## Go Build

Firstly make sure the `MongoURI` Environment variable set to log into your database.
For example:

```
.zshrc

export MongoURI="your mongo string with username and password" 
```

Then build the application using Go Build.
`go build .` or `go build main.go` from the project directory.

Finally, when inserting posts, stick to this schema:
```
author: "Joe Bloggs",
content: "This is a new blog",
metaTitle: "new-blog",
published: "01/11/2022",
slug: "newblog",
title: "Joes New Blog",
updated: "01/11/2022",
summary: "new blog"
```


## Docker

For this, you just need docker and a mongodb instance running in the cloud.

First clone the repository and then cd into it.
```
git clone https://github.com/xavierhazzardadmin/blog.git
cd blog
```

Then build the docker container
```docker
sudo docker build . --file Dockerfile.multi --tag YOUR_TAG_NAME:RELEASE 
```

then run the container

```docker
sudo docker run YOUR_TAGE_NAME:RELEASE -p 8080:8080 -e MongoURI='your string'
```



## API Reference

#### Get one Post

```http
  GET /api/posts/{id}
```

| URL Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. The ID of the post. |

#### Get all posts of one athor

```http
  GET /api/posts
```

| Form Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `author`      | `string` | **Required**. Author of posts to fetch. |

#### Get the 10 most recent posts

```http
  GET /api/cache
```

#### Delete one post

```http
  DELETE /api/posts/{id}
```

| URL Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Author of post to delete. |

#### Insert a Post

```http
  POST /api/posts/{id}
```

| JSON Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `author`      | `string` | **Required**. Author of post. |
| `content`      | `string` | **Required**. Main content of post. |
| `metaTitle`      | `string` | **Required**. Title of post for embeds. |
| `published`      | `string` | **Required**. Date of post creation. |
| `slug`      | `string` | **Required**. URL slug of post. |
| `title`      | `string` | **Required**. Title of post. |
| `updated`      | `string` | **Required**. Date of post update. |
| `summary`      | `string` | **Required**. Short summary of post.  |


#### Update a Post

```http
  PUT /api/posts/{id}
```

| JSON Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `author`      | `string` | **Required**. Author of post. |
| `content`      | `string` | **Required**. Main content of post. |
| `metaTitle`      | `string` | **Required**. Title of post for embeds. |
| `published`      | `string` | **Required**. Date of post creation. |
| `slug`      | `string` | **Required**. URL slug of post. |
| `title`      | `string` | **Required**. Title of post. |
| `updated`      | `string` | **Required**. Date of post update. |
| `summary`      | `string` | **Required**. Short summary of post.  |
## License

[MIT](https://choosealicense.com/licenses/mit/)

