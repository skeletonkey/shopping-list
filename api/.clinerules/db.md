# Database

This application is designed to be self hosted, therefor, the database being used is SQLite.

uuid fields are unique identifiers that can be exposed by the API and to users. These IDs should be auto generated so the caller doesn't need to provide them. This ensures uniqueness and a standard format.

id fields should only be used internally and not be exposed.

## Tables

### User

User table describes a user.

Name: user

Fields:

- id int auto incrementing not null
- uuid not null
- username string no null unique
- first_name string not null
- last_name string not null
- password encrypted string not null
- active bool not null  default true

### User Type

The type of the user - this indicates security level and what the user has the ability to do or see.

Root will have access to everything.
Admin will be able to create shopping lists.
User will be able to be assigned to a shopping list and add/remove items to that list.
Root can do everything that an admin can and admin can do everything that a user can.

Table Name: user_type

Fields:

- id int not null auto incrementing
- name string not null
- level int not null default 99

#### Values

|name|level|
|--|--|
|root|0|
|admin|5|
|user|10|

### User Level

Ties a User to a User Type for a given List.

The root user type will never have an entry here as that user has access to everything.

There should only be one entry per user per list. When a user is removed from a list the entry in this table will be deleted. This information does not need to be preserved as logging/historical information will be kept showing the user's participation in a list.

Table Name: user_level

Fields:

- id int not null auto increment
- user_id int not null must exist in the user table
- list_id int non null must exist in the list table
- user_type_id int not null must exist in the user_type table

### List

Name of lists that users can add/remove items to.

Names of lists are not unique. When they are set to not active they will remain inactive, they will remain for logging/history, but will not be deleted nor re-activated.

Table Name: list

Fields:

- id int not null auto incrementing
- uuid not null
- name string not null
- active bool default true

### Item

This table holds all items that can be added/removed from lists.
The list_id field indicates if the item is custom to a specific list and therefore is only available for that list. A list_id of 0 indicates that the item is globally available.
Item names should be unique when combined with the list_id, but should also not be duplicated if it's available globally.

Table Name: item

Fields:

- id int not null auto increment
- uuid not null
- item string not null
- list_id not null default 0 or it must exist in the list table

#### Item/List ID uniqueness

|item|list_id|
|--|--|
|Green Apple|1|
|Green Apple|2|
|Egg|0|

I can add another "Green Apple" to a new list, however, I can not add a "Green Apple" entry for list 1 or list 2. I can not add new "Egg" entry since it's list_id is 0 indicating that it is globally available.
