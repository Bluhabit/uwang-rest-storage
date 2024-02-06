package common

import "gorm.io/gen"

func GenerateEntity() {
	connection := GetDbConnection()
	if connection != nil {
		g := gen.NewGenerator(gen.Config{
			ModelPkgPath: "./entity",
			Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		})
		g.UseDB(GetDbConnection())
		g.GenerateAllTable()
		g.Execute()
	}
}
