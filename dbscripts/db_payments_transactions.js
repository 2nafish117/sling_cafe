var insertRandomPaymentTransactions = (num) => {
    let i = 0
    for(i = 0;i < num;i++) {
        let n = db.users.count();
        let r = Math.floor(Math.random() * n)
        let randomuser = db.users.find().limit(1).skip(r).toArray()[0]

        let modes = ['upi', 'credit card', 'debit card', 'cash']
        n = modes.length
        r = Math.floor(Math.random() * n)
        let randommode = modes[r]

        let randomamount = Math.round(Math.random() * 70) + 10
        db.payments.insertOne(
            { uid: randomuser.uid, mode: randommode, amount: randomamount, datetime: ISODate() }
        )
        sleep(2000)
    }
}
