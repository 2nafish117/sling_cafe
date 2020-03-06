// find bill of a particular user <empid>
// get all meals for <empid>
// for each meal get price of coresponding mealtype

db.meals.aggregate(
    { $group: { _id: "$empid" } },
    { $project: { "_id": 1 } }
)

db.meals.aggregate(
    { $group: { _id: "$empid" , mealtypeid: "$mealtypeid" } },
    { $project: { "_id": 1 } },
    {  }
)

db.meals.aggregate(
    { $group: { _id: { empid : "$empid" , mealtypeid: "$mealtypeid" } } },
    { $lookup: { from:  } }
)

db.meals.aggregate(
    { $group: { _id: { empid : "$empid" , mealtypeid: "$mealtypeid" } } }
)

// get all meals of that empid
empid = "INT4"
usermeals = db.meals.find({ empid : empid }, { _id : 0, mealtypeid : 1})
// for each mealtype find its cost 
// and keep adding them

// db.mealtypess.aggregate(
//     { $group : { mealtypesid : "$mealtypesid", due : { $sum : "" } } }
//     { total : { $sum : 1 } }
// )

// db.collection1.aggregate(
//     { $group : { name : "$name" , total : { $sum : "$amt" } } }
// )

// db.collection1.insertMany(
//     [
//         { _id: 1, cust_id: "abc1", ord_date: ISODate("2012-11-02T17:04:11.102Z"), status: "A", amount: 50 },
//         { _id: 2, cust_id: "xyz1", ord_date: ISODate("2013-10-01T17:04:11.102Z"), status: "A", amount: 100 },
//         { _id: 3, cust_id: "xyz1", ord_date: ISODate("2013-10-12T17:04:11.102Z"), status: "D", amount: 25 },
//         { _id: 4, cust_id: "xyz1", ord_date: ISODate("2013-10-11T17:04:11.102Z"), status: "D", amount: 125 },
//         { _id: 5, cust_id: "abc1", ord_date: ISODate("2013-11-12T17:04:11.102Z"), status: "A", amount: 25 }
//     ]
// )

// db.collection1.aggregate(
//     [
//         { $match: { status: "A" } },
//         { $group: { _id: "$cust_id", total: { $sum: "$amount" } } },
//         { $sort: { total: -1 } }
//     ]
// )

// db.collection1.aggregate(
//     [
//         { $group: { _id: "$cust_id", total: { $sum: "$amount" } } }
//     ]
// )


// db.collection1.aggregate(
//     [
//         { $group: { _id: "$cust_id", total: {  } } }
//     ]
// )

// works !!!!!!!!!!!
var meals_for_user = db.meals.find({ empid : "INT4" });
var types = meals_for_user.projection({_id: 0, mealtypeid: 1});

types.forEach()

var meal_types = db.mealtyes.find(types.toArray()[0])
meal_types_eaten.forEach({due: {$sum: db.mealtypes.find({})}})
var cost = db.mealtypes.find(mealtype.toArray()[0], { _id : 0, costtouser : 1 } );

db.meals.aggregate(
    { $group : { empids : "$empid", due: { $sum: 1 } } }
)


