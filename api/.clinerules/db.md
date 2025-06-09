# Database

This application is designed to be self hosted, therefor, the database being used is SQLite.

pressly/goose is used to handle DB changes and migration.

uuid fields are unique identifiers that can be exposed by the API and to users. These IDs should be auto generated so the caller doesn't need to provide them. This ensures uniqueness and a standard format.

id fields should only be used internally and not be exposed.

## Tables

### Family

Family is the owner of a list(s). These are created via a config file which the server will handle syncing to the database.

Table Name: family

Fields:

- id int not null auto incrementing
- name string not null unique
- display_name string

### List

Name of lists that users can add/remove items to.

Names of lists are not unique. When they are set to not active they will remain inactive, they will remain for logging/history, but will not be deleted nor re-activated.

If no display_name is found/provided the "name" will be the display_name.

Table Name: list

Fields:

- id int not null auto incrementing
- uuid not null
- name string not null
- display_name string
- family_id int not null

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
