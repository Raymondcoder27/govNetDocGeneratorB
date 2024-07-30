The jsonwId file is a separate implementation where I create a document struct that accepts ID and Body fields.
The program assumes that the json data is divided into an id field and a body object which includes the document body, this was done to provide for when the user specifys the id and body of the json data separately.
The program maps and binds the json data to the body object and creates a random id incase the json data has no id specified.
