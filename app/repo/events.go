package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"log"
	"sling_cafe/app/model"
	"sling_cafe/util"
	"time"
)

func MonthEndEvent() {
	for true {
		from := util.BeginningOfMonth(time.Now())
		to := util.EndOfMonth(time.Now())
		util.WaitUntil(context.Background(), to)
		fmt.Println("month passed", from.Format(time.RFC3339))
		CreateEmployeePosting(context.TODO(), from, to)
	}
}

func DayEndEvent() {
	for true {
		t := util.EndOfDay(time.Now())
		util.WaitUntil(context.Background(), t)
		fmt.Println("day passed", t.Format(time.RFC3339))
	}
}

func HourEndEvent() {
	for true {
		t := util.EndOfHour(time.Now())
		util.WaitUntil(context.Background(), t)
		fmt.Println("hour passed", t.Format(time.RFC3339))
	}
}

func MinuteEndEvent() {
	for true {
		from := util.BeginningOfMinute(time.Now())
		to := util.EndOfMinute(time.Now())
		util.WaitUntil(context.Background(), to)
		fmt.Println("minute passed", from.Format(time.RFC3339))
		CreateEmployeePosting(context.TODO(), from, to)
	}
}

func SecondEndEvent() {
	for true {
		t := util.EndOfSecond(time.Now())
		util.WaitUntil(context.Background(), t)
		fmt.Println("second passed", t.Format(time.RFC3339))
	}
}

/*
// employee transaction monthly posting
db.meal_entries.aggregate([
    { $match: { date_time: { $gte: ISODate("2020-01-01T00:00:00.000Z"), $lte: ISODate("2020-10-01T00:00:00.000Z") } } },
    { $group: { _id: { employee_id: "$employee_id" }, amount: { $sum: "$employee_cost" }} },
    { $project: { _id: 0, employee_id: '$_id.employee_id', transaction_id: "", amount: 1, date_time: new Date(), transaction_type: "MonthlyPosting" }},
])
*/

func CreateEmployeePosting(ctx context.Context, fromDate, toDate time.Time) error {
	mealEntriesAggregate := util.CustomPipeline{}
	mealEntriesAggregate.Match(
		bson.D{
			{
				Key: "date_time",
				Value: bson.D{
					{Key: "$gte", Value: fromDate},
					{Key: "$lte", Value: toDate},
				},
			},
		},
	).Group(
		bson.D{
			{Key: "_id", Value: bson.D{{Key: "employee_id", Value: "$employee_id"}}},
			{Key: "amount", Value: bson.D{{Key: "$sum", Value: "$employee_cost"}}},
		},
	).Project(
		bson.D{
			{Key: "_id", Value: 0}, {Key: "employee_id", Value: "$_id.employee_id"}, {Key: "amount", Value: 1},
			{Key: "transaction_id", Value: ""}, {Key: "date_time", Value: toDate}, {Key: "transaction_type", Value: model.MonthlyPosting},
		},
	)

	cursor, err := MealEntriesAggregate(ctx, mealEntriesAggregate.Pipe)

	if err != nil {
		log.Print(err.Error())
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		transaction := model.EmployeeTransaction{}

		if err := cursor.Decode(&transaction); err != nil {
			log.Print(err.Error())
			return err
		}
		if _, err := EmployeeTransactionsInsertOne(ctx, &transaction); err != nil {
			log.Print(err.Error())
			return err
		}
	}
	if err := cursor.Err(); err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
