conn = new Mongo();
db = conn.getDB("mydb");

db.Stats.createIndex({ "count": 1 }, { unique: false });
db.Stats.insert({ date: new Date(), count: 1250 });
db.Stats.insert({ date: new Date(), count: 1354 });
db.Stats.insert({ date: new Date(), count: 1556 });
db.Stats.insert({ date: new Date(), count: 1757 });
