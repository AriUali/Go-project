Working on refactoring code and testing the Rating and Commenting items functions , edit my code because of some errors 
happened when im trying to do it by log in in to the account.So im add function getItem because will be easier to select and commenting it:
func getItems(db *sql.DB) ([]Item, error) {
    rows, err := db.Query("SELECT i.id, i.name, i.description, i.price, c.id, c.item_id, c.content FROM items i LEFT JOIN comments c ON i.id = c.item_id")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

And now i can select any item i want to select.The next im creating an html pages for functions. Main part is Post method.
<form method="POST">
        <input type="hidden" name="item_id" value="{{ $item.ID }}">
        <textarea name="comment_content"></textarea><br>
        <input type="submit" value="Add Comment">
    </form>
    
And also editing other html pages in one style.
