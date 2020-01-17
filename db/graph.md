# Pass Around Novel Graph Database Schema

## Nodes

### User

 Property           | Type
--------------------|--------
 username           | string
 password           | string
 password_token     | string?
 email              | string
 email_token        | string?
 phone              | string
 phone_token        | string?
 since              | long
 picture            | string

### Token

 Property           | Type
--------------------|--------
 value              | string
 since              | long
 until              | long

### Role

 Property           | Type
--------------------|--------
 name               | string
 (propertyname)     | any

### Novel

 Property           | Type
--------------------|--------
 name               | string
 uuid               | string
 num_chapters       | int
 since              | long
 last_turn          | long
 done               | bool
 picture            | string
 public             | bool
 turn_min           | int
 turn_max           | int
 turn_type          | int

## Edges

### Auth (Token -> User)

 Property           | Type
--------------------|--------
 ip                 | string

### Friend (User -> User)

 Property           | Type
--------------------|--------
 nickname_fwd       | string
 nickname_rev       | string
 since              | long

### Permissions (User -> Role)

 Property           | Type
--------------------|--------
 priority           | int
 since              | long

### Writer (User -> Novel)

 Property           | Type
--------------------|--------
 starred            | bool
 notify             | bool
 chapter            | int
 pos                | int

### Reader (Novel -> User)

 Property           | Type
--------------------|--------
 starred            | bool
 chapter            | int
 pos                | int
