curl -X POST http://localhost:8080/posts \
  -H 'Content-Type: application/json' \
  -d '{"author": "Joe Bloggs", "content": "This is a new blog", "metaTitle": "new-blog","published": "01/11/2022", "slug": "newblog", "title": "Joe\'s New Blog", "updated": "01/11/2022", "summary": "new blog"}'\
