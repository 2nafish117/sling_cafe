// get report of company
// db.meal_entries.aggregate([
// 	{ $project: { employee_id: 1, date_time: 1, meal_id: 1, is_breakfast: {$cond: [{$eq: ['$type', "breakfast"]}, 1, 0]},cost_breakfast: {$cond: [{$eq: ['$type', "breakfast"]}, '$cost', 0]},is_lunch: {$cond: [{$eq: ['$type', "lunch"]}, 1, 0]},cost_lunch: {$cond: [{$eq: ['$type', "lunch"]}, '$cost', 0]},is_snack: {$cond: [{$eq: ['$type', "snack"]}, 1, 0]},cost_snack: {$cond: [{$eq: ['$type', "snack"]}, '$cost', 0]}}},
// 	{ $group: { _id: null, breakfast_quantity: { $sum: '$is_breakfast' }, breakfast_cost: { $sum: '$cost_breakfast' } ,lunch_quantity: { $sum: '$is_lunch' }, lunch_cost: { $sum: '$cost_lunch' }, snack_quantity: { $sum: '$is_snack' }, snack_cost: { $sum: '$cost_snack' }, total: { $sum: "$cost" }} },
// 	{ $project: { _id: 0, total: 1, breakfast: { quantity: '$breakfast_quantity', cost: '$breakfast_cost' }, lunch: { quantity: '$lunch_quantity', cost: '$lunch_cost' }, snack: { quantity: '$snack_quantity', cost: '$snack_cost' }, }},
// ]).pretty()

// employee wise report of meal_entries
db.meal_entries.aggregate([
	{ $project: { employee_id: 1, is_breakfast: {$cond: [{$eq: ['$meal_id', "breakfast"]}, 1, 0]}, is_lunch: {$cond: [{$eq: ['$meal_id', "lunch"]}, 1, 0]}, is_snack: {$cond: [{$eq: ['$meal_id', "snack"]}, 1, 0]}}},
	{ $group: { _id: { employee_id: "$employee_id" }, breakfast_quantity: { $sum: '$is_breakfast' }, lunch_quantity: { $sum: '$is_lunch' }, snack_quantity: { $sum: '$is_snack' } } },
	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
	{ $unwind: { path: "$employee" } },
	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', breakfast_quantity: '$breakfast_quantity', lunch_quantity: '$lunch_quantity', snack_quantity: '$snack_quantity' }},
]).pretty()

// employee wise report of meal_entries by a single employee id
db.meal_entries.aggregate([
	{ $match: { employee_id: "4" }},
	{ $project: { employee_id: 1, is_breakfast: {$cond: [{$eq: ['$meal_id', "breakfast"]}, 1, 0]}, is_lunch: {$cond: [{$eq: ['$meal_id', "lunch"]}, 1, 0]}, is_snack: {$cond: [{$eq: ['$meal_id', "snack"]}, 1, 0]}}},
	{ $group: { _id: { employee_id: "$employee_id" }, breakfast_quantity: { $sum: '$is_breakfast' }, lunch_quantity: { $sum: '$is_lunch' }, snack_quantity: { $sum: '$is_snack' } } },
	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
	{ $unwind: { path: "$employee" } },
	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', breakfast_quantity: '$breakfast_quantity', lunch_quantity: '$lunch_quantity', snack_quantity: '$snack_quantity' }},
]).pretty()

// employee transaction report for all employees
db.employee_transactions.aggregate([
	{ $project: { employee_id: 1, amount: {$cond: [{$eq: ['$transaction_type', "EmployeePayment"]}, {'$multiply':['$amount', NumberLong(-1)]}, '$amount']}}},
	{ $group: { _id: { employee_id: "$employee_id" }, amount: { $sum: '$amount' } } },
	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
	{ $unwind: { path: "$employee" } },
	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', amount: 1 }},
])

// employee transaction report for one employee based on employee_id
db.employee_transactions.aggregate([
	{ $match: { employee_id: "4" }},
	{ $project: { employee_id: 1, amount: {$cond: [{$eq: ['$transaction_type', "EmployeePayment"]}, {'$multiply':['$amount', NumberLong(-1)]}, '$amount']}}},
	{ $group: { _id: { employee_id: "$employee_id" }, amount: { $sum: '$amount' } } },
	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
	{ $unwind: { path: "$employee" } },
	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', amount: 1 }},
])

// employee transaction report for all employees
db.admin_transactions.aggregate([
	{ $project: { admin_id: 1, amount: {$cond: [{$eq: ['$transaction_type', "EmployeePayment"]}, {'$multiply':['$amount', NumberLong(-1)]}, '$amount']}}},
	{ $group: { _id: { employee_id: "$employee_id" }, amount: { $sum: '$amount' } } },
	{ $lookup: { from: "employees", localField: "_id.employee_id", foreignField: "employee_id", as: "employee" } },
	{ $unwind: { path: "$employee" } },
	{ $project: { _id: 0, name: '$employee.name', employee_id: '$employee.employee_id', amount: 1 }},
])

// employee transaction monthly posting
db.meal_transactions.aggregate([
    { $match: { date_time: { $gte: ISODate("2020-01-01T00:00:00.000Z"), $lte: ISODate("2020-10-01T00:00:00.000Z") } } },
    { $group: { _id: { employee_id: "$employee_id" }, amount: { $sum: "$cost" }} },
    { $project: { _id: 0, employee_id: '$_id.employee_id', transaction_id: "", amount: 1, date_time: new Date(), transaction_type: "MonthlyPosting" }},
])

// caterer wise amount due
db.meal_entries.aggregate([
    { $match: { date_time: { $gte: ISODate("2020-01-01T00:00:00.000Z"), $lte: ISODate("2020-10-01T00:00:00.000Z") } } },
    { $group: { _id: { caterer_id: "$caterer_id" }, amount: { $sum: "$company_cost" }} },
    { $project: { _id: 0, caterer_id: '$_id.caterer_id', transaction_id: "", amount: 1, date_time: new Date(), transaction_type: "MonthlyPosting" }},
])
