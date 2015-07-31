# ambition

ambition is an action tracker for the purpose of quantitative self and habit tracking.

The core will be database with a useful structure and an api that can be accessed by many different interfaces for both data input and viewing.

## Note

This is a learning project, please point out anything that is dumb or that could be done better.

## Table structure

Users have Sets, Sets have Actions, and Actions have Occurrences.

A user would be person (or a group? What about groups?) A set would be a logical group of actions. An action is an thing that can be done, binary in completion. An occurrences is an instance of an action being done. An occurrence may have metadata. The action determines whether their occurrences will have metadata, the metadata structure, and whether the metadata is optional or not.

Structure that will be added after the above is added:

Actions belonging to multiple sets. i.e. a many to many relationship between actions and sets. That way an action can be in multiple sets if that makes logical sense.

Sets belonging to multiple users. The idea of user privileges over sets needs to be discussed. The hope is for groups to have a set of actions that can be completed as a group and this would be a useful tracking tool.

Actions could belong to a set which belongs to multiple users, and all users could do any action but the metadata has the user who did the action in it. Or the occurrence could belong to both a user and an action. I actually like that. No action belongs to a user but the occurrences do belong to a particular user. Which is true in real life too. No one owns a specific activity, but they do own when they do the activity.

The problem with the paragraph above is then who is able to edit the action name and the metadata? There are two routes as far as I am concerned.

- Either actions are owned by one user and then other users can "clone" an action.
- Or actions are not owned by anyone and cannot be mutated. But can be cloned and have different attributes

Very quickly I can see the first one being better. I am still wanting to work out an idea for group sets, but I think I need to build the first part before I worry about it.

The metadata structure that an action defines and that an occurrence makes an instance of, must be valid json. The action provides the structure, by defining objects, arrays, and values. Typing all values in the json. Arrays are not allowed to have mixed types, all structures in any array must be of the same type. The base required structure is the outmost structure is an object with at least one key and one value. i.e.

	{ key: string }

Some example uses:

	{ weight: number }

	{ data: string }

	{ duriation: number, unit: text }



### Sets

id  | SetName
--- | --------
1   | Health

### Actions

id  | ActionName (char 255)   | metadata (text)
--- | ----------------------- | ------
1   | Brush Teeth             | null
2   | Fluoride Tray           | null
3   | Weigh Self              | { weight: number }

### Occurrences

id | Time    | ActionId | metadata (text)
-- | -----   | ------   | --------
1  | ISO8601 | 1        |
2  | ISO8601 | 2        |
3  | ISO8601 | 2        |
4  | ISO8601 | 3        | { weight: 155.0 }
5  | ISO8601 | 2        |
6  | ISO8601 | 3        | { weight: 156.3 }
7  | ISO8601 | 1        |

### Users

id  | username   | email               | hashed_password
--- | ---------- | -------             | -----------------
1   | adamryman  | adamryman@gmail.com | fkj4382fmwf83nvyd0h4nfi3gha04nfysb3ymoq1

## Relations

- Users have many Sets
- Sets have many Users
- Sets have many Actions
- Actions have many Sets
- Actions have many Occurrences
- Occurrences have one Action

### Relation Tables

Possible joins / How a user would want data

- All occurrences of an action
- All occurrences of an action within a time frame (order / id on time)
- All sets of a user
- All actions of a set
- All sets that contain an action

#### Sets -> Actions Table

SetId  | ActionId
------ | ----------
1      | 1
1      | 2
1      | 3

#### Users -> Sets

UserId  | SetId
------- | ---------
1       | 1

### Future Relations Tables

#### Actions -> Sets

#### Sets -> Users/Groups

#### Privileges Tables

## API

Research into api best practices has to be done.

Things to learn:
- How to properly do api versioning.
- Authentication
	- Just an http auth header?
	- Cookies?
	- Oath?
	- Other services?
- Public entities.
- How much data to return
- Query params for verbosity

**Note: For now, only select data will return. All data will be accessable through some request. After a discussion, will determine way to specify what data to return, possibly with query params.**


### Base URL

	myambition.io

### Authentication

I think that going the way of one of the answers in this stackoverflow post will work: http://stackoverflow.com/questions/25218903/how-are-people-managing-authentication-in-go

For now though, there will be a secrete query param that will be sent along to validate requests and there will not be users. Once auth is setup we can make an association between users to sets and users to actions.

	POST myambition.io/auth/login
	JSON { username: "", hashed_password: }

This will return some kind of auth token, this token is sent with all requests to indicate a user. This will handle privileges

How about public sets and actions?

### Sets

	// List all sets of a user
	GET	myambition.io/sets

	// Details about set :id
	GET myambition.io/sets/:id

	// List all actions in set :id
	GET myambition.io/sets/:id/actions

	// Add a set
	POST myambition.io/sets
	JSON { setname : "Exercise", actions: [actionid, actionid] }

	// Modify a set, (change name, add actions, remove actions, eventually users)
	PUT myambition.io/sets/:id
	JSON, all optional { setname: "" , addactions: [], removeactions: [] }

	// Delete a set
	DELETE myambition.io/sets/:id

### Actions

	// List all actions assosiated with a user
	GET myambition.io/actions
	// Response [ { id: id, actionname: name, data: json }, ... ]

	// Details about action :id
	GET myambition.io/actions/:id
	// Response { id: id, actionname: name, data: json }

	// Add an action
	POST myambition.io/actions
	JSON { actionname: "Handstands", metadata: { duration : double, otherdata: [something, another], onelastthing: { inner: text } } }
	// Response	{ status: "OK" } or { status : "ERROR" }

	PUT myambition.io/actions/:id
	JSON, all optional { actionname: "", metadata: json }

	// Delete an action ( And all occurrences of that action )
	DELETE myambition.io/actions/:id

### Occurrences

	// Occurrences by current user of a action :id ( default: last week of occurrences )
	GET myambition.io/actions/:id/occurrences
	// Response [ { id: id, time: ISO8601, metadata: json }, ... ]
	// Query params: ?from=ISO8061&until=ISO8061 or ?time=( all, month, week, day, hour )

	// Add an occurrence
	POST myambition.io/actions/:id/occurrences
	JSON { time: ISO8601, metadata: json }
	// To start, data will not be validated or required. It is up to clients to throw out malformed data and handle missing portions

	// Modify an occurrence
	PUT myambition.io/actions/:id/occurrences/:id or PUT myambition.io/occurrences/:id
	JSON { time: ISO8601, metadata: json, actionid: differentactionid }

