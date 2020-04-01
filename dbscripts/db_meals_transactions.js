var insertRandomFoodTransactions = (num) => {
    let i = 0
    for(i = 0;i < num;i++) {
        let n = db.mealtypes.count()
        let r = Math.floor(Math.random() * n)
        let meal = db.mealtypes.find().limit(1).skip(r).toArray()[0]
        n = db.users.count();
        r = Math.floor(Math.random() * n)
        let user = db.users.find().limit(1).skip(r).toArray()[0]
        db.meals.insertOne(
            { uid: user.uid, datetime: ISODate(), type: meal.type, cost: meal.cost }
        )
        sleep(2000)
    }
}

// var randomDate = function () {
//     var day = 1000 * 60 * 60 * 24;
//     return new Date(Date.now() - (Math.floor(Math.random() * day)));
// }

// var randomName = function() {
//     return (Math.random()+1).toString(36).substring(2);
// }

// var randomInt = function(min, max) {
//     return Math.random() * (max - min) + min;
// }

// var insertRandom = (num) => {
//     let i = 0
//     for(i = 0;i < num;i++) {
//         db.stuff.insertOne(
//             { uid: randomName(), datetime: randomDate(), type: randomName(), cost: randomInt(0, 10000) }
//         )
//     }
// }


// var food = db.food_types.findOne({type: "lunch"})
// db.food_transactions.insertOne(
//     { uid: "2", datetime: ISODate(), type: food.type, cost: food.cost }
// )