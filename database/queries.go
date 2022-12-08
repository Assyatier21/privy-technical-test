package database

const (
	GetListOfCakes       = "SELECT * FROM privy_cakes ORDER BY rating DESC, title ASC LIMIT %d OFFSET %d"
	GetDetailsOfCakeByID = "SELECT * FROM privy_cakes WHERE id='%d'"
	InsertCake           = "INSERT INTO privy_cakes VALUES('%d','%s','%s','%f','%s','%s','%s')"
	UpdateCakeByID       = "UPDATE privy_cakes SET title='%s',description='%s',rating='%f',image='%s',created_at='%s',updated_at='%s' WHERE id='%d'"
	DeleteCakeByID       = "DELETE FROM privy_cakes WHERE id='%d'"
)
