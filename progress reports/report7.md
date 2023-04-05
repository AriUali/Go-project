This week working on task Giving rating for items. And for this purpose im add new column “Rating” inside the structure and inside the database . This value taken 
from form inside html code.
 
  itemId := r.FormValue("id")
	rating := r.FormValue("rating")

Then we use range for visualize inside the localhost page all items and rating. The name of function is rateHandler().

for i, item := range items {
