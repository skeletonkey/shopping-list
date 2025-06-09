# Family

Families are created via configuration. In the config there is a 'family' field with the following data:

- name: Name of the family and the first part of the url
- display_name: this is what is displayed inside the UI

When a family is added it will be inserted into the family table and create in the database.

When a family is removed it leads a cascade of removal:

- set all items to all lists to removed (false)
- delete all lists
- delete the family entry

As an effect of the family entry removal is that the URL http://base_url/family_name no longer resolves.
