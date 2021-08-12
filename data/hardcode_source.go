package data

import (
	"image/color"
	"image/png"
	"os"

	"github.com/mbaraa/ligma/errors"
	"github.com/mbaraa/ligma/models"
	"github.com/mbaraa/ligma/utils"
)

var (
	img0, _    = os.Open("./res/pics/society_service.png")
	img, _     = png.Decode(img0)
	ssRawImage = models.NewFormImage(img)

	_ = img0.Close()

	blueGoogle = color.RGBA64{R: 66, G: 133, B: 244, A: 255}

	ssFields = map[string]models.Field{
		"StudentName": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 960, Y: 435},
			ssRawImage,
		),
		"StudentId": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 447, Y: 435},
			ssRawImage, true,
		),
		"AcademicAdvisor": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 904, Y: 492},
			ssRawImage,
		),
		"Major": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 971, Y: 547},
			ssRawImage, true,
		),
		"Date": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 940, Y: 600},
			ssRawImage, true,
		),
		"Semester": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 435, Y: 600},
			ssRawImage, true,
		),
		"ActivityGoal": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 940, Y: 655},
			ssRawImage,
		),
		"TargetedPersonnel": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 803, Y: 1034},
			ssRawImage,
		),
		"ActivityTitle": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 946, Y: 1089},
			ssRawImage, true,
		),
		"DeservedPoints": models.NewTextField(
			utils.NewText("", blueGoogle, 20.5, "Default"),
			&models.Point2{X: 972, Y: 1143},
			ssRawImage, true,
		),
	}

	societyServiceForm = models.NewForm(
		"SocietyService",
		ssRawImage,
		ssFields,
	)
	forms = []*models.Form{
		societyServiceForm,
	}
)

type HardCodeSource struct{}

// ExistsByName reports whether the Form exists or not, and an occurring error
func (s *HardCodeSource) ExistsByName(string) (bool, error) {
	panic(errors.ErrNotImplemented)
}

// Get returns a form depending on its name, and an occurring error
func (s *HardCodeSource) Get(name string) (*models.Form, error) {
	for _, form := range forms {
		if form.GetName() == name {
			return form, nil
		}
	}
	return nil, errors.ErrNoFormFound
}

// GetAll returns all available forms, and an occurring error
func (s *HardCodeSource) GetAll() ([]*models.Form, error) {
	return forms, nil
}

// Count returns the number of available forms, and an occurring error
func (s *HardCodeSource) Count() (int64, error) {
	return int64(len(forms)), nil
}
