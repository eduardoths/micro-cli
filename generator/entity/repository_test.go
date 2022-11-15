package entity_test

import (
	"testing"

	"github.com/eduardoths/micro-cli/generator/entity"
	"github.com/eduardoths/micro-cli/generator/file"
	"github.com/eduardoths/micro-cli/tests/utils"
)

func TestNewRepository(t *testing.T) {
	t.Run("it should return valid interfaces", func(t *testing.T) {
		repo := entity.NewRepository(
			entity.NewEntityName("XptoStructName", "src/structs", "github.com/eduardoths/microservice"),
			"github.com/eduardoths/microservice",
		)

		actual := repo.Interface
		want := "\ntype XptoStructNameRepository interface {\n" +
			"\tGetAll(ctx context.Context) (xptoStructName []structs.XptoStructName, err error)\n" +
			"\tGet(ctx context.Context, id uuid.UUID) (xptoStructName structs.XptoStructName, err error)\n" +
			"}\n"
		if want != actual.String() {
			utils.Error(t, want, actual)
		}
	})

	t.Run("it should return valid imports", func(t *testing.T) {
		repo := entity.NewRepository(
			entity.NewEntityName("XptoStructName", "src/structs", "github.com/eduardoths/microservice"),
			"github.com/eduardoths/microservice",
		)

		actual := repo.Imports
		want := file.Imports{
			{Path: "github.com/google/uuid"},
			{Path: "context"},
			{Path: "github.com/eduardoths/microservice/src/structs"},
		}
		if want.String() != actual.String() {
			utils.Error(t, want, actual)
		}
	})

	t.Run("it should return valid file", func(t *testing.T) {
		repo := entity.NewRepository(
			entity.NewEntityName("XptoStructName", "src/structs", "github.com/eduardoths/microservice"),
			"github.com/eduardoths/microservice",
		)

		actual := repo.File()
		want := "package xptostructname\n\n" +
			"import (\n" +
			"\t\"context\"\n" +
			"\t\"github.com/eduardoths/microservice/src/structs\"\n" +
			"\t\"github.com/google/uuid\"\n" +
			")\n\n" +
			"type XptoStructNameRepository struct {}\n\n" +
			"func (xsnr XptoStructNameRepository) GetAll(ctx context.Context) (xptoStructName []structs.XptoStructName, err error) {\n}\n\n" +
			"func (xsnr XptoStructNameRepository) Get(ctx context.Context, id uuid.UUID) (xptoStructName structs.XptoStructName, err error) {\n}\n"
		if want != actual.String() {
			utils.Error(t, want, actual)
		}
	})
}
