HOST=localhost
PORT=27017
ADMIN_USER=root
ADMIN_PWD=rootpassword
ADMIN_DB=admin
DB=superheroes
USER=test
PWD=12345

mongo --host $HOST --port $PORT -u $ADMIN_USER -p $ADMIN_PWD --authenticationDatabase $ADMIN_DB <<EOF

use $DB;

db.createCollection("${COLLECTION}");

db.createUser({ user: "${USER}", pwd:  "${PWD}", roles: [ { role: "readWrite", db: "${DB}" }] });

db["test"].insertOne( { admin: "admin", user: "user" } );

EOF
