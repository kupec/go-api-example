CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "visits" int DEFAULT(0)
);
