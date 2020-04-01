// https://farenda.com/mongodb/how-to-generate-random-test-data-in-mongodb/

var day = 1000 * 60 * 60 * 24;
var randomDate = function () {
  return new Date(Date.now() - (Math.floor(Math.random() * day)));
}

// > db.food_transactions.findOne()
// {
//         "_id" : ObjectId("5e73b446304b884c445b55c1"),
//         "uid" : "2",
//         "datetime" : ISODate("2020-03-19T18:04:54.558Z"),
//         "type" : "breakfast",
//         "cost" : 30
// }

db.food_transactions.aggregate([
  { $match: { datetime: { $gte: ISODate("2020-03-01T00:00:00.000Z"), $lt: ISODate("2020-03-20T00:00:00.000Z") } } },
  { $group: { _id: {uid: "$uid", type: "$type"}, count: {$sum: 1}, total: {$sum: "$cost"} } },
  // { 
  //   $lookup: 
  //   {
  //     from: "users",
  //     localField: "uid",
  //     foreignField: "uid",
  //     as: "user"
  //   }
  // },
  // { $project: { uid: 0 } },
  // { $unwind: "$user" }
]).pretty()

db.food_transactions.aggregate([
  { $match: { datetime: { $gte: ISODate("2020-03-01T00:00:00.000Z"), $lt: ISODate("2020-03-20T00:00:00.000Z") } } },
  { $group: { _id: {uid: "$uid", type: "$type"}, count: {$sum: 1}, total: {$sum: "$cost"} } },
]).pretty()

db.food_transactions.aggregate([
  { $group: { _id: {uid: "$uid"}, 
  breakfast: {
    $switch: {
      branches: [
        {case: {type: "breakfast"}, then: {$sum: 1} }
      ]
    }
  } 
}}
]).pretty()