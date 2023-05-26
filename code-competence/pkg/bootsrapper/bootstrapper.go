package bootsrapper

import (
	"time"

	itemControllerPkg "github.com/DewaBiara/Code-Competence/internal/item/controller"
	itemRepositoryPkg "github.com/DewaBiara/Code-Competence/internal/item/repository/impl"
	itemServicePkg "github.com/DewaBiara/Code-Competence/internal/item/service/impl"
	userControllerPkg "github.com/DewaBiara/Code-Competence/internal/user/controller"
	userRepositoryPkg "github.com/DewaBiara/Code-Competence/internal/user/repository/impl"
	userServicePkg "github.com/DewaBiara/Code-Competence/internal/user/service/impl"
	"github.com/DewaBiara/Code-Competence/pkg/routes"
	jwtPkg "github.com/DewaBiara/Code-Competence/pkg/utils/jwt_service/impl"
	passwordPkg "github.com/DewaBiara/Code-Competence/pkg/utils/password/impl"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

func InitController(e *echo.Echo, db *gorm.DB, conf map[string]string) {
	passwordFunc := passwordPkg.NewPasswordFuncImpl()
	jwtService := jwtPkg.NewJWTService(conf["JWT_SECRET"], 1*time.Hour)

	// User
	userRepository := userRepositoryPkg.NewUserRepositoryImpl(db)
	userService := userServicePkg.NewUserServiceImpl(userRepository, passwordFunc, jwtService)
	userController := userControllerPkg.NewUserController(userService, jwtService)

	//Item
	itemRepository := itemRepositoryPkg.NewItemRepositoryImpl(db)
	itemService := itemServicePkg.NewItemServiceImpl(itemRepository)
	itemController := itemControllerPkg.NewItemController(itemService, jwtService)

	route := routes.NewRoutes(userController, itemController)
	route.Init(e, conf)
}
