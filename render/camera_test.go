package render

import (
	"image/color"
	"testing"

	"github.com/Shamanskiy/go-ray-tracer/core"
	"github.com/Shamanskiy/go-ray-tracer/utils"
)

func TestCamera_Default(t *testing.T) {
	t.Log("Default camera without randomness")
	settings := DefaultCameraSettings()
	camera := NewCamera(&settings)
	core.Random().Disable()
	defer core.Random().Enable()

	utils.CheckResult(t, "Camera origin", camera.origin, settings.LookFrom)

	t.Log("  GetRay u=0.5 v=0.5")
	ray := camera.GetRay(0.5, 0.5)
	core.CheckVec3Tol(t, "Ray origin", ray.Origin, settings.LookFrom)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), settings.LookAt)

	t.Log("  GetRay u=0.5 v=0.0")
	ray = camera.GetRay(0.5, 0.0)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), core.Vec3{0., 1., -1})

	t.Log("  GetRay u=0.5 v=1.0")
	ray = camera.GetRay(0.5, 1.0)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), core.Vec3{0., -1., -1})

	t.Log("  GetRay u=0.0 v=0.0")
	ray = camera.GetRay(0.0, 0.0)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), core.Vec3{-2., 1., -1})

	t.Log("  GetRay u=1.0 v=1.0")
	ray = camera.GetRay(1.0, 1.0)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), core.Vec3{2., -1., -1})
}

func TestCamera_Custom(t *testing.T) {
	t.Log("Camera with custom settings without randomness")
	settings := DefaultCameraSettings()
	settings.LookAt = core.Vec3{3., 0., 4}
	settings.AspectRatio = 1.0
	camera := NewCamera(&settings)
	core.Random().Disable()
	defer core.Random().Enable()

	core.CheckVec3Tol(t, "Camera origin", camera.origin, settings.LookFrom)

	t.Log("  GetRay u=0.5 v=0.5")
	ray := camera.GetRay(0.5, 0.5)
	core.CheckVec3Tol(t, "Ray origin", ray.Origin, settings.LookFrom)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), settings.LookAt)

	t.Log("  GetRay u=0.5 v=0.0")
	ray = camera.GetRay(0.5, 0.0)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), core.Vec3{3., 5., 4.})

	t.Log("  GetRay u=0.0 v=0.5")
	ray = camera.GetRay(0.0, 0.5)
	core.CheckVec3Tol(t, "Ray target", ray.Eval(1.0), core.Vec3{7., 0., 1.})
}

func TestCamera_indexToU(t *testing.T) {
	t.Log("Camera with 100 px height and 1:1 aspect ratio")
	settings := DefaultCameraSettings()
	settings.ImagePixelHeight = 100
	settings.AspectRatio = 2.0
	camera := NewCamera(&settings)

	core.Random().Disable()
	defer core.Random().Enable()

	utils.CheckResult(t, "Image height", camera.pixelHeight, 100)
	utils.CheckResult(t, "Image width", camera.pixelWidth, 200)

	t.Log("  Pixel 0 to u")
	utils.CheckResult(t, "u param", camera.indexToU(0), core.Real(0.))
	t.Log("  Pixel 100 to u")
	utils.CheckResult(t, "u param", camera.indexToU(100), core.Real(0.5))
	t.Log("  Pixel 200 to u")
	utils.CheckResult(t, "u param", camera.indexToU(200), core.Real(1.))

	t.Log("  Pixel 0 to v")
	utils.CheckResult(t, "v param", camera.indexToV(0), core.Real(0.))
	t.Log("  Pixel 50 to v")
	utils.CheckResult(t, "v param", camera.indexToV(50), core.Real(0.5))
	t.Log("  Pixel 100 to v")
	utils.CheckResult(t, "v param", camera.indexToV(100), core.Real(1.))
}

func TestCamera_toRGBA(t *testing.T) {
	t.Log("Black to RGB")
	colorIn := core.Vec3{0., 0., 0.}
	colorOut := color.RGBA{0, 0, 0, 255}
	utils.CheckResult(t, "RGBA color", toRGBA(colorIn), colorOut)

	t.Log("White to RGB")
	colorIn = core.Vec3{1., 1., 1.}
	colorOut = color.RGBA{255, 255, 255, 255}
	utils.CheckResult(t, "RGBA color", toRGBA(colorIn), colorOut)

	t.Log("Gray to RGB with gamma correction")
	colorIn = core.Vec3{0.64, 0.64, 0.64}
	colorOut = color.RGBA{204, 204, 204, 255}
	utils.CheckResult(t, "RGBA color", toRGBA(colorIn), colorOut)
}

func TestCamera_RenderEmptyScene(t *testing.T) {
	t.Log("Given an empty scene with white background")
	scene := Scene{SkyColorTop: core.White, SkyColorBottom: core.White}

	imageSize := 2
	t.Logf("and a camera with %vx%v resolution,\n", imageSize, imageSize)
	settings := DefaultCameraSettings()
	settings.ImagePixelHeight = imageSize
	settings.AspectRatio = 1.0
	settings.Antialiasing = 1
	camera := NewCamera(&settings)

	t.Logf("  the rendered image should be a %vx%v white square:\n", imageSize, imageSize)
	renderedImage := camera.Render(&scene)

	utils.CheckResult(t, "Image width", renderedImage.Bounds().Size().X, imageSize)
	utils.CheckResult(t, "Image height", renderedImage.Bounds().Size().Y, imageSize)

	expectedColor := color.RGBA{255, 255, 255, 255}
	for x := 0; x < imageSize; x++ {
		for y := 0; y < imageSize; y++ {
			utils.CheckResult(t, "Pixel color", renderedImage.At(x, y), expectedColor)
		}
	}

}
