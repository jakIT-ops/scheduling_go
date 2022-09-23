package main

import (
	"github.com/go-co-op/gocron"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"scheduler/Jobs"
	"time"
)

//	s.Every(1).Seconds().Do(func() {
//		hello("Jawhaa")
//	})
//
//	// 5
//	s.StartBlocking()
//}
//
var jobs = new(Jobs.Jobs)

func Start(ctx *fiber.Ctx) error {
	//job := ctx.Params("jobName")
	//sch.Jobs()
	return nil
}

func Stop(sch gocron.Scheduler, ctx *fiber.Ctx) []*gocron.Job {
	//job := ctx.Params("jobName")
	//return nil
	return sch.Jobs()
}

func List(sch gocron.Scheduler, ctx *fiber.Ctx) ([]*gocron.Job, error) {
	return sch.Jobs(), nil
}

func main() {

	// Timer jobs
	//_, err := sch.Every(1).Weekday(time.Friday).At("10:46:00").Do(func() {
	//	jobs.hello()
	//})
	sch := gocron.NewScheduler(time.Local)
	// Job 0
	_, err := sch.Cron("0/10 * * * *").Do(func() {
		jobs.Hello()
	})
	if err != nil {
		log.Fatal("from job 0: ", err)
	}

	// Job 1
	_, err = sch.Cron("0/5 * * * *").Do(func() {
		jobs.Hello1()
	})
	if err != nil {
		log.Fatal("from job 1: ", err)
	}

	// Job 2
	_, err = sch.Cron("0/3 * * * *").Do(func() {
		jobs.Hello2()
	})
	if err != nil {
		log.Fatal("from job 2: ", err)
	}

	// Job 3
	_, err = sch.Every(1).Second().Do(func() {
		jobs.Hello()
	})
	if err != nil {
		log.Fatal("from job 3: ", err)
	}
	// bvh job uud async - aar ajillana
	sch.StartAsync()

	app := fiber.New()

	app.Get("/start/:jobName", Start)

	//app.Get("/stop/:jobName", Stop)

	app.Get("/list", func(ctx *fiber.Ctx) error {
		//i := []*gocron.Job
		jobs := sch.Jobs()
		ij := jobs[0]

		sch.Stop()
		//return sch.Jobs()
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"success": true,
			"jobs: ":  ij,
		})
	})
	//log.Fatal(http.ListenAndServe(":8080", nil))

	app.Listen(":8080")

}
