db.users.insertMany([
    {uid: "1", fname: "shashank", lname: "c", email: "shashank.c@dish.com"},
    {uid: "2", fname: "amruth", lname: "r", email: "amruth.r@dish.com"},
    {uid: "3", fname: "beep", lname: "is", email: "beep.is@dish.com"},
    {uid: "4", fname: "yeet", lname: "skeet", email: "yeet.skeet@dish.com"},
    {uid: "5", fname: "beet", lname: "meet", email: "beet.meet@dish.com"},
    {uid: "6", fname: "simp", lname: "son", email: "simp.son@dish.com"},
    {uid: "7", fname: "rocket", lname: "man", email: "rocket.man@dish.com"},
    {uid: "8", fname: "ocean", lname: "man", email: "ocean.man@dish.com"},
    {uid: "9", fname: "bart", lname: "simpson", email: "bart.simpson@dish.com"},
    {uid: "10", fname: "mike", lname: "wazauski", email: "mike.wazauski@dish.com"},
    {uid: "11", fname: "gordon", lname: "freeman", email: "gordon.freeman@dish.com"},
    {uid: "12", fname: "alyx", lname: "vance", email: "alyx.vance@dish.com"},
    {uid: "13", fname: "eli", lname: "vance", email: "eli.vance@dish.com"},
    {uid: "14", fname: "bob", lname: "butcher", email: "bob.butcher@dish.com"},
    {uid: "15", fname: "cortana", lname: "wort", email: "cortana.wort@dish.com"},
    {uid: "16", fname: "thel", lname: "vadam", email: "thel.vadam@dish.com"},
    {uid: "17", fname: "master", lname: "chief", email: "master.chief@dish.com"},
    {uid: "18", fname: "bob", lname: "marley", email: "bob.marley@dish.com"},
    {uid: "19", fname: "sahaj", lname: "khota", email: "sahaj.khota@dish.com"},
    {uid: "20", fname: "weeb", lname: "lord", email: "weeb.lord@dish.com"},
    {uid: "INT1", fname: "red", lname: "eye", email: "red.eye@dish.com"},
    {uid: "INT2", fname: "thunder", lname: "thund", email: "thunder.thund@dish.com"},
    {uid: "INT3", fname: "sparks", lname: "sporks", email: "sparks.sporks@dish.com"},
    {uid: "INT4", fname: "nader", lname: "nades", email: "nader.nades@dish.com"},
    {uid: "INT5", fname: "fragger", lname: "frags", email: "fragger.frags@dish.com"},
    {uid: "INT6", fname: "doom", lname: "guy", email: "doom.guy@dish.com"},
    {uid: "INT7", fname: "light", lname: "yagami", email: "light.yagami@dish.com"},
    {uid: "INT8", fname: "nishimiya", lname: "shouko", email: "nishimiya.shouko@dish.com"},
    {uid: "INT9", fname: "bubble", lname: "sort", email: "bubble.sort@dish.com"},
    {uid: "INT10", fname: "bucket", lname: "sort", email: "bucket.sort@dish.com"}
])
db.users.createIndex({ uid: 1})

db.meals.insertMany([
    {type: "breakfast", cost: 30, start_time: ISODate("2020-01-01T09:00:00.000Z"), end_time: ISODate("2020-01-01T11:00:00.000Z")},
    {type: "lunch", cost: 60, start_time: ISODate("2020-01-01T13:00:00.000Z"), end_time: ISODate("2020-01-01T15:00:00.000Z")},
    {type: "snack", cost: 0, start_time: ISODate("2020-01-01T16:30:00.000Z"), end_time: ISODate("2020-01-01T18:00:00.000Z")}
])