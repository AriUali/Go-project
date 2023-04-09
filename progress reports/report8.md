This week working on part with commenting items . For this Im add new structure

type Comment struct {
    Author string
    Text   string
}
 Comments []Comment
 
  And as Main function i use this and everything appear in mysql Database:
  
  func addComment(db *sql.DB, id string, comment Comment) error {
    _, err := db.Exec("INSERT INTO comments (item_id, author, text) VALUES (?, ?, ?)", id, comment.Author, comment.Text)
    if err != nil {
        return err
