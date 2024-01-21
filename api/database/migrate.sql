
-- Table: user
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
	"id"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT NOT NULL,
	"nick"	TEXT NOT NULL UNIQUE,
	"email"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	"created_at"	DATETIME NOT NULL DEFAULT current_timestamp,
	PRIMARY KEY("id" AUTOINCREMENT)
);

-- Table: follower
DROP TABLE IF EXISTS "follower";
CREATE TABLE "follower" (
	"user_id"	INTEGER NOT NULL,
	"follower_id"	INTEGER NOT NULL,
	FOREIGN KEY("user_id") REFERENCES "user"("id") ON DELETE CASCADE,
	FOREIGN KEY("follower_id") REFERENCES "user"("id") ON DELETE CASCADE,
	PRIMARY KEY("user_id","follower_id")
);