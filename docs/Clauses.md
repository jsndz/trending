SQL clauses are parts of an SQL statement that define what the query does.

Common clauses:

* `SELECT` → choose columns
* `FROM` → choose table
* `WHERE` → filter rows
* `GROUP BY` → group rows
* `HAVING` → filter grouped rows
* `ORDER BY` → sort results
* `LIMIT` → restrict number of rows
* `JOIN` → combine tables

Example:

```sql 
SELECT name, age
FROM users
WHERE age > 18
ORDER BY age DESC
LIMIT 10;
```

Breakdown:

* `SELECT name, age` → columns
* `FROM users` → source table
* `WHERE age > 18` → filter
* `ORDER BY age DESC` → sorting
* `LIMIT 10` → max rows returned

Clauses together form a complete SQL query.



```go 
db.Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: "url"}},
    DoNothing: true,
})
```

here we create a clause if any confilct arises with columns URL just ignore it

same as `ON CONFLICT (url) DO NOTHING`