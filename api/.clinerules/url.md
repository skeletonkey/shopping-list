# URL

The HTTP server being used is github.com/labstack/echo/v4.

If a URL does not exist (meaning that the family_name has been deleted or list has been deleted) then the user is redirected to a 404 page.

The URL structure of the API service follows this pattern:

base_url - the host name plus the current API version (e.g., "/app/v1")
family_name - represents the "name" from the family table
list_name - represents the list name from the list table

## GET http://base_url/family_name

return array of all list for the family

HTTP Code: 200

```json
{
    "data": [
        {
            "name": "list1",
            "display_name": "List One",
            "uuid": "xxxxxxx"
        }
    ]
}
```

## POST http://base_url/family_name

Creates the list

JSON BODY:

```json
{
    "name": "list name",
    "display_name": "Optional Display Nane"
}
```

Returns all lists including the new one:

HTTP Code: 201

```json
{
    "data": [
        {
            "name": "list name",
            "display_name": "List Name",
            "uuid": "xxxxxxx"
        }
    ]
}
```

## GET http://base_url/family_name/list_name/items

Get a list of available item for this list. It will be any items that are globally available and available for the list_name.

HTTP Code: 200

```json
{
    "data": [
        {
            "name": "item name",
            "uuid": "xxxxxxxxx"
        }
    ]
}
```

## DELETE http://base_url/family_name/list_name

Delete list_name

This should cause any items on the list to be marked as deleted first.

HTTP Code: 204

## GET http://base_url/family_name/list_name

return array of all items on the list

HTTP Code: 200

```json
{
    "data": [
        {
            "name": "item name",
            "uuid": "xxxxxxxxxx"
        }
    ]
}
```

## POST http://base_url/family_name/list_name/XXXXXXXXXXXX

XXXXXXXXX can be a UUID or string representing a new item for that list

add item "XXXXXXXXX" (from url) to list_name

"XXXXXXXXX" may not be a UUID. If it is a string then it will be added to the item/list table. This will be done in the background and will be part of the items return list on the next call to the items endpoint.

HTTP Code: 201

## DELETE http://base_url/family_name/list_name/XXXXXXXXXXX

XXXXXXXXXXXX has to be a UUID - general strings will not be honored

remove item "XXXXXXXXX" (from url) to list_name

HTTP Code: 204
