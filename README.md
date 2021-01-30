# TalentPro
Assignment

### Steps: First Create the Staticfile softlink
ln -s static_path localhost

### How to Run the Code
Make app

## API's with Request and Response

### Prime Number Request
curl -H "Content-type:application/json" -X GET 'http://:50051/talentpro/1/prime-number?number=13'
#### Prime Number Response
{"success":true,"message":"Success","code":200,"data":[13,2,3,5,7,11]}

### Date Request
curl -X POST -k http://localhost:50051/talentpro/1/last-day --header "Content-Type:application/json" -d '{"Date":"2021-01-12T11:45"}'
### Date Response
{"success":true,"message":"Success","code":200,"data":{"Date":"2021-01-12T11:45","LastDayOfMonth":31}}

### User List Request 
curl -H "Content-type:application/json" -X GET 'http://:50051/talentpro/1/user/list'
### User List Response 
{"success":true,"message":"Success","code":200,"data":{"id":"fa1be559-7b34-486b-add6-fee3e9b52220","loginId":"1","fullName":"Sumit Thakur","state":1}}

### User Add Request
curl -X POST -k http://localhost:50051/talentpro/1/user/add --header "Content-Type:application/json" -d '{"loginId":"RDua2.0","fullName":"Rahul Dua2"}'
### User Add Response
{"success":true,"message":"Success","code":200,"data":null}

### User Delete Request
curl -H "Content-type:application/json" -X DELETE 'http://:50051/talentpro/1/user/delete?id=233d3c39-d0c8-49ce-baae-4310d2379f36'
### User Delete Response
{"success":true,"message":"Delete with Success","code":200,"data":null}

### User Edit Request
curl -X PATCH -k http://localhost:50051/talentpro/1/user/edit --header "Content-Type:application/json" -d '{"loginId":"RDua3.0","fullName":"Rahul Dua3","id":"7560a184-c848-4b8f-b3f0-033dcfcc9eaf"}'
### User Edit Response
{"success":true,"message":"Update user details with Success","code":200,"data":null}

### Words count crawl Request
curl -H "Content-type:application/json" -X GET 'http://:50051/talentpro/1/url-crawl?url=https://www.amazon.com'
### Words count crawl Response
{"success":true,"message":"Success","code":200,"data":{"ACX":1,"AbeBooks":1,"AccessoriesShop":1,"AccessoriesVideo":1,"Account":3,"AccountYour":1,"Actionable":1,"Ads":1,"Advertising":1,"AffiliateAdvertise":1,"After":2,"Alerts":1,"Alexa":1,"All":3,"Amazon":23,"AmazonBasicsSee":1,"AmazonBecome":1,"AmazonClick":1,"AmazonGlobal":1,"AmazonInvestor":1,"AmazonSell":1,"Analytics":1,"Animals":1,"App":1,"Apparel":1,"Arts":1,"AssistantHelp":1,"Audible":1,"Audio":1,"Audiobook":1,"Automotive":1,"Baby":1,"BabyShop":1,"Back":1,"BalanceAmazon":1,"Based":1,"Beauty":3,"Best":1,"Book":2,"Books":5,"Box":2,"Boxes":1,"Boys":1,"Brands":1,"Business":3,"CDs":1,"COVID":2,"CardShop":1,"Cards":1,"Care":2,"Cart":1,"CategoryComputers":1,"Celebrities":1,"Chance":1,"Click":1,"Clothing":1,"Cloud":2,"Comfy":1,"ComiXology":1,"Comics":1,"Computers":3,"Computing":1,"Conditions":1,"Content":1,"ConverterLet":1,"Crafts":1,"Create":1,"Crime":1,"Currency":1,"Customer":2,"Customers":1,"DPReview":1,"DVDs":1,"Dane":1,"Data":1,"Deals":4,"Deliver":1,"Delivery":2,"Departments":1,"Depository":1,"Designer":2,"DevicesAmazon":2,"Digital":5,"Direct":2,"Disability":1,"Discover":1,"Distribution":1,"Dollar":1,"Drive":1,"East":1,"Easy":3,"Educational":1,"Electronics":2,"Entertainment":1,"Every":1,"Everything":1,"Experienced":1,"Explore":1,"FREE":1,"Fabric":1,"Fashion":6,"Find":3,"For":1,"Free":1,"Fun":1,"Games":3,"GamesBabyToys":1,"GamesShop":1,"Get":5,"Gift":1,"Gifts":1,"Girls":1,"Goodreads":1,"Guarantee":1,"Happiness":1,"Health":1,"Hello":1,"Help":1,"Home":4,"Household":1,"Hub":1,"IMDb":1,"IMDbPro":1,"If":1,"Ignite":1,"Improvement":1,"Inc":1,"India":2,"Indie":1,"Industrial":1,"Info":1,"Internationally":1,"K":1,"Kindle":3,"Kitchen":1,"Knitting":1,"Know":1,"Laptops":1,"Learn":1,"Listen":1,"Lists":1,"Luggage":1,"Made":3,"Make":1,"Men":2,"Mojo":1,"Money":2,"More":1,"Movie":1,"Movies":2,"Music":3,"Must":1,"Need":1,"Neighbors":1,"Next":1,"NoticeInterest":1,"Office":2,"Online":1,"Orders":2,"OrdersShipping":1,"Original":1,"Our":1,"Outdoors":1,"Pass":1,"Payment":1,"Performances":1,"Personal":2,"Pet":1,"Pharmacy":1,"Photography":1,"PillPack":1,"PointsReload":1,"PoliciesReturns":1,"Previous":1,"Prime":5,"Print":1,"ProductsAmazon":1,"ProductsSelf":1,"Professionals":1,"PromotionsShop":1,"Pros":1,"Publish":1,"Publishing":3,"Quick":1,"Quilting":1,"Rapids":1,"Rates":1,"Real":1,"Registry":1,"RelationsAmazon":1,"ReplacementsManage":1,"Resources":1,"Returns":1,"Ring":1,"Room":1,"S":1,"Safety":1,"Scalable":1,"Scientific":1,"Score":1,"Second":1,"Security":1,"See":2,"Select":1,"Sell":3,"Sellers":1,"SellersShop":1,"Selling":1,"Service":1,"Services":3,"Sewing":1,"Shenanigans":1,"Ship":1,"Shoes":1,"Shop":5,"Shopbop":1,"Shopping":1,"Sign":2,"Simplified":1,"Skip":1,"Smart":1,"Software":1,"Sports":1,"Start":1,"StatesChoose":1,"Stay":1,"Store":1,"Stream":2,"Stuffed":1,"Subscription":1,"Supplies":1,"Support":1,"Systems":1,"TV":3,"TVSee":1,"TabletsSee":1,"The":1,"There":1,"Thousands":1,"Time":1,"Today":1,"Tools":1,"Top":3,"ToursMake":1,"Toys":2,"ToysShop":1,"U":1,"USD":1,"Us":1,"UsAmazon":1,"UsCareersBlogAbout":1,"UsHost":1,"UsSell":1,"UsePrivacy":1,"Video":6,"View":2,"Vinyl":1,"Web":2,"Welcome":1,"WiFi":1,"Wireless":1,"With":1,"Women":1,"Woot":1,"Worldwide":1,"You":2,"YouAmazon":1,"Your":7,"Zappos":1,"a":6,"about":1,"access":3,"affiliates":1,"all":1,"also":2,"amazon":3,"amp":7,"an":4,"and":8,"app":1,"apps":1,"are":3,"art":1,"at":2,"attract":1,"audio":1,"bEnglishChoose":1,"bUnited":1,"back":2,"beddingSee":1,"best":1,"books":1,"bought":1,"boxes":1,"brands":1,"browsing":2,"by":1,"can":1,"categoriesSee":1,"collectibles":1,"com":6,"content":1,"country":1,"customers":1,"deals":1,"delivery":2,"department":1,"detail":2,"details":1,"doctype":1,"door":1,"easy":2,"edit":2,"eero":1,"engage":1,"enjoy":1,"exclusive":1,"experience":1,"experienceSign":1,"fashion":1,"fast":1,"favorite":1,"featured":2,"features":1,"find":2,"fit":1,"for":9,"free":1,"from":1,"give":1,"go":2,"gt":1,"have":1,"herSweatshirtsJoggersCardigansEasy":1,"here":3,"history":2,"home":1,"homeExplore":1,"html":1,"ideal":1,"in":11,"interested":2,"is":1,"it":3,"items":2,"its":1,"kids":1,"language":1,"learn":1,"life":1,"lightsShop":1,"loading":1,"local":1,"look":3,"main":2,"members":1,"menu":1,"millions":2,"mobile":3,"more":9,"movies":1,"music":1,"navigate":2,"nbsp":8,"now":14,"of":6,"on":9,"or":3,"original":2,"page":2,"pages":4,"picksShop":1,"pm":1,"prefer":1,"problem":1,"product":3,"products":2,"productsShop":2,"recently":2,"recommendations":3,"region":1,"response":1,"reviews":1,"right":2,"rsaquo":1,"s":6,"search":1,"second":1,"securely":1,"series":1,"shop":1,"shopping":3,"shows":1,"similar":1,"simplified":1,"songs":1,"started":1,"storage":1,"stories":1,"strip":1,"styles":1,"subscription":1,"teesSee":1,"the":8,"this":1,"to":21,"top":2,"trade":1,"try":1,"under":2,"us":1,"version":2,"viewed":2,"viewing":2,"want":1,"way":2,"web":2,"website":1,"with":7,"www":1,"x":2,"you":4,"your":5}}
