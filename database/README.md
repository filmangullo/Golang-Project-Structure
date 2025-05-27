# `database`

### Contoh 1: Fungsi FindWhere (dinamis)

```sh
func (r *userRepository) FindWhere(condition map[string]interface{}) ([]models.User, error) {
	var users []models.User
	err := r.db.Where(condition).Find(&users).Error
	return users, err
}
```
####  Cara Pakai
```sh
users, err := repo.FindWhere(map[string]interface{}{
    "email": "example@mail.com",
    "status": "active",
})
```

### Contoh 2: Fungsi FindByField (spesifik satu field)

```sh
func (r *userRepository) FindByField(field string, value interface{}) ([]models.User, error) {
	var users []models.User
	err := r.db.Where(field+" = ?", value).Find(&users).Error
	return users, err
}
```
####  Cara Pakai
```sh
users, err := repo.FindByField("username", "captain123")
```