# Shopping List Design - Grug's Simple Way

DATE: 2025-06-14
AUTHOR: grug (simplicity shaman)

* GRUG UNDERSTAND PROBLEM 🪨

grug read task. grug see family need shopping list. family not want learn new thing. family want add milk to list. other family see milk on list. simple like rock.

existing solution too complex. google keep try do everything. apple reminder try be smart. family just want list. grug make list.

** what family actually need:

* see list of thing to buy
* add thing to list  
* remove thing from list
* other family see same list
* work on phone and computer
* not break when internet slow

** what family NOT need (complexity demon try trick grug):

* fancy animation
* 47 different view
* ai suggestion
* integration with recipe app
* price tracking
* analytics dashboard
* notification system
* dark mode toggle
* user preference panel

grug say: if family need explain how use app, grug failed.

* GRUG DESIGN PHILOSOPHY FOR THIS PROBLEM 🔥

** core grug rule for shopping list:

1. one screen do everything
2. big button for add thing
3. big button for remove thing  
4. list show all thing
5. no menu. no setting. no configuration.

** grug simplicity test:

* can 5 year old use? yes = good
* can grandma use without help? yes = good  
* need instruction manual? no = good
* work when wifi slow? yes = good

** grug complexity demon detector:

* need login screen? DEMON DETECTED
* need tutorial? DEMON DETECTED  
* need setting page? DEMON DETECTED
* need user manual? DEMON DETECTED

* GRUG RESEARCH EXISTING SIMPLE THING 🦣

grug look at what work in real world:

** physical shopping list (paper):

* write thing down
* cross thing off
* everyone see same paper
* work when power out
* no learning curve

** what make paper list good:

* one action: write
* one action: cross off
* visual: see all thing at once
* shared: everyone use same paper

** what make app list bad:

* need account creation
* need permission setting
* need sync configuration  
* need app download
* need notification permission

grug learn: best app copy paper, not try be smarter than paper.

* GRUG SOLUTION: SHOPPING LIST THAT WORK LIKE PAPER 📝

** grug main screen design:

```text
    FAMILY SHOPPING LIST
    
    [  type new item here...  ] [ADD]
    
    ☐ milk (2 gallon)           [X]
    ☐ bread                     [X] 
    ☐ banana (bunch)            [X]
    ☐ chicken breast (2 lb)     [X]
    
    4 items on list
```

that it. one screen. do everything.

** how grug screen work:

1. family open web page
2. family see list of thing to buy
3. family type new thing in box
4. family click ADD or press enter
5. thing appear on list for everyone
6. family click X when buy thing
7. thing disappear from list for everyone

no login. no account. no setting. no menu. just list.

** grug make it even simpler:

* text box always focused (can start typing immediately)
* enter key add item (no need click button)
* click anywhere on item to remove (big target)
* auto-save everything (no save button)
* auto-refresh every 30 second (no refresh button)

* GRUG TECHNICAL DESIGN (KEEP SIMPLE) ⚙️

** grug use existing go api structure:

* db module: already have user/list/item table (good rock)
* server module: serve web page and api (good rock)  
* eventing module: handle sync between user (good rock)

grug not change database. database already simple. grug like.

** grug web ui architecture:

```text
index.html (one file)
├── html: show list and input box
├── css: make button big and text readable  
└── javascript: add item, remove item, refresh list
```

one html file. no framework. no build step. no npm. no webpack. no complexity demon.

** grug api endpoint (use existing go structure):

```text
GET  /                    -> serve index.html
GET  /api/items          -> get all item for family list
POST /api/items          -> add new item to list
DELETE /api/items/{id}   -> remove item from list
```

four endpoint. that all. no rest api with 47 endpoint. no graphql. no complexity demon.

** grug sync strategy:

javascript poll /api/items every 30 second. if list change, update screen. simple like rock.

no websocket. no server-sent event. no real-time complexity demon. 30 second good enough for shopping list. family not buying milk every 5 second.

* GRUG HANDLE MULTI-USER (KEEP SIMPLE) 👥

** grug family model:

one family = one list. family share same url. no user account. no permission. no complexity demon.

example:

* family get url: [http://pi.local:8080/family/smith](http://pi.local:8080/family/smith)
* everyone in smith family use same url
* everyone see same list
* everyone can add/remove item

** grug security model:

security through obscurity. url have random part. only family know url.

example url: [http://pi.local:8080/family/abc123def456](http://pi.local:8080/family/abc123def456)

if someone guess url, they see smith family shopping list. but who care? is just shopping list. not bank account.

** grug handle quantity:

when add item, can type quantity in parenthesis:

* "milk (2 gallon)"
* "banana (bunch)"  
* "chicken (2 lb)"

grug parse this simple. no dropdown. no unit selector. no complexity demon.

* GRUG IMPLEMENTATION PLAN 🔨

** phase 1: make it work (1 day)

1. create index.html with list and input box
2. add 4 api endpoint to existing go server
3. add javascript for add/remove/refresh
4. test with one family

** phase 2: make it reliable (1 day)  

1. add error handling (what if api down?)
2. add offline support (save to browser storage)
3. add auto-retry (if add fail, try again)
4. test with slow internet

** phase 3: make it fast (1 day)

1. optimize database query
2. add simple caching
3. minimize javascript
4. test with many item

total: 3 day. not 3 month. not 3 year. 3 day.

** grug file structure:

```text
api/
├── app/
│   ├── main.go (already exist)
│   ├── db/ (already exist)  
│   ├── server/ (already exist)
│   └── eventing/ (already exist)
├── static/
│   └── index.html (grug create this)
└── family/
    └── handler.go (grug create this)
```

grug add two file. not 200 file. two file.

* GRUG HANDLE EDGE CASE (BUT KEEP SIMPLE) 🪨

** what if two person add same item same time?
grug say: so what? now have two milk on list. family buy two milk. not end of world.

** what if internet slow?
grug say: item save to browser first. sync when internet back. family still see their item.

** what if raspberry pi crash?
grug say: family write on paper until pi back. like old time. family survive.

** what if family want delete whole list?
grug say: refresh page. start new list. old list still in database for history.

** what if family want multiple list?
grug think more about this. family actually need multiple list. grocery list different from hardware list. grug make simple way.

grug update design: family page show all list. still simple.

* GRUG UPDATED DESIGN: FAMILY AND LIST MANAGEMENT 📋

grug realize family need multiple list. but grug keep simple. no complexity demon allowed.

** grug new url structure:

```text
http://pi.local:8080/smith              -> family page (show all list)
http://pi.local:8080/smith/grocery      -> grocery list
http://pi.local:8080/smith/hardware     -> hardware list  
http://pi.local:8080/smith/christmas    -> christmas list
```

simple like file system. family understand folder concept.

** grug family page design:

```text
    SMITH FAMILY LISTS
    
    [  new list name...  ] [CREATE LIST]
    
    📝 Grocery (5 items)        →
    🔨 Hardware (2 items)       →  
    🎄 Christmas (12 items)     →
    
    3 lists total
```

one screen show all list. big button to make new list. click list name to open.

** grug individual list page design:

```text
    ← BACK TO FAMILY    GROCERY LIST
    
    [  type new item here...  ] [ADD]
    
    ☐ milk (2 gallon)           [X]
    ☐ bread                     [X] 
    ☐ banana (bunch)            [X]
    
    3 items on list
```

same as before but with back button. back button big and obvious.

** grug navigation flow:

1. family go to [http://pi.local:8080/smith](http://pi.local:8080/smith)
2. family see all their list
3. family click "Grocery" to open grocery list
4. family add/remove item from grocery list
5. family click "BACK TO FAMILY" to see all list again
6. family click "Hardware" to open hardware list

simple like paper folder. no complexity demon.

** grug updated api endpoint:

```text
GET  /smith                     -> serve family page
GET  /smith/grocery             -> serve grocery list page
POST /smith/grocery/items       -> add item to grocery list
DELETE /smith/grocery/items/{id} -> remove item from grocery list
GET  /api/smith/lists           -> get all list for smith family
POST /api/smith/lists           -> create new list for smith family
GET  /api/smith/grocery/items   -> get all item from grocery list
```

still simple. just few more endpoint for list management.

** grug list creation:

family type "Hardware" in box, click "CREATE LIST". new list appear immediately. family click on it to start adding item.

no form. no setting. no permission. just name and create.

** grug handle list name:

grug make list name simple:

* only letter and number allowed
* space become dash (Hardware Store -> hardware-store)
* no special character (complexity demon live there)
* if list name already exist, add number (grocery-2, grocery-3)

** grug file structure updated:

```text
api/
├── static/
│   ├── family.html (grug create - show all list)
│   └── list.html (grug create - show one list)
└── family/
    └── handler.go (grug update - handle family and list)
```

still only two html file. grug not make 47 different page.

** grug keep it simple rule:

* family page just show list, no fancy feature
* list page just show item, no fancy feature  
* back button always visible and big
* create list just need name, nothing else
* no delete list button (complexity demon trap)
* no rename list button (complexity demon trap)
* no list setting (complexity demon paradise)

if family want delete list, they stop using it. empty list not hurt anyone.

* GRUG CONCERN AND RISK 🚨

** grug worry about:

1. **family expect fancy ui**: grug solution look simple. maybe too simple? family used to fancy app?
   * grug solution: make button big and colorful. simple not mean ugly.

2. **sync not fast enough**: 30 second polling maybe too slow for some family?
   * grug solution: start with 30 second. if family complain, make 10 second. still simple.

3. **no user account feel weird**: family expect login screen?
   * grug solution: add optional family name input. just for show. not for security.

4. **database grow too big**: what if family add 10000 item over time?
   * grug solution: auto-delete item older than 30 day. family not need old item.

** grug not worry about:

1. **security**: is shopping list, not nuclear code
2. **scalability**: is for family, not facebook  
3. **performance**: sqlite fast enough for shopping list
4. **mobile responsive**: big button work on phone too

* GRUG SUCCESS METRIC 📊

** how grug know solution work:

1. **5 year old test**: give to 5 year old. if they can add "cookie" to list without help, grug win.

2. **grandma test**: give to grandma. if she can use without calling tech support, grug win.

3. **wifi slow test**: use on slow internet. if still work, grug win.

4. **family happy test**: family use for 1 week. if they not complain, grug win.

** grug failure condition:

1. family ask "how do i...?" = grug failed
2. family need help to use = grug failed  
3. family go back to paper list = grug failed
4. family ask for more feature = maybe grug failed, maybe family infected by complexity demon

* GRUG FINAL WISDOM 🧙‍♂️

shopping list is simple problem. grug give simple solution.

no framework. no build tool. no package manager. no complexity demon.

just html + css + javascript + go api. work on any browser. work on any device. work when internet slow.

family get url. family use list. family happy. grug happy.

remember grug rule: **if need explain how use, already too complex.**

shopping list should work like paper list, but sync between family. that all.

grug done. time for mammoth soup. 🦣🍲

---

*"best shopping list app is one that feel like not using app at all"* - ancient grug wisdom

** APPENDIX: GRUG ANTI-PATTERN DETECTED IN OTHER SOLUTION

grug see many shopping list app. all infected by complexity demon:

* **todoist**: try be project manager. shopping not project.
* **any.do**: have 47 view and smart suggestion. family just want add milk.
* **google keep**: try do everything. note, reminder, photo, drawing. complexity demon paradise.
* **apple reminder**: location reminder, time reminder, smart list. family brain already know when buy milk.

all these app make simple thing complex. grug make complex thing simple.

grug way: one screen, big button, work like paper. done.
