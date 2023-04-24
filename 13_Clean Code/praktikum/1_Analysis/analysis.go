package main

type user struct {
	id       int
	username string
	password string
}

// type userservice struct { <- tidak menggunakan camelCase pada nama struct
type userService struct {
	// t []user				<- tidak menggunakan nama yang jelas untuk nama variable di dalam struct
	users []user
}

// func (u userservice) getallusers() []user {    <- tidak menggunakan camelCase pada nama function serta mengganti nama struct
func (u userService) getAllUsers() []user {
	// return u.t 								 <- mengganti sesuai dengan nama variable yang ada di dalam struct
	return u.users
}

// func (u userService) getuserbyid(id int) user {  <- tidak menggunakan camelCase pada nama function serta mengganti nama struct
func (u userService) getUserById(id int) user {
	// for _, r := range u.t {     					<- tidak menggunakan nama yang jelas untuk nama variable yang ada dalam for loop,
	//												   serta mengganti sesuai dengan nama variable yang  ada di dalam struct
	for _, user := range u.users {
		// if id == r.id { 							<- mengganti sesuai dengan nama variable yang ada di dalam perulagan
		if id == user.id {
			// return r 							<- mengganti sesuai dengan nama variable yang ada di dalam perulagan
			return user
		}
	}
	return user{}
}
