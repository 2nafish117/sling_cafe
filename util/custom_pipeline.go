package util

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomPipeline struct {
	Pipe mongo.Pipeline
}

func (cp *CustomPipeline) AddStage(stage bson.D) *CustomPipeline {
	cp.Pipe = append(cp.Pipe, stage)
	return cp
}

// All stages
// $addFields
// $bucket
// $bucketAuto
// $collStats
// $count
// $facet
// $geoNear
// $graphLookup
// $group
// $indexStats
// $limit
// $listSessions
// $lookup
// $match
// $merge
// $out
// $planCacheStats
// $project
// $redact
// $replaceRoot
// $replaceWith
// $sample
// $set
// $skip
// $sort
// $sortByCount
// $unset
// $unwind

/*
db.meals.aggregate([
	{ $match: { datetime: { $gt: ISODate("2020-01-01T00:00:00.000Z"), $lte: ISODate("2020-05-01T00:00:00.000Z") } } },
	{ $group: { _id: { uid: "$uid" }, total: { $sum: "$cost" } } },
	{ $lookup: { from: "users", localField: "_id.uid", foreignField: "uid", as: "user" } },
	{ $unwind: { path: "$user" } },
	{ $project: { _id: 0 } },
	{ $sort: { total: -1 } }
]).pretty()
*/

func (cp *CustomPipeline) Match(doc bson.D) *CustomPipeline {
	return cp.AddStage(bson.D{bson.E{Key: "$match", Value: doc}})
}

func (cp *CustomPipeline) Group(doc bson.D) *CustomPipeline {
	return cp.AddStage(bson.D{bson.E{Key: "$group", Value: doc}})
}

func (cp *CustomPipeline) Lookup(doc bson.D) *CustomPipeline {
	return cp.AddStage(bson.D{bson.E{Key: "$lookup", Value: doc}})
}

func (cp *CustomPipeline) Unwind(doc bson.D) *CustomPipeline {
	return cp.AddStage(bson.D{bson.E{Key: "$unwind", Value: doc}})
}

func (cp *CustomPipeline) Project(doc bson.D) *CustomPipeline {
	return cp.AddStage(bson.D{bson.E{Key: "$project", Value: doc}})
}

func (cp *CustomPipeline) Sort(doc bson.D) *CustomPipeline {
	return cp.AddStage(bson.D{bson.E{Key: "$sort", Value: doc}})
}
